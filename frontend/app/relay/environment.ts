import { Environment, Network, RecordSource, Store } from 'relay-runtime';
import { QueryResponseCache } from 'relay-runtime';

import { getCookieAction, setCookieAction } from '@actions/cookieActions';

import type { GraphQLError } from '@lib/constants';
import { ErrorMessages } from '@lib/constants';

import type { RefreshTokenMutation$data } from '@relay/__generated__/RefreshTokenMutation.graphql';

import type { CacheConfig, GraphQLResponse, RequestParameters, Variables } from 'relay-runtime';

const PUBLIC_QUERIES = [
  'SignupFormMutation',
  'SigninFormMutation',
  'ScreenNameFieldQuery',
  'EmailFieldQuery',
  'RefreshTokenMutation',
];

// Reference:
// https://github.com/relayjs/relay-examples/pull/241/files#diff-bd1d567983978138ef31a0368880fe8ff1c1f24ed172cd70f1826a13a0d26727

const IS_SERVER = typeof window === typeof undefined;
const CACHE_TTL = 5 * 1000; // 5 seconds, to resolve preloaded results

// All relay query will eventually call this fetch with params.text which is a graphql body compiled by relay
export async function networkFetch(params: RequestParameters, variables: Variables): Promise<GraphQLResponse> {
  const accessToken = await getCookieAction('accessToken');

  const attachHdrAndFetchGQL = async (token: string | undefined) => {
    let headers: HeadersInit = { 'Content-Type': 'application/json' };
    if (token != null) {
      headers = { ...headers, Authorization: `Bearer ${token}` };
    }

    const response = await fetch(process.env.NEXT_PUBLIC_GRAPHQL_URL, {
      method: 'POST',
      headers: headers,
      body: JSON.stringify({
        query: params.text,
        variables,
      }),
    });

    return await response.json();
  };

  const json = await attachHdrAndFetchGQL(accessToken);

  if (json.errors) {
    // Try refreshing accessToken if it is expired.
    if (
      (json.errors as GraphQLError[])[0].extensions.code === 'UNAUTHORIZED' &&
      !PUBLIC_QUERIES.includes(params.name)
    ) {
      const refreshToken = await getCookieAction('refreshToken');
      if (!refreshToken) {
        throw new Error(ErrorMessages.RefreshTokenNotFound);
      }

      const newAccessToken = await refreshAccessToken(refreshToken);
      setCookieAction('accessToken', newAccessToken);

      // Retry the original request with the new access token
      const jsonWithNewAccToken = await attachHdrAndFetchGQL(newAccessToken);

      if (jsonWithNewAccToken && jsonWithNewAccToken.errors) {
        throw (jsonWithNewAccToken.errors as GraphQLError[])[0];
      }

      return jsonWithNewAccToken;
    }

    // TODO: How to deal with multiple errors (what if there are errors[1], errors[2], etc.)?
    throw (json.errors as GraphQLError[])[0];
  }

  return json;
}

async function refreshAccessToken(refreshToken: string): Promise<string> {
  const query = `mutation RefreshAccessToken($refreshToken: String!) {
    refreshToken(refreshToken: $refreshToken) {
      accessToken
    }
  }`;

  const response = await fetch(process.env.NEXT_PUBLIC_GRAPHQL_URL, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      query,
      variables: { refreshToken },
    }),
  });

  const json = await response.json();
  if (json.errors) {
    throw json.errors as GraphQLError[];
  }

  return (json.data as RefreshTokenMutation$data).refreshToken.accessToken;
}

export const responseCache: QueryResponseCache | null = IS_SERVER
  ? null
  : new QueryResponseCache({ size: 100, ttl: CACHE_TTL });

function createNetwork() {
  async function fetchResponse(params: RequestParameters, variables: Variables, cacheConfig: CacheConfig) {
    const isQuery = params.operationKind === 'query';
    const cacheKey = params.id ?? params.cacheID;
    const forceFetch = cacheConfig && cacheConfig.force;
    if (responseCache != null && isQuery && !forceFetch) {
      const fromCache = responseCache.get(cacheKey, variables);
      if (fromCache !== null) {
        return Promise.resolve(fromCache);
      }
    }

    return networkFetch(params, variables);
  }

  const network = Network.create(fetchResponse);
  return network;
}

function createEnvironment() {
  return new Environment({
    network: createNetwork(),
    store: new Store(RecordSource.create()),
    isServer: IS_SERVER,
  });
}

export const environment = createEnvironment();

export function getCurrentEnvironment() {
  if (IS_SERVER) {
    return createEnvironment();
  }

  return environment;
}

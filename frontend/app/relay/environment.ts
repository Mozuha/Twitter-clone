import { Environment, Network, RecordSource, Store } from 'relay-runtime';
import { QueryResponseCache } from 'relay-runtime';

import type { GraphQLError } from '@types-constants/form';

import type { CacheConfig, GraphQLResponse, RequestParameters, Variables } from 'relay-runtime';

// Reference:
// https://github.com/relayjs/relay-examples/pull/241/files#diff-bd1d567983978138ef31a0368880fe8ff1c1f24ed172cd70f1826a13a0d26727

const IS_SERVER = typeof window === typeof undefined;
const CACHE_TTL = 5 * 1000; // 5 seconds, to resolve preloaded results

// All relay query will eventually call this fetch with params.text which is a graphql body compiled by relay
export async function networkFetch(params: RequestParameters, variables: Variables): Promise<GraphQLResponse> {
  const accessToken = localStorage.getItem('accessToken');
  let headers: HeadersInit = { 'Content-Type': 'application/json' };
  if (accessToken != null) {
    headers = { ...headers, Authorization: `Bearer ${accessToken}` };
  }

  const response = await fetch(process.env.NEXT_PUBLIC_GRAPHQL_URL, {
    method: 'POST',
    headers: headers,
    body: JSON.stringify({
      query: params.text,
      variables,
    }),
  });
  const json = await response.json();

  if (json && json.errors) {
    throw { details: json.errors as GraphQLError };
  }

  return json;
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

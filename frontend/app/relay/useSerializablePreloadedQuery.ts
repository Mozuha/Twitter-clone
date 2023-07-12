// Reference:
// https://github.com/relayjs/relay-examples/pull/241/files#diff-e7f510749e877b2146eb798e0ab75c0a7db9a6c8483f6127ca0056067782050d

// Convert preloaded query object (with raw GraphQL Response) into
// Relay's PreloadedQuery.

import { useMemo } from 'react';

import { responseCache } from './environment';

import type { SerializablePreloadedQuery } from './loadSerializableQuery';
import type { PreloadedQuery, PreloadFetchPolicy } from 'react-relay';
import type { ConcreteRequest, IEnvironment, OperationType } from 'relay-runtime';

// This hook convert serializable preloaded query
// into Relay's PreloadedQuery object.
// It is also writes this serializable preloaded query
// into QueryResponseCache, so we the network layer
// can use these cache results when fetching data
// in `usePreloadedQuery`.
export default function useSerializablePreloadedQuery<TRequest extends ConcreteRequest, TQuery extends OperationType>(
  environment: IEnvironment,
  preloadQuery: SerializablePreloadedQuery<TRequest, TQuery>,
  fetchPolicy: PreloadFetchPolicy = 'store-or-network'
): PreloadedQuery<TQuery> {
  useMemo(() => {
    writePreloadedQueryToCache(preloadQuery);
  }, [preloadQuery]);

  return {
    environment,
    fetchKey: preloadQuery.params.id ?? preloadQuery.params.cacheID,
    fetchPolicy,
    isDisposed: false,
    name: preloadQuery.params.name,
    kind: 'PreloadedQuery',
    variables: preloadQuery.variables,
    dispose: () => {
      return;
    },
  };
}

function writePreloadedQueryToCache<TRequest extends ConcreteRequest, TQuery extends OperationType>(
  preloadedQueryObject: SerializablePreloadedQuery<TRequest, TQuery>
) {
  const cacheKey = preloadedQueryObject.params.id ?? preloadedQueryObject.params.cacheID;
  responseCache?.set(cacheKey, preloadedQueryObject.variables, preloadedQueryObject.response);
}

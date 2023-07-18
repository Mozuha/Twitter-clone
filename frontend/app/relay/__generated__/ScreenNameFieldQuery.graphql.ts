/**
 * @generated SignedSource<<4749159d084b83f5ce572fd74b177333>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
export type ScreenNameFieldQuery$variables = {
  screenName: string;
};
export type ScreenNameFieldQuery$data = {
  readonly screenNameExists: boolean | null;
};
export type ScreenNameFieldQuery = {
  response: ScreenNameFieldQuery$data;
  variables: ScreenNameFieldQuery$variables;
};

const node: ConcreteRequest = (function () {
  var v0 = [
      {
        defaultValue: null,
        kind: 'LocalArgument',
        name: 'screenName',
      },
    ],
    v1 = [
      {
        alias: null,
        args: [
          {
            kind: 'Variable',
            name: 'screenName',
            variableName: 'screenName',
          },
        ],
        kind: 'ScalarField',
        name: 'screenNameExists',
        storageKey: null,
      },
    ];
  return {
    fragment: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Fragment',
      metadata: null,
      name: 'ScreenNameFieldQuery',
      selections: v1 /*: any*/,
      type: 'Query',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Operation',
      name: 'ScreenNameFieldQuery',
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: '61b390a8b5381018225e6be00ba4034d',
      id: null,
      metadata: {},
      name: 'ScreenNameFieldQuery',
      operationKind: 'query',
      text: 'query ScreenNameFieldQuery(\n  $screenName: String!\n) {\n  screenNameExists(screenName: $screenName)\n}\n',
    },
  };
})();

(node as any).hash = '6606503be03a94aa6997fd43640835eb';

export default node;

/**
 * @generated SignedSource<<cbb05a8510439b99c98b82cf9c533959>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ClientRequest, ClientQuery } from 'relay-runtime';
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

const node: ClientRequest = (function () {
  var v0 = [
      {
        defaultValue: null,
        kind: 'LocalArgument',
        name: 'screenName',
      },
    ],
    v1 = [
      {
        kind: 'ClientExtension',
        selections: [
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
        ],
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
      cacheID: '5f470779f36b396e6aeec715bdc579a0',
      id: null,
      metadata: {},
      name: 'ScreenNameFieldQuery',
      operationKind: 'query',
      text: 'query screennameFieldQuery(\n  $screenName: String!\n) {\n  screenNameExists(screenName: $screenName)\n}\n',
    },
  };
})();

(node as any).hash = '6606503be03a94aa6997fd43640835eb';

export default node;

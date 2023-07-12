/**
 * @generated SignedSource<<d706171997b54ba37cd324ee8e0f2d5f>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ClientRequest, ClientQuery } from 'relay-runtime';
export type screennameFieldQuery$variables = {
  screenName: string;
};
export type screennameFieldQuery$data = {
  readonly screenNameExists: boolean | null;
};
export type screennameFieldQuery = {
  response: screennameFieldQuery$data;
  variables: screennameFieldQuery$variables;
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
      name: 'screennameFieldQuery',
      selections: v1 /*: any*/,
      type: 'Query',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Operation',
      name: 'screennameFieldQuery',
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: '441a34c23ab712801c816c1837b5b23f',
      id: null,
      metadata: {},
      name: 'screennameFieldQuery',
      operationKind: 'query',
      text: 'query screennameFieldQuery(\n  $screenName: String!\n) {\n  screenNameExists(screenName: $screenName)\n}\n',
    },
  };
})();

(node as any).hash = 'b40f30efb8cc7e2a0d9d518a7dd1f751';

export default node;

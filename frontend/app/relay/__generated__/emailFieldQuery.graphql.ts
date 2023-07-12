/**
 * @generated SignedSource<<dcf42cf9e87b1126717b8a87f54c04dd>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ClientRequest, ClientQuery } from 'relay-runtime';
export type emailFieldQuery$variables = {
  email: string;
};
export type emailFieldQuery$data = {
  readonly emailExists: boolean | null;
};
export type emailFieldQuery = {
  response: emailFieldQuery$data;
  variables: emailFieldQuery$variables;
};

const node: ClientRequest = (function () {
  var v0 = [
      {
        defaultValue: null,
        kind: 'LocalArgument',
        name: 'email',
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
                name: 'email',
                variableName: 'email',
              },
            ],
            kind: 'ScalarField',
            name: 'emailExists',
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
      name: 'emailFieldQuery',
      selections: v1 /*: any*/,
      type: 'Query',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Operation',
      name: 'emailFieldQuery',
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: '3cd359f24b2ee8c91aa34a9afa1c8dd5',
      id: null,
      metadata: {},
      name: 'emailFieldQuery',
      operationKind: 'query',
      text: 'query emailFieldQuery(\n  $email: String!\n) {\n  emailExists(email: $email)\n}\n',
    },
  };
})();

(node as any).hash = 'e4c8917c4c1c3e472fd8f1f74f3ad55f';

export default node;

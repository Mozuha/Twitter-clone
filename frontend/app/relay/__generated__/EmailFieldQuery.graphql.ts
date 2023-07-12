/**
 * @generated SignedSource<<bf3b8524f8372994ebb63db5f9629fae>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ClientRequest, ClientQuery } from 'relay-runtime';
export type EmailFieldQuery$variables = {
  email: string;
};
export type EmailFieldQuery$data = {
  readonly emailExists: boolean | null;
};
export type EmailFieldQuery = {
  response: EmailFieldQuery$data;
  variables: EmailFieldQuery$variables;
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
      name: 'EmailFieldQuery',
      selections: v1 /*: any*/,
      type: 'Query',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Operation',
      name: 'EmailFieldQuery',
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: '8bf572ff2bb584e9a3aa4df417d896b4',
      id: null,
      metadata: {},
      name: 'EmailFieldQuery',
      operationKind: 'query',
      text: 'query emailFieldQuery(\n  $email: String!\n) {\n  emailExists(email: $email)\n}\n',
    },
  };
})();

(node as any).hash = 'e4e2ae42c5e14c221b09a4e055589de6';

export default node;

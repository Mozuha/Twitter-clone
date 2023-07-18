/**
 * @generated SignedSource<<5a0460fc79ce705fdc193070e6f471da>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
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

const node: ConcreteRequest = (function () {
  var v0 = [
      {
        defaultValue: null,
        kind: 'LocalArgument',
        name: 'email',
      },
    ],
    v1 = [
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
      cacheID: '6fa1a90a5a637f8497283e5e9f68993a',
      id: null,
      metadata: {},
      name: 'EmailFieldQuery',
      operationKind: 'query',
      text: 'query EmailFieldQuery(\n  $email: String!\n) {\n  emailExists(email: $email)\n}\n',
    },
  };
})();

(node as any).hash = 'e4e2ae42c5e14c221b09a4e055589de6';

export default node;

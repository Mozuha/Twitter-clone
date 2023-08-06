/**
 * @generated SignedSource<<b24255d227dadd7159dddf00ec5a5d02>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type SigninFormMutation$variables = {
  email: string;
  password: string;
};
export type SigninFormMutation$data = {
  readonly signin: {
    readonly accessToken: string;
    readonly refreshToken: string;
    readonly userId: string;
  };
};
export type SigninFormMutation = {
  response: SigninFormMutation$data;
  variables: SigninFormMutation$variables;
};

const node: ConcreteRequest = (function () {
  var v0 = [
      {
        defaultValue: null,
        kind: 'LocalArgument',
        name: 'email',
      },
      {
        defaultValue: null,
        kind: 'LocalArgument',
        name: 'password',
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
          {
            kind: 'Variable',
            name: 'password',
            variableName: 'password',
          },
        ],
        concreteType: 'SigninResponse',
        kind: 'LinkedField',
        name: 'signin',
        plural: false,
        selections: [
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'userId',
            storageKey: null,
          },
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'accessToken',
            storageKey: null,
          },
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'refreshToken',
            storageKey: null,
          },
        ],
        storageKey: null,
      },
    ];
  return {
    fragment: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Fragment',
      metadata: null,
      name: 'SigninFormMutation',
      selections: v1 /*: any*/,
      type: 'Mutation',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Operation',
      name: 'SigninFormMutation',
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: '998f7ee910323fa70ae8fff97589aa69',
      id: null,
      metadata: {},
      name: 'SigninFormMutation',
      operationKind: 'mutation',
      text: 'mutation SigninFormMutation(\n  $email: String!\n  $password: String!\n) {\n  signin(email: $email, password: $password) {\n    userId\n    accessToken\n    refreshToken\n  }\n}\n',
    },
  };
})();

(node as any).hash = 'a97787adc6d8c012ac0833b87dc42366';

export default node;

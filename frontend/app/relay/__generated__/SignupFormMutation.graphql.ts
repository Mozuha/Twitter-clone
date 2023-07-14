/**
 * @generated SignedSource<<f6585ca1de88ede6f5ffa7d5d4fdca07>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type SignupFormMutation$variables = {
  email: string;
  name: string;
  password: string;
  screenName: string;
};
export type SignupFormMutation$data = {
  readonly createUser: {
    readonly id: string;
  };
};
export type SignupFormMutation = {
  response: SignupFormMutation$data;
  variables: SignupFormMutation$variables;
};

const node: ConcreteRequest = (function () {
  var v0 = {
      defaultValue: null,
      kind: 'LocalArgument',
      name: 'email',
    },
    v1 = {
      defaultValue: null,
      kind: 'LocalArgument',
      name: 'name',
    },
    v2 = {
      defaultValue: null,
      kind: 'LocalArgument',
      name: 'password',
    },
    v3 = {
      defaultValue: null,
      kind: 'LocalArgument',
      name: 'screenName',
    },
    v4 = [
      {
        alias: null,
        args: [
          {
            fields: [
              {
                kind: 'Variable',
                name: 'email',
                variableName: 'email',
              },
              {
                kind: 'Variable',
                name: 'name',
                variableName: 'name',
              },
              {
                kind: 'Variable',
                name: 'password',
                variableName: 'password',
              },
              {
                kind: 'Variable',
                name: 'screenName',
                variableName: 'screenName',
              },
            ],
            kind: 'ObjectValue',
            name: 'input',
          },
        ],
        concreteType: 'User',
        kind: 'LinkedField',
        name: 'createUser',
        plural: false,
        selections: [
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'id',
            storageKey: null,
          },
        ],
        storageKey: null,
      },
    ];
  return {
    fragment: {
      argumentDefinitions: [v0 /*: any*/, v1 /*: any*/, v2 /*: any*/, v3 /*: any*/],
      kind: 'Fragment',
      metadata: null,
      name: 'SignupFormMutation',
      selections: v4 /*: any*/,
      type: 'Mutation',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: [v1 /*: any*/, v3 /*: any*/, v0 /*: any*/, v2 /*: any*/],
      kind: 'Operation',
      name: 'SignupFormMutation',
      selections: v4 /*: any*/,
    },
    params: {
      cacheID: 'b10402e44baf877e5bd07f15be44d4ca',
      id: null,
      metadata: {},
      name: 'SignupFormMutation',
      operationKind: 'mutation',
      text: 'mutation SignupFormMutation(\n  $name: String!\n  $screenName: String!\n  $email: String!\n  $password: String!\n) {\n  createUser(input: {name: $name, screenName: $screenName, email: $email, password: $password}) {\n    id\n  }\n}\n',
    },
  };
})();

(node as any).hash = '843735477be8d7cdd50c4602509ae064';

export default node;

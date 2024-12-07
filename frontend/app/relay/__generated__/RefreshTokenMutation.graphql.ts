/**
 * @generated SignedSource<<15499c0e047e5f125a094aa8bea239f5>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type RefreshTokenMutation$variables = {
  refreshToken: string;
};
export type RefreshTokenMutation$data = {
  readonly refreshToken: {
    readonly accessToken: string;
  };
};
export type RefreshTokenMutation = {
  response: RefreshTokenMutation$data;
  variables: RefreshTokenMutation$variables;
};

const node: ConcreteRequest = (function () {
  var v0 = [
      {
        defaultValue: null,
        kind: 'LocalArgument',
        name: 'refreshToken',
      },
    ],
    v1 = [
      {
        alias: null,
        args: [
          {
            kind: 'Variable',
            name: 'refreshToken',
            variableName: 'refreshToken',
          },
        ],
        concreteType: 'RefreshTokenResponse',
        kind: 'LinkedField',
        name: 'refreshToken',
        plural: false,
        selections: [
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'accessToken',
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
      name: 'RefreshTokenMutation',
      selections: v1 /*: any*/,
      type: 'Mutation',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Operation',
      name: 'RefreshTokenMutation',
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: '6bff5164e6b3f47852039d6a78a8937e',
      id: null,
      metadata: {},
      name: 'RefreshTokenMutation',
      operationKind: 'mutation',
      text: 'mutation RefreshTokenMutation(\n  $refreshToken: String!\n) {\n  refreshToken(refreshToken: $refreshToken) {\n    accessToken\n  }\n}\n',
    },
  };
})();

(node as any).hash = '9f41d27601e57d97c3a78ef71d57eb77';

export default node;

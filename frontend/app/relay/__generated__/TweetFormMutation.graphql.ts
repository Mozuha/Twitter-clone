/**
 * @generated SignedSource<<ceb2b7776d1ca21fe980ddac5cb1fc1f>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Mutation } from 'relay-runtime';
export type TweetFormMutation$variables = {
  postedByID: string;
  text: string;
};
export type TweetFormMutation$data = {
  readonly createTweet: {
    readonly id: string;
  };
};
export type TweetFormMutation = {
  response: TweetFormMutation$data;
  variables: TweetFormMutation$variables;
};

const node: ConcreteRequest = (function () {
  var v0 = {
      defaultValue: null,
      kind: 'LocalArgument',
      name: 'postedByID',
    },
    v1 = {
      defaultValue: null,
      kind: 'LocalArgument',
      name: 'text',
    },
    v2 = [
      {
        alias: null,
        args: [
          {
            fields: [
              {
                kind: 'Variable',
                name: 'postedByID',
                variableName: 'postedByID',
              },
              {
                kind: 'Variable',
                name: 'text',
                variableName: 'text',
              },
            ],
            kind: 'ObjectValue',
            name: 'input',
          },
        ],
        concreteType: 'Tweet',
        kind: 'LinkedField',
        name: 'createTweet',
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
      argumentDefinitions: [v0 /*: any*/, v1 /*: any*/],
      kind: 'Fragment',
      metadata: null,
      name: 'TweetFormMutation',
      selections: v2 /*: any*/,
      type: 'Mutation',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: [v1 /*: any*/, v0 /*: any*/],
      kind: 'Operation',
      name: 'TweetFormMutation',
      selections: v2 /*: any*/,
    },
    params: {
      cacheID: '69f65503bf5c823b04f0abbcfbb5be4d',
      id: null,
      metadata: {},
      name: 'TweetFormMutation',
      operationKind: 'mutation',
      text: 'mutation TweetFormMutation(\n  $text: String!\n  $postedByID: ID!\n) {\n  createTweet(input: {text: $text, postedByID: $postedByID}) {\n    id\n  }\n}\n',
    },
  };
})();

(node as any).hash = '0be563752979850441bf3c15e1b5e8f6';

export default node;

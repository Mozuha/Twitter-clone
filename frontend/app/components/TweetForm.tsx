'use client';

import { useState } from 'react';

import { useRouter } from 'next/navigation';

import { graphql, useMutation } from 'react-relay';

import { Button, Spinner } from '@components/material-tailwind';

import { getCookieAction } from '@actions/cookieActions';

import { ErrorMessages, type GraphQLError } from '@lib/constants';

import type { TweetFormMutation } from '@relay/__generated__/TweetFormMutation.graphql';

interface TweetFormProps {
  parentId?: string;
}

export const tweetMutation = graphql`
  mutation TweetFormMutation($text: String!, $postedByID: ID!) {
    createTweet(input: { text: $text, postedByID: $postedByID }) {
      id
    }
  }
`;

export default function TweetForm({ parentId }: TweetFormProps) {
  const [text, setText] = useState('');
  const [error, setError] = useState('');
  const router = useRouter();
  const [commitMutation, isMutationInFlight] = useMutation<TweetFormMutation>(tweetMutation);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    const userId = await getCookieAction('userId');

    commitMutation({
      variables: {
        text: text,
        postedByID: userId!,
      },
      onCompleted(res) {
        // TODO: Pop toast? router.refresh()?
        setText('');
      },
      onError(err) {
        console.log(err);
        const gerr = err as GraphQLError;

        setError(gerr.extensions?.userMessage ?? gerr.message);

        if (gerr.message === ErrorMessages.RefreshTokenNotFound || gerr.extensions?.code === 'UNAUTHORIZED') {
          router.push('/signin');
        }
      },
    });
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <textarea
        value={text}
        onChange={(e) => setText(e.target.value)}
        maxLength={140}
        className="w-full p-2 rounded-md resize-none bg-twitter-black text-white placeholder-twitter-grey focus:outline-none"
        placeholder="What's happening?"
      />
      <div className="flex justify-between items-center">
        <span className="text-sm text-twitter-grey">{text.length}/140</span>
        <Button
          variant="filled"
          size="sm"
          type="submit"
          className="bg-twitter-blue rounded-full px-4 py-2 normal-case text-[13px]"
          ripple={false}
          disabled={!text}
        >
          Post
          {isMutationInFlight && <Spinner className="text-twitter-blue/10" />}
        </Button>
        {/* <button
          type="submit"
          className="px-4 py-2 bg-twitter-blue text-white rounded-full font-bold hover:opacity-[0.85] hover:shadow-none"
        >
          Post
        </button> */}
      </div>
      {error && <p className="text-red-500 text-sm">{error}</p>}
    </form>
  );
}

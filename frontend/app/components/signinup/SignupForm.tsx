import { useState } from 'react';

import Link from 'next/link';
import { useRouter } from 'next/navigation';

import { useForm } from 'react-hook-form';
import { graphql, useMutation } from 'react-relay';

import EmailField from '@components/field/EmailField';
import NameField from '@components/field/NameField';
import PasswordField from '@components/field/PasswordField';
import ScreenNameField from '@components/field/ScreenNameField';
import { Button, Spinner } from '@components/material-tailwind';

import { emailRegex } from '@types-constants/form';
import type { FormData, GraphQLError } from '@types-constants/form';

import type { SignupFormMutation } from '@relay/__generated__/SignupFormMutation.graphql';

import type { SubmitHandler } from 'react-hook-form';

export const createUserMutation = graphql`
  mutation SignupFormMutation($name: String!, $screenName: String!, $email: String!, $password: String!) {
    createUser(input: { name: $name, screenName: $screenName, email: $email, password: $password }) {
      id
    }
  }
`;

export default function SignupForm() {
  const router = useRouter();
  const [isSubmitErr, setIsSubmitErr] = useState(false);

  const {
    handleSubmit,
    control,
    formState: { isValid },
  } = useForm<FormData>({
    mode: 'onChange',
    criteriaMode: 'all',
    defaultValues: { name: '', screenName: '', email: '', password: '' },
  });

  const [commitMutation, isMutationInFlight] = useMutation<SignupFormMutation>(createUserMutation);

  const onSubmit: SubmitHandler<FormData> = (data) => {
    commitMutation({
      variables: {
        name: data.name,
        screenName: data.screenName,
        email: data.email,
        password: data.password,
      },
      onCompleted(res) {
        localStorage.setItem('userId', res.createUser.id);
        router.push('/signin');
      },
      onError(err) {
        console.log((err as GraphQLError).details);
        setIsSubmitErr(true);
      },
    });
  };

  // TODO: It is better to disable fields upon submitting the form, but the response arrives quite fast
  //       so setting disabled props will just make fields flickered. Remove them?
  return (
    <>
      <form className="w-11/12 mt-7" onSubmit={handleSubmit(onSubmit)} aria-label="signup-form">
        <div className="flex flex-col gap-6">
          <NameField
            control={control}
            name="name"
            rules={{
              required: "What's your name?",
              maxLength: { value: 50, message: 'Name must be less than 50 characters.' },
            }}
            // disabled={isMutationInFlight}
          />
          <ScreenNameField
            control={control}
            name="screenName"
            rules={{
              required: 'Screen name is required.',
              maxLength: { value: 15, message: 'Screen name must be less than 15 characters.' },
            }}
            checkExistenceOnBlur
            toggleAlert
            // disabled={isMutationInFlight}
          />
          <EmailField
            control={control}
            name="email"
            rules={{
              required: 'Email is required.',
              pattern: { value: emailRegex, message: 'Please enter a valid email.' },
            }}
            checkExistenceOnBlur
            toggleAlert
            // disabled={isMutationInFlight}
          />
          <PasswordField
            control={control}
            name="password"
            rules={{
              required: 'Password is required.',
              minLength: { value: 8, message: 'Password must be more than 8 characters.' },
              maxLength: { value: 50, message: 'Password must be less than 50 characters.' },
            }}
            showTooltip
            // disabled={isMutationInFlight}
          />
          {isMutationInFlight ? (
            <div className="m-auto">
              <Spinner className="text-twitter-blue/10" />
            </div>
          ) : (
            <Button
              variant="filled"
              size="sm"
              type="submit"
              className="bg-twitter-blue rounded-full mt-2 normal-case text-[13px]"
              ripple={false}
              disabled={!isValid}
            >
              Sign up
            </Button>
          )}
          {isSubmitErr && (
            <span role="alert" className="text-xs font-light text-red-500 -mt-5">
              Something went wrong on signing up. Please try again later.
            </span>
          )}
        </div>
      </form>
      <p className="font-medium text-sm text-twitter-grey mt-4 ml-4 self-start">
        Have an account already?{' '}
        <Link href="/signin" className="text-twitter-blue">
          Sign in
        </Link>
      </p>
    </>
  );
}

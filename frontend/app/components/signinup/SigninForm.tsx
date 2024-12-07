import Link from 'next/link';
import { useRouter } from 'next/navigation';

import { useForm } from 'react-hook-form';
import { graphql, useMutation } from 'react-relay';

import EmailField from '@components/field/EmailField';
import PasswordField from '@components/field/PasswordField';
import { Button, Spinner } from '@components/material-tailwind';

import { setCookieAction } from '@actions/cookieActions';

import { EmailRegex } from '@lib/constants';
import type { AuthFormData, GraphQLError } from '@lib/constants';

import type { SigninFormMutation } from '@relay/__generated__/SigninFormMutation.graphql';

import type { SubmitHandler } from 'react-hook-form';

export const signinMutation = graphql`
  mutation SigninFormMutation($email: String!, $password: String!) {
    signin(email: $email, password: $password) {
      userId
      accessToken
      refreshToken
    }
  }
`;

export default function SigninForm() {
  const router = useRouter();

  const {
    handleSubmit,
    control,
    formState: { isValid },
    setError,
  } = useForm<AuthFormData>({
    mode: 'onChange',
    criteriaMode: 'all',
    defaultValues: { email: '', password: '' },
  });

  const [commitMutation, isMutationInFlight] = useMutation<SigninFormMutation>(signinMutation);

  const onSubmit: SubmitHandler<AuthFormData> = (data) => {
    commitMutation({
      variables: {
        email: data.email,
        password: data.password,
      },
      onCompleted(res) {
        Promise.all([
          setCookieAction('userId', res.signin.userId),
          setCookieAction('accessToken', res.signin.accessToken),
          setCookieAction('refreshToken', res.signin.refreshToken),
        ]).then(() => {
          router.push('/home');
        });
      },
      onError(err) {
        console.log(err);
        const gerr = err as GraphQLError;

        if (gerr.extensions.code === 'NOT_FOUND') {
          setError('email', {
            types: {
              not_found: 'This email address is not registered',
            },
          });
        }

        if (gerr.extensions.code === 'UNAUTHORIZED') {
          setError('password', {
            types: {
              unauthorized: 'Password incorrect',
            },
          });
        }
      },
    });
  };

  return (
    <>
      <form className="w-11/12 mt-7" onSubmit={handleSubmit(onSubmit)} aria-label="signin-form">
        <div className="flex flex-col gap-6">
          <EmailField
            control={control}
            name="email"
            rules={{
              required: 'Email is required.',
              pattern: { value: EmailRegex, message: 'Please enter a valid email.' },
            }}
            checkExistenceOnBlur={false}
            toggleAlert={false}
          />
          <PasswordField
            control={control}
            name="password"
            rules={{
              required: 'Password is required.',
            }}
            showTooltip={false}
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
              Sign in
            </Button>
          )}
        </div>
      </form>
      <p className="font-medium text-sm text-twitter-grey mt-4 ml-4 self-start">
        Don&#39;t have an account?{' '}
        <Link href="/signup" className="text-twitter-blue">
          Sign up
        </Link>
      </p>
    </>
  );
}

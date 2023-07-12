import { useForm } from 'react-hook-form';

import EmailField from '@components/field/email-field';
import NameField from '@components/field/name-field';
import PasswordField from '@components/field/password-field';
import ScreenNameField from '@components/field/screenname-field';
import { Button } from '@components/material-tailwind';

import { emailRegex } from '@types-constants/form';
import type { FormData } from '@types-constants/form';

import type { SubmitHandler } from 'react-hook-form';

export default function SignupForm() {
  const {
    handleSubmit,
    control,
    formState: { isValid },
  } = useForm<FormData>({
    mode: 'onChange',
    criteriaMode: 'all',
    defaultValues: { name: '', screenName: '', email: '', password: '' },
  });

  const onSubmit: SubmitHandler<FormData> = (data) => console.log(data);

  return (
    <>
      <form className="w-11/12 mt-7" onSubmit={handleSubmit(onSubmit)}>
        <div className="flex flex-col gap-6">
          <NameField
            control={control}
            name="name"
            rules={{
              required: "What's your name?",
              maxLength: { value: 50, message: 'Name must be less than 50 characters.' },
            }}
          />
          <ScreenNameField
            control={control}
            name="screenName"
            rules={{
              required: 'Screen name is required.',
              maxLength: { value: 15, message: 'Screen name must be less than 15 characters.' },
            }}
          />
          <EmailField
            control={control}
            name="email"
            rules={{
              required: 'Email is required.',
              pattern: { value: emailRegex, message: 'Please enter a valid email.' },
            }}
          />
          <PasswordField
            control={control}
            name="password"
            rules={{
              required: 'Password is required.',
              minLength: { value: 8, message: 'Password must be more than 8 characters.' },
              maxLength: { value: 50, message: 'Password must be less than 50 characters.' },
            }}
          />
          <Button
            variant="filled"
            size="sm"
            type="submit"
            className="bg-twitter-blue rounded-full mt-2 normal-case text-[13px]"
            disabled={!isValid}
          >
            Sign up
          </Button>
        </div>
      </form>
      <p className="font-medium text-sm mt-4 ml-4 self-start">
        Have an account already? <a className="text-twitter-blue">Sign in</a>
      </p>
    </>
  );
}

'use client';
import { useState } from 'react';

import { ErrorMessage } from '@hookform/error-message';
import { Controller, useForm } from 'react-hook-form';

import { HiOutlineInformationCircle } from 'react-icons/hi';
import { VscEye, VscEyeClosed } from 'react-icons/vsc';

import { Button, Input } from '@components/material-tailwind';

import type { SubmitHandler } from 'react-hook-form';

const emailRegex =
  // eslint-disable-next-line
  /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;

type FormData = {
  name: string;
  screenName: string;
  email: string;
  password: string;
};

/**
 *
 * TODO: Split parts that depends on useState and watch into separate components to avoid unnecessary re-renders.
 * https://zenn.dev/takepepe/articles/rhf-usewatach
 *
 * TODO?: Make min-width work
 */

export default function SignupForm() {
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);
  const [isNameFocused, setIsNameFocused] = useState(false);
  const [isScreenNameFocused, setIsScreenNameFocused] = useState(false);

  const {
    handleSubmit,
    control,
    formState: { errors, isValid },
    watch,
  } = useForm<FormData>({
    mode: 'onChange',
    criteriaMode: 'all',
    defaultValues: { name: '', screenName: '', email: '', password: '' },
  });

  const watchName = watch('name');
  const watchScreenName = watch('screenName');

  const togglePasswordVisibility = () => {
    setIsPasswordVisible(!isPasswordVisible);
  };

  const onSubmit: SubmitHandler<FormData> = (data) => console.log(data);

  return (
    <>
      <form className="w-11/12 mt-7" onSubmit={handleSubmit(onSubmit)}>
        <div className="flex flex-col gap-6">
          <Controller
            name="name"
            control={control}
            rules={{
              required: "What's your name?",
              maxLength: { value: 50, message: 'Name must be less than 50 characters.' },
            }}
            render={({ field: { onBlur, ...rest } }) => (
              <div className="relative flex w-full max-w-[24rem]">
                <Input
                  size="lg"
                  label="Name"
                  type="text"
                  maxLength={50}
                  className="text-white text-[15px]"
                  labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
                  error={!!errors.name}
                  onFocus={() => setIsNameFocused(true)}
                  onBlur={() => setIsNameFocused(false)}
                  {...rest}
                />
                {isNameFocused && (
                  <span className="text-xs text-twitter-grey font-normal !absolute right-1 top-1 pr-1 pt-0.5">
                    {watchName.length + ' / 50'}
                  </span>
                )}
              </div>
            )}
          />
          <ErrorMessage
            errors={errors}
            name="name"
            render={({ messages }) =>
              messages
                ? Object.entries(messages).map(([type, message]) => (
                    <span key={type} className="text-xs font-light text-red-500 -mt-5">
                      {message}
                    </span>
                  ))
                : null
            }
          />

          <Controller
            name="screenName"
            control={control}
            rules={{
              required: 'Screen name is required.',
              maxLength: { value: 15, message: 'Screen name must be less than 15 characters.' },
            }}
            render={({ field: { onBlur, ...rest } }) => (
              <div className="relative flex w-full max-w-[24rem]">
                <Input
                  size="lg"
                  label="@Screen_name"
                  type="text"
                  maxLength={15}
                  className="text-white text-[15px]"
                  labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
                  error={!!errors.screenName}
                  onFocus={() => setIsScreenNameFocused(true)}
                  onBlur={() => setIsScreenNameFocused(false)}
                  {...rest}
                />
                {isScreenNameFocused && (
                  <span className="text-xs text-twitter-grey font-normal !absolute right-1 top-1 pr-1 pt-0.5">
                    {watchScreenName.length + ' / 15'}
                  </span>
                )}
              </div>
            )}
          />
          <ErrorMessage
            errors={errors}
            name="screenName"
            render={({ messages }) =>
              messages
                ? Object.entries(messages).map(([type, message]) => (
                    <span key={type} className="text-xs font-light text-red-500 -mt-5">
                      {message}
                    </span>
                  ))
                : null
            }
          />

          <Controller
            name="email"
            control={control}
            rules={{
              required: 'Email is required.',
              pattern: { value: emailRegex, message: 'Please enter a valid email.' },
            }}
            render={({ field }) => (
              <Input
                size="lg"
                label="Email"
                type="email"
                autoComplete="email"
                className="text-white text-[15px]"
                labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
                error={!!errors.email}
                {...field}
              />
            )}
          />
          <ErrorMessage
            errors={errors}
            name="email"
            render={({ messages }) =>
              messages
                ? Object.entries(messages).map(([type, message]) => (
                    <span key={type} className="text-xs font-light text-red-500 -mt-5">
                      {message}
                    </span>
                  ))
                : null
            }
          />

          <Controller
            name="password"
            control={control}
            rules={{
              required: 'Password is required.',
              minLength: { value: 8, message: 'Password must be more than 8 characters.' },
              maxLength: { value: 50, message: 'Password must be less than 50 characters.' },
            }}
            render={({ field }) => (
              <div className="relative flex w-full max-w-[24rem]">
                <Input
                  size="lg"
                  label="Password"
                  type={isPasswordVisible ? 'text' : 'password'}
                  autoComplete="new-password"
                  className="text-white text-[15px] pr-12"
                  labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
                  containerProps={{ className: 'min-w-0' }}
                  error={!!errors.password}
                  {...field}
                />
                <Button
                  variant="text"
                  size="sm"
                  color="gray"
                  ripple={false}
                  className="w-9 h-9 p-0 !absolute right-1 top-1"
                  onClick={togglePasswordVisibility}
                >
                  {isPasswordVisible ? (
                    <VscEyeClosed color="gray" className="w-6 h-6 ml-1.5" />
                  ) : (
                    <VscEye color="gray" className="w-6 h-6 ml-1.5" />
                  )}
                </Button>
                <div className="group round-full !absolute -right-7 top-3">
                  <span className="bg-slate-500 text-white opacity-0 invisible rounded group-hover:visible opacity-100 absolute bottom-7 -right-14 w-80 text-sm p-1">
                    {'Uppercase, Lowercase, Numbers, Symbols (~`!@#$%^&*()_-+={[}]|\\:;"\'<,>.?/) are allowed.'}
                  </span>
                  <HiOutlineInformationCircle size="1.2rem" color="gray" />
                </div>
              </div>
            )}
          />

          <ErrorMessage
            errors={errors}
            name="password"
            render={({ messages }) =>
              messages
                ? Object.entries(messages).map(([type, message]) => (
                    <span key={type} className="text-xs font-light text-red-500 -mt-5">
                      {message}
                    </span>
                  ))
                : null
            }
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

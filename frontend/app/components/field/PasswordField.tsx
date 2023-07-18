import { useState } from 'react';

import { ErrorMessage } from '@hookform/error-message';
import { useController } from 'react-hook-form';

import { HiOutlineInformationCircle } from 'react-icons/hi';
import { VscEye, VscEyeClosed } from 'react-icons/vsc';

import { Button, Input } from '@components/material-tailwind';

import type { FormFieldProps } from '@types-constants/form';

export default function PasswordField(props: FormFieldProps) {
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);
  const {
    field,
    formState: { errors },
  } = useController(props);

  const togglePasswordVisibility = () => {
    setIsPasswordVisible(!isPasswordVisible);
  };

  return (
    <>
      <div className="relative flex w-full">
        <Input
          size="lg"
          label="Password"
          aria-label="Password"
          type={isPasswordVisible ? 'text' : 'password'}
          autoComplete="new-password"
          className="text-white text-[15px] pr-12"
          labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
          containerProps={{ className: 'min-w-0' }}
          error={!!errors.password}
          {...field}
          disabled={props.disabled}
        />
        <Button
          aria-label="PasswordVisibility"
          variant="text"
          size="sm"
          color="gray"
          ripple={false}
          className="w-9 h-9 p-0 !absolute right-1 top-1"
          onClick={togglePasswordVisibility}
          disabled={props.disabled}
        >
          {isPasswordVisible ? (
            <VscEyeClosed color="gray" className="w-6 h-6 ml-1.5" />
          ) : (
            <VscEye color="gray" className="w-6 h-6 ml-1.5" />
          )}
        </Button>
        <div className="group round-full !absolute -right-7 top-3">
          <span className="bg-slate-500 text-white opacity-0 invisible rounded group-hover:visible opacity-95 absolute bottom-7 right-1 w-80 sm:w-[80vw] text-sm p-1">
            {'Uppercase, Lowercase, Numbers, Symbols (~`!@#$%^&*()_-+={[}]|\\:;"\'<,>.?/) are allowed.'}
          </span>
          <HiOutlineInformationCircle size="1.2rem" color="gray" />
        </div>
      </div>
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
    </>
  );
}

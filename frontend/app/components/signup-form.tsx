'use client';
import { useState } from 'react';

import { VscEye, VscEyeClosed } from 'react-icons/vsc';

import { Button, Input } from '@components/material-tailwind';

export default function SignupForm() {
  const [isPasswordVisible, setIsPasswordVisible] = useState(false);

  const togglePasswordVisibility = () => {
    setIsPasswordVisible(!isPasswordVisible);
  };

  return (
    <form className="w-11/12 mt-7">
      <div className="flex flex-col gap-6">
        <Input
          size="lg"
          label="Name"
          type="text"
          maxLength={50}
          className="text-white"
          labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
        />
        <Input
          size="lg"
          label="@Screen_name"
          type="text"
          maxLength={15}
          className="text-white"
          labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
        />
        <Input
          size="lg"
          label="Email"
          type="email"
          className="text-white"
          labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
        />
        <div className="relative flex w-full max-w-[24rem]">
          <Input
            size="lg"
            label="Password"
            type={isPasswordVisible ? 'text' : 'password'}
            minLength={8}
            maxLength={50}
            className="text-white pr-12"
            labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
            containerProps={{ className: 'min-w-0' }}
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
        </div>
        <Button variant="filled" size="sm" className="bg-twitter-blue rounded-full mt-2 normal-case text-[13px]">
          Sign up
        </Button>
        <p className="font-medium text-sm -mt-1.5">
          Have an account already? <a className="text-twitter-blue">Sign in</a>
        </p>
      </div>
    </form>
  );
}

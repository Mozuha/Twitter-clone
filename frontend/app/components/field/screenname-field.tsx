import { useState } from 'react';

import { ErrorMessage } from '@hookform/error-message';
import { useController, type UseControllerProps, useWatch } from 'react-hook-form';

import { Input } from '@components/material-tailwind';

import type { FormData } from '@types-constants/form';

export default function ScreenNameField(props: UseControllerProps<FormData>) {
  const [isScreenNameFocused, setIsScreenNameFocused] = useState(false);
  const {
    field,
    formState: { errors },
  } = useController(props);
  const { onBlur: _, ...rest } = field;
  const watchScreenName = useWatch({ control: props.control, name: props.name });

  return (
    <>
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
    </>
  );
}

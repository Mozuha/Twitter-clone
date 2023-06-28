import { ErrorMessage } from '@hookform/error-message';
import { useController, type UseControllerProps } from 'react-hook-form';

import { Input } from '@components/material-tailwind';

import type { FormData } from '@types-constants/form';

export default function EmailField(props: UseControllerProps<FormData>) {
  const {
    field,
    formState: { errors },
  } = useController(props);

  return (
    <>
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
    </>
  );
}

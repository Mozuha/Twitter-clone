import { Suspense, useState } from 'react';

import { ErrorMessage } from '@hookform/error-message';
import { useController, useWatch } from 'react-hook-form';
import { graphql, useQueryLoader } from 'react-relay';

import EmailNotExistMsg from '@components/EmailNotExistMsg';
import { Input } from '@components/material-tailwind';

import type { FormFieldProps } from '@types-constants/form';

import type { EmailFieldQuery } from '@relay/__generated__/EmailFieldQuery.graphql';

export const emailExistsQuery = graphql`
  query EmailFieldQuery($email: String!) {
    emailExists(email: $email)
  }
`;

export default function EmailField(props: FormFieldProps) {
  const [isEmailFocused, setIsEmailFocused] = useState(false);
  const [prevWatch, setPrevWatch] = useState('');

  const {
    field,
    formState: { errors },
  } = useController(props);
  const { onBlur: _, ...rest } = field;
  const emailWatch = useWatch({ control: props.control, name: props.name });

  const [queryRef, loadQuery] = useQueryLoader<EmailFieldQuery>(emailExistsQuery);

  const handleBlur = () => {
    if (prevWatch !== emailWatch) {
      emailWatch.length && !errors.email && loadQuery({ email: emailWatch }, { fetchPolicy: 'network-only' });
      setPrevWatch(emailWatch);
    }
    setIsEmailFocused(false);
  };

  return (
    <>
      <Input
        size="lg"
        label="Email"
        aria-label="Email"
        type="email"
        autoComplete="email"
        className="text-white text-[15px]"
        labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
        error={
          !!errors.email ||
          Boolean(
            queryRef?.environment.getStore().getSource().get('client:root')?.[
              `emailExists(email:"${queryRef.variables.email}")`
            ]
          )
        }
        onFocus={() => setIsEmailFocused(true)}
        onBlur={handleBlur}
        {...rest}
        disabled={props.disabled}
      />
      <ErrorMessage
        errors={errors}
        name="email"
        render={({ messages }) =>
          messages
            ? Object.entries(messages).map(([type, message]) => (
                <span key={type} role="alert" className="text-xs font-light text-red-500 -mt-5">
                  {message}
                </span>
              ))
            : null
        }
      />
      <Suspense>
        {queryRef != null && !isEmailFocused && emailWatch && <EmailNotExistMsg queryRef={queryRef} />}
      </Suspense>
    </>
  );
}

import { Suspense, useState } from 'react';

import { ErrorMessage } from '@hookform/error-message';
import { useController, type UseControllerProps, useWatch } from 'react-hook-form';
import { graphql, useQueryLoader } from 'react-relay';

import EmailNotExistMsg from '@components/email-not-exist-message';
import { Input } from '@components/material-tailwind';

import type { FormData } from '@types-constants/form';

import type { emailFieldQuery } from '@relay/__generated__/emailFieldQuery.graphql';

// Replace params.text in relay/__generated__/screennameFieldQuery.graphql.ts with this if it is null
// 'query emailFieldQuery(\n  $email: String!\n) {\n  emailExists(email: $email)\n}\n'
export const emailExistsQuery = graphql`
  query emailFieldQuery($email: String!) {
    emailExists(email: $email)
  }
`;

export default function EmailField(props: UseControllerProps<FormData>) {
  const [isEmailFocused, setIsEmailFocused] = useState(false);
  const [prevWatch, setPrevWatch] = useState('');

  const {
    field,
    formState: { errors },
  } = useController(props);
  const { onBlur: _, ...rest } = field;
  const emailWatch = useWatch({ control: props.control, name: props.name });

  const [queryRef, loadQuery] = useQueryLoader<emailFieldQuery>(emailExistsQuery);

  const handleBlur = () => {
    if (prevWatch !== emailWatch) {
      emailWatch.length && loadQuery({ email: emailWatch }, { fetchPolicy: 'network-only' });
      setPrevWatch(emailWatch);
    }
    setIsEmailFocused(false);
  };

  return (
    <>
      <Input
        size="lg"
        label="Email"
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
      <Suspense>
        {queryRef != null && !isEmailFocused && emailWatch && <EmailNotExistMsg queryRef={queryRef} />}
      </Suspense>
    </>
  );
}

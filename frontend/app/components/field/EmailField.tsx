import { Suspense, useState } from 'react';

import { ErrorMessage } from '@hookform/error-message';
import { useController, useWatch } from 'react-hook-form';
import { graphql, useQueryLoader } from 'react-relay';

import ExistenceMsg from '@components/ExistenceMsg';
import { Input } from '@components/material-tailwind';

import type { FormFieldProps } from '@types-constants/form';

import type { EmailFieldQuery } from '@relay/__generated__/EmailFieldQuery.graphql';

export const emailExistsQuery = graphql`
  query EmailFieldQuery($email: String!) {
    emailExists(email: $email)
  }
`;

type Props = FormFieldProps & {
  toggleAlert: boolean;
  checkExistenceOnBlur: boolean;
};

export default function EmailField(props: Props) {
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
          (props.checkExistenceOnBlur &&
            Boolean(
              queryRef?.environment.getStore().getSource().get('client:root')?.[
                `emailExists(email:"${queryRef.variables.email}")`
              ]
            ) === props.toggleAlert)
        }
        onFocus={() => setIsEmailFocused(true)}
        onBlur={props.checkExistenceOnBlur ? handleBlur : undefined}
        {...rest}
        disabled={props.disabled}
      />
      <ErrorMessage
        errors={errors}
        name="email"
        render={({ messages }) =>
          messages &&
          Object.entries(messages).map(([type, message]) => (
            <span key={type} role="alert" className="text-xs font-light text-red-500 -mt-5">
              {message}
            </span>
          ))
        }
      />
      <Suspense>
        {queryRef != null && !isEmailFocused && emailWatch && !errors.email && (
          <ExistenceMsg queryRef={queryRef} gqlQuery={emailExistsQuery} toggleAlert={props.toggleAlert} />
        )}
      </Suspense>
    </>
  );
}

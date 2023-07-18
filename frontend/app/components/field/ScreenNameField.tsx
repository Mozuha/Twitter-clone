import { Suspense, useState } from 'react';

import { ErrorMessage } from '@hookform/error-message';
import { useController, useWatch } from 'react-hook-form';
import { graphql, useQueryLoader } from 'react-relay';

import { Input } from '@components/material-tailwind';
import ScreenNameNotExistMsg from '@components/ScreenNameNotExistMsg';

import type { FormFieldProps } from '@types-constants/form';

import type { ScreenNameFieldQuery } from '@relay/__generated__/ScreenNameFieldQuery.graphql';

export const screenNameExistsQuery = graphql`
  query ScreenNameFieldQuery($screenName: String!) {
    screenNameExists(screenName: $screenName)
  }
`;

export default function ScreenNameField(props: FormFieldProps) {
  const [isScreenNameFocused, setIsScreenNameFocused] = useState(false);
  const [prevWatch, setPrevWatch] = useState('');

  const {
    field,
    formState: { errors },
  } = useController(props);
  const { onBlur: _, ...rest } = field;
  const screenNameWatch = useWatch({ control: props.control, name: props.name });

  const [queryRef, loadQuery] = useQueryLoader<ScreenNameFieldQuery>(screenNameExistsQuery);

  const handleBlur = () => {
    if (prevWatch !== screenNameWatch) {
      // always send query to check the latest availability of the screen name
      // send query only when screen name is not empty and is changed on blur to reduce the number of queries
      screenNameWatch.length && loadQuery({ screenName: screenNameWatch }, { fetchPolicy: 'network-only' });
      setPrevWatch(screenNameWatch);
    }
    setIsScreenNameFocused(false);
  };

  return (
    <>
      <div className="relative flex w-full">
        <Input
          size="lg"
          label="@Screen_name"
          aria-label="ScreenName"
          type="text"
          maxLength={15}
          className="text-white text-[15px]"
          labelProps={{ className: 'peer-placeholder-shown:text-twitter-grey' }}
          error={
            !!errors.screenName ||
            Boolean(
              queryRef?.environment.getStore().getSource().get('client:root')?.[
                `screenNameExists(screenName:"${queryRef.variables.screenName}")`
              ]
            )
          }
          onFocus={() => setIsScreenNameFocused(true)}
          onBlur={handleBlur}
          {...rest}
          disabled={props.disabled}
        />
        {isScreenNameFocused && (
          <span className="text-xs text-twitter-grey font-normal !absolute right-1 top-1 pr-1 pt-0.5">
            {screenNameWatch.length + ' / 15'}
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
      <Suspense>
        {queryRef != null && !isScreenNameFocused && screenNameWatch && <ScreenNameNotExistMsg queryRef={queryRef} />}
      </Suspense>
    </>
  );
}

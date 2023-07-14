import { usePreloadedQuery } from 'react-relay';

import { screenNameExistsQuery } from '@components/field/ScreenNameField';

import type { ScreenNameFieldQuery } from '@relay/__generated__/ScreenNameFieldQuery.graphql';

import type { PreloadedQuery } from 'react-relay';

type Props = {
  queryRef: PreloadedQuery<ScreenNameFieldQuery>;
};

export default function ScreenNameNotExistMsg(props: Props) {
  const res = usePreloadedQuery(screenNameExistsQuery, props.queryRef);

  return res.screenNameExists ? (
    <span className="text-xs font-light text-red-500 -mt-5">
      {props.queryRef.variables.screenName} is already taken
    </span>
  ) : null;
}

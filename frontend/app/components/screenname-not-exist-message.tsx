import { usePreloadedQuery } from 'react-relay';

import { screenNameExistsQuery } from '@components/field/screenname-field';

import type { screennameFieldQuery } from '@relay/__generated__/screennameFieldQuery.graphql';

import type { PreloadedQuery } from 'react-relay';

type Props = {
  queryRef: PreloadedQuery<screennameFieldQuery>;
};

export default function ScreenNameNotExistMsg(props: Props) {
  const res = usePreloadedQuery(screenNameExistsQuery, props.queryRef);

  return res.screenNameExists ? (
    <span className="text-xs font-light text-red-500 -mt-5">
      {props.queryRef.variables.screenName} is already taken
    </span>
  ) : null;
}

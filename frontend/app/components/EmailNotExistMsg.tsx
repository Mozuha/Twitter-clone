import { usePreloadedQuery } from 'react-relay';

import { emailExistsQuery } from '@components/field/EmailField';

import type { EmailFieldQuery } from '@relay/__generated__/EmailFieldQuery.graphql';

import type { PreloadedQuery } from 'react-relay';

type Props = {
  queryRef: PreloadedQuery<EmailFieldQuery>;
};

export default function EmailNotExistMsg(props: Props) {
  const res = usePreloadedQuery(emailExistsQuery, props.queryRef);

  return res.emailExists ? (
    <span className="text-xs font-light text-red-500 -mt-5">This email address is already registered</span>
  ) : null;
}

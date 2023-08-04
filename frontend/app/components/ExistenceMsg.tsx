import { usePreloadedQuery } from 'react-relay';

import type { GraphQLTaggedNode, PreloadedQuery } from 'react-relay';

type Props = {
  queryRef: PreloadedQuery<any>;
  gqlQuery: GraphQLTaggedNode;
  toggleAlert: boolean; // true: xxx exists, false: xxx does not exists
};

export default function ExistenceMsg(props: Props) {
  const res = usePreloadedQuery(props.gqlQuery, props.queryRef);

  if (props.toggleAlert) {
    if (res.screenNameExists) {
      return (
        <span role="alert" className="text-xs font-light text-red-500 -mt-5">
          {props.queryRef.variables.screenName} is already taken
        </span>
      );
    } else if (res.emailExists) {
      return (
        <span role="alert" className="text-xs font-light text-red-500 -mt-5">
          This email address is already registered
        </span>
      );
    } else {
      return null;
    }
  } else {
    if (res.screenNameExists != undefined && !res.screenNameExists) {
      return (
        <span role="alert" className="text-xs font-light text-red-500 -mt-5">
          {props.queryRef.variables.screenName} is not registered
        </span>
      );
    } else if (res.emailExists != undefined && !res.emailExists) {
      return (
        <span role="alert" className="text-xs font-light text-red-500 -mt-5">
          This email address is not registered
        </span>
      );
    } else {
      return null;
    }
  }
}

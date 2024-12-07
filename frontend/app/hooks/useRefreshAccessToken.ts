import { useMutation } from 'react-relay';

import { setCookieAction } from '@actions/cookieActions';

import type { GraphQLError } from '@lib/constants';
import { refreshTokenMutation } from '@lib/RefreshToken';

import type { RefreshTokenMutation } from '@relay/__generated__/RefreshTokenMutation.graphql';

export function useRefreshAccessToken() {
  const [commitMutation] = useMutation<RefreshTokenMutation>(refreshTokenMutation);

  const refreshAccessToken = async (refreshToken: string): Promise<string> => {
    return new Promise((resolve, reject) => {
      commitMutation({
        variables: {
          refreshToken: refreshToken!,
        },
        onCompleted(res) {
          const newAccessToken = res.refreshToken.accessToken;
          setCookieAction('accessToken', newAccessToken);
          resolve(newAccessToken);
        },
        onError(err) {
          reject(err as GraphQLError);
        },
      });
    });
  };

  return refreshAccessToken;
}

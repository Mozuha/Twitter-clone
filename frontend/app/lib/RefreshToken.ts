import { graphql } from 'react-relay';

export const refreshTokenMutation = graphql`
  mutation RefreshTokenMutation($refreshToken: String!) {
    refreshToken(refreshToken: $refreshToken) {
      accessToken
    }
  }
`;

import type { UseControllerProps } from 'react-hook-form';

export type AuthFormData = {
  name: string;
  screenName: string;
  email: string;
  password: string;
};

export type FormFieldProps = UseControllerProps<AuthFormData> & {
  disabled?: boolean;
};

export const EmailRegex =
  // eslint-disable-next-line
  /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;

export type GraphQLError = Error & {
  message: string;
  path: string;
  extensions: {
    code: string;
    userMessage?: string;
  };
};

export const ErrorMessages = {
  RefreshTokenNotFound: 'No refresh token available',
};

import type { UseControllerProps } from 'react-hook-form';

export type FormData = {
  name: string;
  screenName: string;
  email: string;
  password: string;
};

export type FormFieldProps = UseControllerProps<FormData> & {
  disabled?: boolean;
};

export const emailRegex =
  // eslint-disable-next-line
  /^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;

export type GraphQLError = Error & {
  details: [
    {
      message: string;
      path: string;
      extensions: {
        code: string;
        userMessage?: string;
      };
    }
  ];
};

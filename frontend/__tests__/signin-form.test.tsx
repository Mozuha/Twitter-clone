import '@testing-library/jest-dom';

import { useRouter } from 'next/navigation';

import { cleanup, fireEvent, render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { RelayEnvironmentProvider } from 'react-relay';
import { createMockEnvironment, MockPayloadGenerator } from 'relay-test-utils';

import SigninForm, { signinMutation } from '@components/signinup/SigninForm';

import type { GraphQLError } from '@types-constants/form';

import type { RelayMockEnvironment } from 'relay-test-utils/lib/RelayModernMockEnvironment';

jest.mock('next/navigation', () => ({
  useRouter: jest.fn(),
}));

describe('SigninForm', () => {
  let environment: RelayMockEnvironment;

  beforeEach(() => {
    environment = createMockEnvironment();
  });

  afterEach(cleanup);

  const callQueueOperationResolver = () => {
    environment.mock.queueOperationResolver((operation) => {
      return MockPayloadGenerator.generate(operation, {
        SigninResponse() {
          return {
            userId: '123456789',
            accessToken: 'accesstoken',
            refreshToken: 'refreshtoken',
          };
        },
      });
    });
  };

  const renderSigninForm = () => {
    render(
      <RelayEnvironmentProvider environment={environment}>
        <SigninForm />
      </RelayEnvironmentProvider>
    );
  };

  describe('rendering', () => {
    test('should render the form', () => {
      renderSigninForm();

      expect(screen.getByRole('form', { name: 'signin-form' })).toBeInTheDocument();
      expect(screen.getByLabelText('Email')).toBeInTheDocument();
      expect(screen.getByLabelText('Password')).toBeInTheDocument();
      expect(screen.getByRole('button', { name: 'PasswordVisibility' })).toBeInTheDocument();
      expect(screen.getByRole('button', { name: 'Sign in' })).toBeInTheDocument();
      expect(screen.getByText("Don't have an account?")).toBeInTheDocument();
      expect(screen.getByText('Sign up')).toBeInTheDocument();
    });
  });

  describe('email field', () => {
    test('should show error message when the field is empty', async () => {
      renderSigninForm();

      await userEvent.type(screen.getByLabelText('Email'), 'a{backspace}');

      expect(screen.getByRole('alert')).toHaveTextContent('Email is required.');
    });

    test('should show error message when the email is invalid', async () => {
      renderSigninForm();

      await userEvent.type(screen.getByLabelText('Email'), 'invalid');

      expect(screen.getByRole('alert')).toHaveTextContent('Please enter a valid email.');
    });

    test('should show error message when the email is not registered', async () => {
      environment.mock.queuePendingOperation(signinMutation, {
        email: 'unexisting@gmail.com',
        password: 'arbitrarypass',
      });
      renderSigninForm();

      await userEvent.type(screen.getByLabelText('Email'), 'unexisting@gmail.com');
      await userEvent.type(screen.getByLabelText('Password'), 'arbitrarypass');
      await userEvent.click(screen.getByRole('button', { name: 'Sign in' }));

      await waitFor(() =>
        environment.mock.rejectMostRecentOperation({
          details: [
            {
              message: 'email not found',
              path: 'signin',
              extensions: {
                code: 'NOT_FOUND',
              },
            },
          ],
        } as GraphQLError)
      );

      expect(await screen.findByRole('alert')).toHaveTextContent('This email address is not registered');
    });
  });

  describe('password field', () => {
    test('should show error message when the field is empty', async () => {
      renderSigninForm();

      await userEvent.type(screen.getByLabelText('Password'), 'a{backspace}');

      expect(screen.getByRole('alert')).toHaveTextContent('Password is required.');
    });

    test('should make the password visible when the eye icon is clicked', async () => {
      renderSigninForm();

      await userEvent.click(screen.getByRole('button', { name: 'PasswordVisibility' }));

      expect((await screen.findByLabelText('Password')).getAttribute('type')).toBe('text');
    });

    test('should show error message when the password is not correct', async () => {
      environment.mock.queuePendingOperation(signinMutation, {
        email: 'existing@gmail.com',
        password: 'incorrectpass',
      });
      renderSigninForm();

      await userEvent.type(screen.getByLabelText('Email'), 'existing@gmail.com');
      await userEvent.type(screen.getByLabelText('Password'), 'incorrectpass');
      await userEvent.click(screen.getByRole('button', { name: 'Sign in' }));

      await waitFor(() =>
        environment.mock.rejectMostRecentOperation({
          details: [
            {
              message: 'password incorrect',
              path: 'signin',
              extensions: {
                code: 'UNAUTHORIZED',
              },
            },
          ],
        } as GraphQLError)
      );

      expect(await screen.findByRole('alert')).toHaveTextContent('Password incorrect');
    });
  });

  describe('submit button', () => {
    test('should be disabled when the form is empty', () => {
      renderSigninForm();

      expect(screen.getByRole('button', { name: 'Sign in' })).toBeDisabled();
    });

    test('should be disabled when the form is invalid', async () => {
      renderSigninForm();

      await userEvent.type(screen.getByLabelText('Password'), 'a{backspace}');
      await screen.findByRole('alert');

      expect(screen.getByRole('button', { name: 'Sign in' })).toBeDisabled();
    });

    test('should be enabled when the form is valid', async () => {
      renderSigninForm();

      await userEvent.type(screen.getByLabelText('Email'), 'testtest@gmail.com');
      await userEvent.type(screen.getByLabelText('Password'), 'testtest');
      fireEvent.blur(await screen.findByLabelText('Password'));

      expect(await screen.findByRole('button', { name: 'Sign in' })).toBeEnabled();
    });

    test('should redirect to home page when clicked and received positive response', async () => {
      const mockRouterPush = jest.fn();
      (useRouter as jest.Mock).mockReturnValue({
        push: mockRouterPush,
      });

      callQueueOperationResolver();
      renderSigninForm();

      await userEvent.type(screen.getByLabelText('Email'), 'testtest@gmail.com');
      await userEvent.type(screen.getByLabelText('Password'), 'testtest');
      await userEvent.click(screen.getByRole('button', { name: 'Sign in' }));

      expect(mockRouterPush).toHaveBeenCalledWith('/home');
    });
  });

  describe('signup link', () => {
    test('should have correct path jump to', async () => {
      renderSigninForm();

      expect(await screen.findByRole('link', { name: 'Sign up' })).toHaveAttribute('href', '/signup');
    });
  });
});

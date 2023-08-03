import '@testing-library/jest-dom';
import { useRouter } from 'next/navigation';

import { cleanup, fireEvent, render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { RelayEnvironmentProvider } from 'react-relay';
import { createMockEnvironment, MockPayloadGenerator } from 'relay-test-utils';

import { emailExistsQuery } from '@components/field/EmailField';
import { screenNameExistsQuery } from '@components/field/ScreenNameField';
import SignupForm from '@components/signup/SignupForm';

import type { RelayMockEnvironment } from 'relay-test-utils/lib/RelayModernMockEnvironment';

// TODO: Consider splitting tests to field granularity

jest.mock('next/navigation', () => ({
  useRouter: jest.fn(),
}));

describe('SignupForm', () => {
  let environment: RelayMockEnvironment;

  beforeEach(() => {
    environment = createMockEnvironment();
  });

  afterEach(cleanup);

  const callQueueOperationResolver = ({ screenNameExists = false, emailExists = false }) => {
    environment.mock.queueOperationResolver((operation) => {
      return MockPayloadGenerator.generate(operation, {
        String(context) {
          if (context.name === 'screenNameExists') return screenNameExists;
          if (context.name === 'emailExists') return emailExists;
        },
        User() {
          return {
            id: '123456789',
          };
        },
      });
    });
  };

  const renderSignupForm = () => {
    render(
      <RelayEnvironmentProvider environment={environment}>
        <SignupForm />
      </RelayEnvironmentProvider>
    );
  };

  describe('rendering', () => {
    test('should render the form', () => {
      renderSignupForm();

      expect(screen.getByRole('form', { name: 'signup-form' })).toBeInTheDocument();
      expect(screen.getByLabelText('Name')).toBeInTheDocument();
      expect(screen.getByLabelText('ScreenName')).toBeInTheDocument();
      expect(screen.getByLabelText('Email')).toBeInTheDocument();
      expect(screen.getByLabelText('Password')).toBeInTheDocument();
      expect(screen.getByRole('button', { name: 'PasswordVisibility' })).toBeInTheDocument();
      expect(screen.getByRole('tooltip', { name: 'allowed-characters-for-password' })).toBeInTheDocument();
      expect(screen.getByRole('button', { name: 'Sign up' })).toBeInTheDocument();
      expect(screen.getByText('Have an account already?')).toBeInTheDocument();
      expect(screen.getByText('Sign in')).toBeInTheDocument();
    });
  });

  describe('name field', () => {
    test('should show error message when the field is empty', async () => {
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Name'), 'a');
      await userEvent.clear(screen.getByLabelText('Name'));

      expect(screen.getByRole('alert')).toHaveTextContent("What's your name?");
    });

    test('should not exceed 50 characters', async () => {
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Name'), 'a'.repeat(51));

      expect(screen.getByLabelText('Name').getAttribute('value')).toHaveLength(50);
    });
  });

  describe('screen name field', () => {
    test('should show error message when the field is empty', async () => {
      callQueueOperationResolver({});
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('ScreenName'), 'a{backspace}');

      expect(screen.getByRole('alert')).toHaveTextContent('Screen name is required.');
    });

    test('should not exceed 15 characters', async () => {
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('ScreenName'), 'a'.repeat(16));

      expect(screen.getByLabelText('ScreenName').getAttribute('value')).toHaveLength(15);
    });

    test('should show error message when the screen name already exists', async () => {
      callQueueOperationResolver({ screenNameExists: true });
      environment.mock.queuePendingOperation(screenNameExistsQuery, { screenName: 'existingsn' });
      renderSignupForm();

      await userEvent.type(await screen.findByLabelText('ScreenName'), 'existingsn');
      fireEvent.blur(await screen.findByLabelText('ScreenName'));

      expect(await screen.findByRole('alert')).toHaveTextContent('existingsn is already taken');
    });

    // To actually test the behaviour, I need to check with toHaveCalled but I don't know how to do it
    test('should not throw query if input value has not been changed from before', async () => {
      callQueueOperationResolver({ screenNameExists: true });
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('ScreenName'), 'a{backspace}');
      fireEvent.blur(await screen.findByLabelText('ScreenName'));

      // expect to not see this text after waiting for 100ms
      await expect(screen.findByText('a is already taken', {}, { timeout: 100 })).rejects.toThrow();
    });

    // TODO: test if query was thrown, or if handleBlur was called (need to change implementation of field to pass handleBlur fn as props)
  });

  describe('email field', () => {
    test('should show error message when the field is empty', async () => {
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Email'), 'a{backspace}');

      expect(screen.getByRole('alert')).toHaveTextContent('Email is required.');
    });

    test('should show error message when the email is invalid', async () => {
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Email'), 'invalid');

      expect(screen.getByRole('alert')).toHaveTextContent('Please enter a valid email.');
    });

    test('should show error message when the email already exists', async () => {
      callQueueOperationResolver({ emailExists: true });
      environment.mock.queuePendingOperation(emailExistsQuery, { email: 'existing@gmail.com' });
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Email'), 'existing@gmail.com');
      fireEvent.blur(await screen.findByLabelText('Email'));

      expect(await screen.findByRole('alert')).toHaveTextContent('This email address is already registered');
    });

    test('should not throw query if input value has not been changed from before', async () => {
      callQueueOperationResolver({ emailExists: true });
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Email'), 'a{backspace}');
      fireEvent.blur(await screen.findByLabelText('Email'));

      await expect(
        screen.findByText('This email address is already registered', {}, { timeout: 100 })
      ).rejects.toThrow();
    });

    test('should not throw query if input value is invalid', async () => {
      callQueueOperationResolver({ emailExists: true });
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Email'), 'invalid');
      fireEvent.blur(await screen.findByLabelText('Email'));

      await expect(
        screen.findByText('This email address is already registered', {}, { timeout: 100 })
      ).rejects.toThrow();
    });
  });

  describe('password field', () => {
    test('should show error message when the field is empty', async () => {
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Password'), 'a{backspace}');

      expect(screen.getByRole('alert')).toHaveTextContent('Password is required.');
    });

    test('should show error message when the password is less than 8 characters', async () => {
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Password'), 'a');

      expect(screen.getByRole('alert')).toHaveTextContent('Password must be more than 8 characters.');
    });

    test('should show error message when the password is more than 50 characters', async () => {
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Password'), 'a'.repeat(51));

      expect(screen.getByRole('alert')).toHaveTextContent('Password must be less than 50 characters.');
    });

    test('should make the password visible when the eye icon is clicked', async () => {
      renderSignupForm();

      await userEvent.click(screen.getByRole('button', { name: 'PasswordVisibility' }));

      expect((await screen.findByLabelText('Password')).getAttribute('type')).toBe('text');
    });
  });

  describe('submit button', () => {
    test('should be disabled when the form is empty', () => {
      renderSignupForm();

      expect(screen.getByRole('button', { name: 'Sign up' })).toBeDisabled();
    });

    test('should be disabled when the form is invalid', async () => {
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Password'), 'a');
      await screen.findByRole('alert');

      expect(screen.getByRole('button', { name: 'Sign up' })).toBeDisabled();
    });

    test('should be enabled when the form is valid', async () => {
      callQueueOperationResolver({});
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Name'), 'testtest');
      await userEvent.type(screen.getByLabelText('ScreenName'), 'testtest');
      await userEvent.type(screen.getByLabelText('Email'), 'testtest@gmail.com');
      await userEvent.type(screen.getByLabelText('Password'), 'testtest');
      fireEvent.blur(await screen.findByLabelText('Password'));

      expect(await screen.findByRole('button', { name: 'Sign up' })).toBeEnabled();
    });

    test('should redirect to signin page when clicked and received positive response', async () => {
      const mockRouterPush = jest.fn();
      (useRouter as jest.Mock).mockReturnValue({
        push: mockRouterPush,
      });

      callQueueOperationResolver({});
      renderSignupForm();

      await userEvent.type(screen.getByLabelText('Name'), 'testtest');
      await userEvent.type(screen.getByLabelText('ScreenName'), 'testtest');
      await userEvent.type(screen.getByLabelText('Email'), 'testtest@gmail.com');
      await userEvent.type(screen.getByLabelText('Password'), 'testtest');
      await userEvent.click(screen.getByRole('button', { name: 'Sign up' }));

      expect(mockRouterPush).toHaveBeenCalledWith('/signin');
    });
  });
});

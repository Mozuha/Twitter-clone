'use client';

import { Suspense } from 'react';

import SignupForm from '@/app/components/signup/SignupForm';

import { Spinner } from '@components/material-tailwind';
import SigninupCard from '@components/signup/SigninupCard';

import { getCurrentEnvironment } from '@relay/environment';

const { RelayEnvironmentProvider } = require('react-relay');

const environment = getCurrentEnvironment();

export default function SignupPage() {
  return (
    <RelayEnvironmentProvider environment={environment}>
      <div className="fixed flex justify-center bg-signin-background w-screen h-full">
        <Suspense fallback={<Spinner className="text-twitter-blue/10" />}>
          <SigninupCard>
            <span className="text-white text-3xl font-semibold">Join Twitter today</span>
            <SignupForm />
          </SigninupCard>
        </Suspense>
      </div>
    </RelayEnvironmentProvider>
  );
}

'use client';

import { Suspense } from 'react';

import { RelayEnvironmentProvider } from 'react-relay';

import { Spinner } from '@components/material-tailwind';
import SigninForm from '@components/signinup/SigninForm';
import SigninupCard from '@components/signinup/SigninupCard';

import { getCurrentEnvironment } from '@relay/environment';

const environment = getCurrentEnvironment();

export default function SigninPage() {
  return (
    <RelayEnvironmentProvider environment={environment}>
      <div className="fixed flex justify-center bg-signin-background w-screen h-full">
        <Suspense
          fallback={
            <div className="m-[50vh_50vw]">
              <Spinner className="text-twitter-blue/10" />
            </div>
          }
        >
          <SigninupCard>
            <span className="text-white text-3xl sm:text-xl font-semibold">Sign in to Twitter</span>
            <SigninForm />
          </SigninupCard>
        </Suspense>
      </div>
    </RelayEnvironmentProvider>
  );
}

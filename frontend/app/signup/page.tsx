import SignupForm from '@/app/components/signup/signup-form';

import SigninupCard from '@components/signup/signinup-card';

export default function SignupPage() {
  return (
    <div className="fixed flex justify-center bg-signin-background w-screen h-full">
      <SigninupCard>
        <span className="text-white text-3xl font-semibold">Join Twitter today</span>
        <SignupForm />
      </SigninupCard>
    </div>
  );
}

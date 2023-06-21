import SigninupCard from '@components/signinup-card';
import SignupForm from '@components/signup-form';

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

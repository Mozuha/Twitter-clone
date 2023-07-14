import { FaTwitter } from 'react-icons/fa';

import CloseButton from '@components/CloseButton';
import { Card } from '@components/material-tailwind';

export default function SigninupCard({ children }: { children: React.ReactNode }) {
  return (
    <Card className="grid grid-rows-[70px_1fr] grid-cols-[280px_1fr] sm:grid-cols-[45vw_1fr] h-[650px] sm:h-full max-w-[80vw] max-h-[90vh] sm:max-h-full min-w-[600px] sm:min-w-full m-auto sm:m-0 bg-twitter-black rounded-3xl sm:rounded-none">
      <div className="row-span-1 col-span-1 z-[2] ml-4 pt-4">
        <CloseButton />
      </div>
      <div className="row-span-1 col-auto">
        <FaTwitter color="white" className="w-9 h-9 mt-9" />
      </div>
      <div className="row-auto col-span-2 h-full overflow-x-auto overflow-y-none">
        <div className="flex flex-col flex-nowrap items-center justify-start mx-[100px] sm:mx-7 mt-10 mb-5">
          {children}
        </div>
      </div>
    </Card>
  );
}

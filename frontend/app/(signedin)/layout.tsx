import { Suspense } from 'react';

import SearchField from '@components/field/SearchField';
import { Spinner } from '@components/material-tailwind';
import Navbar from '@components/navbar/Navbar';

export default function MainLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="grid lg:grid-cols-[10%_90%] grid-cols-[25%_50%_25%]">
      <div className="col-start-1 col-span-1 h-screen">
        <Suspense fallback={<Spinner />}>
          <Navbar />
        </Suspense>
      </div>
      <div className="col-start-2 col-span-1">{children}</div>
      <div className="col-start-3 col-span-1">
        <SearchField />
      </div>
    </div>
  );
}

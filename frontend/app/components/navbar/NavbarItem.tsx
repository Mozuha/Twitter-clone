'use client';
import type { ReactElement } from 'react';
import React from 'react';

import Link from 'next/link';
import { usePathname } from 'next/navigation';

import useBreakpoints from '@hooks/useBreakpoints';

import type { IconType } from 'react-icons';

export type NavbarItemProps = {
  href: string;
  icon: ReactElement<IconType>;
  fillIcon: ReactElement<IconType>;
  label: string;
};

export default function NavbarItem(props: NavbarItemProps) {
  const pathname = usePathname();
  const { isXl } = useBreakpoints();
  const fontClassName = pathname.includes(props.href) ? 'font-bold text-xl' : 'font-normal text-xl';

  return (
    <Link href={props.href} className="flex flex-row items-center py-2 my-2 hover:bg-dark-hover rounded-full">
      <div className="lg:mr-0 mr-5">{pathname.includes(props.href) ? props.fillIcon : props.icon}</div>
      {isXl && <div className={fontClassName}>{props.label}</div>}
    </Link>
  );
}

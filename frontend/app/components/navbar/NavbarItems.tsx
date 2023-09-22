'use client';

import { useEffect, useState } from 'react';

import { AiOutlineSetting, AiTwotoneSetting } from 'react-icons/ai';
import { BiSearch } from 'react-icons/bi';
import { BsPerson, BsPersonFill } from 'react-icons/bs';
import { FiSearch } from 'react-icons/fi';
import { RiHome8Fill, RiHome8Line, RiNotification2Fill, RiNotification2Line } from 'react-icons/ri';

import { Spinner } from '@components/material-tailwind';

import NavbarItem from './NavbarItem';

import type { NavbarItemProps } from './NavbarItem';

export default function NavbarItems() {
  const [navbarItems, setNavbarItems] = useState<NavbarItemProps[]>([]);

  useEffect(() => {
    setNavbarItems([
      {
        href: '/home',
        icon: <RiHome8Line color="white" className="w-7 h-7" aria-label="home-icon" />,
        fillIcon: <RiHome8Fill color="white" className="w-7 h-7" aria-label="home-icon-filled" />,
        label: 'Home',
      },
      {
        href: '/explore',
        icon: <BiSearch color="white" className="w-7 h-7" aria-label="explore-icon" />,
        fillIcon: <FiSearch color="white" className="w-7 h-7" aria-label="explore-icon-filled" />,
        label: 'Explore',
      },
      {
        href: '/notifications',
        icon: <RiNotification2Line color="white" className="w-7 h-7" aria-label="notifications-icon" />,
        fillIcon: <RiNotification2Fill color="white" className="w-7 h-7" aria-label="notifications-icon-filled" />,
        label: 'Notifications',
      },
      {
        // TODO: What if the item is empty? Throw some query to get userId? Would it be better to get it from Relay store?
        href: `/${localStorage.getItem('userId')}`,
        icon: <BsPerson color="white" className="w-7 h-7" aria-label="profile-icon" />,
        fillIcon: <BsPersonFill color="white" className="w-7 h-7" aria-label="profile-icon-filled" />,
        label: 'Profile',
      },
      {
        href: '/settings',
        icon: <AiOutlineSetting color="white" className="w-7 h-7" aria-label="settings-icon" />,
        fillIcon: <AiTwotoneSetting color="white" className="w-7 h-7" aria-label="settings-icon-filled" />,
        label: 'Settings',
      },
    ]);
  }, []);

  // TODO: Ideally, we should replace Navbar component with Spinner, not NavbarItems component.
  //       However, using useEffect in here prevents to do so.
  //       Currently using useEffect for getting userId from localstorage but what if use Relay store instead?
  //       Anyway, we would want to check both localstorage and Relay store, so it would be a better idea to separate the userId fetching logic into
  //       e.g. custom hook or into Navbar component.
  if (navbarItems.length < 5)
    return (
      <div className="m-[50%_50%]">
        <Spinner className="h-8 w-8 text-twitter-blue/10" />
      </div>
    );

  return (
    <nav className="flex flex-col">
      <ul>
        {navbarItems.map((item, idx) => (
          <li key={idx}>
            <NavbarItem href={item.href} icon={item.icon} fillIcon={item.fillIcon} label={item.label} />
          </li>
        ))}
      </ul>
    </nav>
  );
}

'use client';
import Image from 'next/image';
import Link from 'next/link';

import { BsThreeDots } from 'react-icons/bs';
import { FaTwitter } from 'react-icons/fa';
import { RiQuillPenLine } from 'react-icons/ri';

import { Button, List, ListItem, Popover, PopoverContent, PopoverHandler } from '@components/material-tailwind';

import useBreakpoints from '@hooks/useBreakpoints';

import NavbarItems from './NavbarItems';

export default function Navbar() {
  const { isXl } = useBreakpoints();

  return (
    <div className="flex flex-col lg:w-20 w-64 mx-[auto] h-full overflow-y-auto p-2">
      <Link href="/home" className="pt-4 pb-2 lg:mx-[auto]">
        <FaTwitter color="white" className="w-7 h-7" aria-label="twitter-icon" />
      </Link>
      <NavbarItems />
      {!isXl ? (
        <div className="mx-[auto]">
          <Button
            variant="filled"
            size="md"
            className="bg-twitter-blue rounded-[50%] hover:opacity-[0.85] hover:shadow-none inline-flex justify-center items-center p-2"
            ripple={false}
            fullWidth
          >
            <RiQuillPenLine color="white" className="w-7 h-7" aria-label="post-icon" />
          </Button>
        </div>
      ) : (
        <div className="pr-8 py-2">
          <Button
            variant="filled"
            size="lg"
            className="bg-twitter-blue rounded-full normal-case font-bold text-base hover:opacity-[0.85] hover:shadow-none"
            ripple={false}
            fullWidth
          >
            Post
          </Button>
        </div>
      )}

      <Popover offset={10}>
        <PopoverHandler>
          <Button
            size="md"
            className="lg:mx-[auto] mt-[auto] my-4 flex justify-between items-center lg:px-0 p-3 bg-transparent hover:bg-dark-hover active:bg-dark-hover rounded-full normal-case shadow-none hover:shadow-none"
            ripple={false}
          >
            <Image src="/img/favicon2.png" alt="profile-picture" width={40} height={40} className="rounded-[50%]" />
            {isXl && (
              <>
                <div className="text-[15px] font-normal text-left w-3/5">
                  <p>Screen name</p>
                  <p className="text-twitter-grey pt-1">@username</p>
                </div>
                <BsThreeDots color="white" className="w-5 h-5" aria-label="more-icon" />
              </>
            )}
          </Button>
        </PopoverHandler>
        <PopoverContent className="bg-black p-2 border-twitter-border-color rounded-lg">
          <List className="p-0 bg-black">
            <ListItem className="p-2 hover:bg-dark-hover active:bg-dark-hover focus:bg-black" ripple={false}>
              <p className="text-[15px] text-white">Log out @username</p>
            </ListItem>
          </List>
        </PopoverContent>
      </Popover>
    </div>
  );
}

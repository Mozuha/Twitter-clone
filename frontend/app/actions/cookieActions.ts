'use server';

import { cookies } from 'next/headers';

export async function setCookieAction(key: string, value: string) {
  let maxAge;

  if (key === 'accessToken') {
    maxAge = 60 * 60; // 1 hour
  } else if (key === 'refreshToken' || key === 'userId') {
    maxAge = 14 * 24 * 60 * 60; // 14 days
  } else {
    maxAge = 7 * 24 * 60 * 60; // 7 days
  }

  (await cookies()).set(key, value, {
    httpOnly: true,
    secure: true,
    sameSite: 'strict',
    path: '/',
    maxAge: maxAge,
  });
}

export async function getCookieAction(key: string): Promise<string | undefined> {
  const cookieStore = await cookies();
  const val = cookieStore.get(key)?.value;

  return val;
}

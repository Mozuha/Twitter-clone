import { type NextRequest, NextResponse } from 'next/server';

import { getCookieAction } from '@actions/cookieActions';

const publicRoutes = ['/signin', '/signup'];

export async function middleware(req: NextRequest) {
  const path = req.nextUrl.pathname;
  const isProtectedRoute = !publicRoutes.includes(path);

  const user = await getCookieAction('userId');

  if (isProtectedRoute && !user) {
    return NextResponse.redirect(new URL('/signin', req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|favicon.ico|.*\\.(?:svg|png|jpg|jpeg|gif|webp)$).*)'],
};

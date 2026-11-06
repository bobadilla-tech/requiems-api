// Next.js 16 proxy (replaces middleware.ts).
// Fast cookie-only check — no database hit, no network call.
// The real session validation happens per-page via auth.api.getSession().

import { NextRequest, NextResponse } from "next/server";
import { getSessionCookie } from "better-auth/cookies";

export function proxy(request: NextRequest) {
  const sessionCookie = getSessionCookie(request);

  if (!sessionCookie) {
    return NextResponse.redirect(new URL("/signup", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/dashboard"],
};

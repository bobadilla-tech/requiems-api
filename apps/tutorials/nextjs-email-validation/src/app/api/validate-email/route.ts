// Server-side proxy route. The client calls /api/validate-email instead of
// api.requiems.xyz directly so the REQUIEMS_API key never reaches the browser.

import { NextRequest, NextResponse } from "next/server";
import { validateEmail } from "@/lib/requiems";

export async function POST(request: NextRequest) {
  let body: { email?: unknown };

  try {
    body = await request.json();
  } catch {
    return NextResponse.json({ error: "Invalid JSON body" }, { status: 400 });
  }

  const email = body?.email;

  if (!email || typeof email !== "string") {
    return NextResponse.json(
      { error: "email field is required" },
      { status: 422 }
    );
  }

  const result = await validateEmail(email.trim());

  if (!result.ok) {
    return NextResponse.json({ error: result.error }, { status: result.status });
  }

  return NextResponse.json({ data: result.data });
}

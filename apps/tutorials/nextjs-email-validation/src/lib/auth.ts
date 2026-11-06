import { betterAuth } from "better-auth";
import { createAuthMiddleware, APIError } from "better-auth/api";
import Database from "better-sqlite3";
import { validateEmail } from "@/lib/requiems";
import { env } from "@/env";

// Middleware that runs before every sign-in and sign-up request.
// This is where the Requiems API validation happens server-side —
// even if someone bypasses the client-side real-time check.
const emailValidationMiddleware = createAuthMiddleware(async (ctx) => {
  const isAuthRoute =
    ctx.path === "/sign-in/email" || ctx.path === "/sign-up/email";

  if (!isAuthRoute) return;

  const email = ctx.body?.email as string | undefined;
  if (!email) return;

  const result = await validateEmail(email);

  // ── API unreachable ────────────────────────────────────────────────────────
  if (!result.ok) {
    throw new APIError("INTERNAL_SERVER_ERROR", {
      message: "Email validation service unavailable. Please try again.",
    });
  }

  const v = result.data;

  // ── RFC 5322 syntax check ──────────────────────────────────────────────────
  if (!v.syntax_valid) {
    throw new APIError("BAD_REQUEST", {
      message: "Invalid email format. Please check the address and try again.",
    });
  }

  // ── Disposable email check ─────────────────────────────────────────────────
  if (v.disposable) {
    throw new APIError("BAD_REQUEST", {
      message:
        "Disposable email addresses are not allowed. Please use a permanent email.",
    });
  }

  // ── MX record check ────────────────────────────────────────────────────────
  if (!v.mx_valid) {
    const hint = v.suggestion ? ` Did you mean \u2026@${v.suggestion}?` : "";
    throw new APIError("BAD_REQUEST", {
      message: `This email domain can't receive mail.${hint}`,
    });
  }
});

export const auth = betterAuth({
  database: new Database("./sqlite.db"),
  secret: env.BETTER_AUTH_SECRET,
  baseURL: env.BETTER_AUTH_URL,

  emailAndPassword: {
    enabled: true,
  },

  hooks: {
    before: emailValidationMiddleware,
  },
});

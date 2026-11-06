import { createEnv } from "@t3-oss/env-nextjs";
import { z } from "zod";

export const env = createEnv({
  server: {
    // Requiems API key — from the requiems-api-tutorials root .env
    REQUIEMS_API: z.string().min(1),
    // Better Auth secret — generate with: openssl rand -base64 32
    BETTER_AUTH_SECRET: z.string().min(32),
    // Better Auth base URL
    BETTER_AUTH_URL: z.url().default("http://localhost:3000"),
  },
  // No public client vars — the API key must stay server-side
  client: {},
  experimental__runtimeEnv: process.env,
});

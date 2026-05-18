import type { BaseWorkerBindings } from "@requiem/workers-shared";
import { createEnvValidator } from "@requiem/workers-shared";

import { z } from "zod";

export interface WorkerBindings extends BaseWorkerBindings {
  BACKEND_URL: string;
  BACKEND_SECRET: string;
  SENTRY_DSN?: string;
}

const envSchema = z.object({
  BACKEND_URL: z.string().url(),
  BACKEND_SECRET: z.string().min(32),
  ENVIRONMENT: z.enum(["development", "staging", "production"]).default("production"),
  SENTRY_DSN: z.string().optional(),
});

export const validateEnv = createEnvValidator(envSchema);

export type ValidatedEnv = z.infer<typeof envSchema>;

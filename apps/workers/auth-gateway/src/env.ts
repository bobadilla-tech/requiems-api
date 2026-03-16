import { z } from "zod";
import { createEnvValidator } from "@requiem/workers-shared";
import type { BaseWorkerBindings } from "@requiem/workers-shared";

export interface WorkerBindings extends BaseWorkerBindings {
  BACKEND_URL: string;
  BACKEND_SECRET: string;
  CORS_ALLOWED_ORIGINS: string;
}

const envSchema = z.object({
  BACKEND_URL: z.string().url(),
  BACKEND_SECRET: z.string().min(32),
  CORS_ALLOWED_ORIGINS: z.string().min(1, "CORS_ALLOWED_ORIGINS must be a non-empty string"),
  ENVIRONMENT: z.enum(["development", "staging", "production"]).default("production"),
});

export const validateEnv = createEnvValidator(envSchema);

export type ValidatedEnv = z.infer<typeof envSchema>;

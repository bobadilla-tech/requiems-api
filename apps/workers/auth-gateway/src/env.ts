import { z } from "zod";
import { createEnvValidator } from "@requiem/workers-shared";
import type { BaseWorkerBindings } from "@requiem/workers-shared";

export interface WorkerBindings extends BaseWorkerBindings {
  BACKEND_URL: string;
  BACKEND_SECRET: string;
  RATE_LIMITER: DurableObjectNamespace;
}

const envSchema = z.object({
  BACKEND_URL: z.string().url(),
  BACKEND_SECRET: z.string().min(32),
  ENVIRONMENT: z.enum(["development", "staging", "production"]).default("production"),
});

export const validateEnv = createEnvValidator(envSchema);

export type ValidatedEnv = z.infer<typeof envSchema>;

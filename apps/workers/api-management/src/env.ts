import { z } from "zod";
import type { BaseWorkerBindings, Logger } from "@requiem/workers-shared";

export interface WorkerBindings extends BaseWorkerBindings {
  API_MANAGEMENT_API_KEY: string;
  SWAGGER_USERNAME?: string;
  SWAGGER_PASSWORD?: string;
}

export interface WorkerVariables {
  log: Logger;
}

export type WorkerEnv = { Bindings: WorkerBindings; Variables: WorkerVariables };

const envSchema = z.object({
  API_MANAGEMENT_API_KEY: z.string().min(32),
  SWAGGER_USERNAME: z.string().optional(),
  SWAGGER_PASSWORD: z.string().optional(),
  ENVIRONMENT: z.enum(["development", "staging", "production"]).default("production"),
});

export function validateEnv(env: WorkerBindings): WorkerBindings {
  envSchema.parse({
    API_MANAGEMENT_API_KEY: env.API_MANAGEMENT_API_KEY,
    SWAGGER_USERNAME: env.SWAGGER_USERNAME,
    SWAGGER_PASSWORD: env.SWAGGER_PASSWORD,
    ENVIRONMENT: env.ENVIRONMENT,
  });

  return env;
}

export type ValidatedEnv = z.infer<typeof envSchema>;

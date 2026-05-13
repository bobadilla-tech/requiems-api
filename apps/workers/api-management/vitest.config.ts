import path from "node:path";

import { defineConfig } from "vitest/config";

export default defineConfig({
  resolve: {
    alias: {
      "@requiem/workers-shared": path.resolve(__dirname, "../shared/src"),
    },
  },
  test: {
    environment: "edge-runtime",
    setupFiles: ["src/__tests__/setup.ts"],
    include: ["src/**/*.{test,spec}.ts"],
    exclude: ["node_modules", "dist"],
    coverage: {
      provider: "v8",
      all: false,
      allowExternal: true,
      reporter: ["text", "json", "html", "lcov"],
      exclude: ["node_modules/", "dist/", "**/*.config.ts", "**/*.d.ts"],
    },
    reporters: "default",
    pool: "threads",
    poolOptions: {
      threads: {
        minThreads: 1,
      },
    },
    testTimeout: 10000,
    hookTimeout: 10000,
    globals: true,
  },
});

/**
 * Configuration loader.
 *
 * Reads API_BASE_URL and REQUIEMS_API_KEY from the environment.
 * Fails fast with a helpful message if either is missing.
 */

export interface Config {
  baseUrl: string;
  apiKey: string;
  /** How many times each endpoint is exercised to compute timing stats */
  runs: number;
  /** Per-request fetch timeout in milliseconds (default 10 000) */
  requestTimeoutMs: number;
}

let _config: Config | undefined;

export function getConfig(): Config {
  if (_config) return _config;

  const baseUrl = process.env["API_BASE_URL"] ?? "https://requiems.xyz";
  const apiKey = process.env["REQUIEMS_API_KEY"] ?? "";
  const runs = Number(process.env["INTEGRATION_RUNS"] ?? "20");
  const requestTimeoutMs = Number(
    process.env["REQUEST_TIMEOUT_MS"] ?? "8000",
  );

  if (!apiKey) {
    throw new Error(
      [
        "",
        "⛔  REQUIEMS_API_KEY is not set.",
        "",
        "   Copy tests/integration/.env.example to tests/integration/.env",
        "   and fill in your production API key, then re-run the tests.",
        "",
      ].join("\n"),
    );
  }

  _config = { baseUrl, apiKey, runs, requestTimeoutMs };
  return _config;
}

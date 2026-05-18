/**
 * Retry an async function with exponential backoff.
 *
 * @param fn - The async function to retry
 * @param attempts - Max number of attempts (default: 3)
 * @param delayMs - Initial delay in ms, doubles each retry (default: 100)
 */
export async function withRetry<T>(fn: () => Promise<T>, attempts = 3, delayMs = 100): Promise<T> {
  let lastError: unknown;
  for (let attempt = 0; attempt < attempts; attempt++) {
    try {
      return await fn();
    } catch (err) {
      lastError = err;
      if (attempt < attempts - 1) {
        await new Promise((resolve) => setTimeout(resolve, delayMs * 2 ** attempt));
      }
    }
  }
  throw lastError;
}

/**
 * Vitest setup — polyfill Cloudflare-specific Web Crypto extensions.
 *
 * `crypto.subtle.timingSafeEqual` is a Cloudflare Workers extension that does
 * not exist in the standard @edge-runtime/vm environment used by Vitest.
 * We provide a functionally equivalent polyfill for testing purposes.
 */
if (typeof crypto !== "undefined" && crypto.subtle && !crypto.subtle.timingSafeEqual) {
  (crypto.subtle as unknown as Record<string, unknown>).timingSafeEqual = (
    a: ArrayBuffer,
    b: ArrayBuffer,
  ): boolean => {
    const aBytes = new Uint8Array(a);
    const bBytes = new Uint8Array(b);
    if (aBytes.byteLength !== bBytes.byteLength) return false;
    let diff = 0;
    for (let i = 0; i < aBytes.byteLength; i++) {
      diff |= aBytes[i] ^ bBytes[i];
    }
    return diff === 0;
  };
}

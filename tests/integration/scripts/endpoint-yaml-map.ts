/**
 * Maps the exact endpoint paths recorded by the integration test client to
 * their corresponding API doc YAML IDs in
 * apps/dashboard/config/api_docs/{api_id}.yml.
 *
 * When multiple paths map to the same api_id (e.g. both /encode and /decode
 * share "base64"), the injection script picks the entry with the most samples.
 */

export const endpointYamlMap: Record<string, string> = {
  // ── Validation ─────────────────────────────────────────────────────────────
  "/v1/validation/email": "email-validate",
  "/v1/text/normalize": "email-validate",
  "/v1/validation/phone": "phone-validation",
  "/v1/validation/profanity": "profanity",

  // ── Networking ─────────────────────────────────────────────────────────────
  "/v1/networking/disposable/check": "disposable-email",
  "/v1/networking/disposable/batch": "disposable-email",
  "/v1/networking/disposable/stats": "disposable-email",
  "/v1/networking/ip": "ip-info",
  "/v1/networking/ip/8.8.8.8": "ip-info",
  "/v1/networking/mx/gmail.com": "mx-lookup",

  // ── Text ───────────────────────────────────────────────────────────────────
  "/v1/text/words/random": "words",
  "/v1/text/lorem": "lorem-ipsum",
  "/v1/text/dictionary/eloquent": "dictionary",
  "/v1/text/thesaurus/happy": "thesaurus",
  "/v1/text/spellcheck": "spell-check",

  // ── Entertainment ──────────────────────────────────────────────────────────
  "/v1/entertainment/advice": "advice",
  "/v1/entertainment/quotes/random": "quotes",
  "/v1/entertainment/chuck-norris": "chuck-norris",
  "/v1/entertainment/jokes/dad": "dad-jokes",
  "/v1/entertainment/facts": "facts",
  "/v1/entertainment/trivia": "trivia",
  "/v1/entertainment/emoji/random": "emoji",
  "/v1/entertainment/sudoku": "sudoku",
  "/v1/entertainment/horoscope/aries": "horoscope",

  // ── Finance ────────────────────────────────────────────────────────────────
  "/v1/finance/mortgage": "mortgage",
  "/v1/finance/inflation": "inflation",
  "/v1/finance/exchange-rate": "exchange-rate",
  "/v1/finance/convert": "exchange-rate",
  "/v1/finance/crypto/BTC": "crypto",
  "/v1/finance/commodities/gold": "commodities",

  // ── Technology ─────────────────────────────────────────────────────────────
  "/v1/technology/password": "password-generator",
  "/v1/technology/useragent": "useragent",
  "/v1/technology/base64/encode": "base64",
  "/v1/technology/base64/decode": "base64",
  "/v1/technology/base": "number-base-conversion",
  "/v1/technology/color": "color-conversion",
  "/v1/technology/markdown": "data-format-conversion",
  "/v1/technology/random-user": "random-user",
  "/v1/technology/convert": "unit-conversion",
};

# Sentiment — Trade-off Analysis and Decision

This document justifies the selection of AFINN-165 as the sentiment lexicon for
the `sentiment` service (`apps/api/services/text/sentiment/service.go`),
replacing the current hardcoded 170-word lexicon, and explains why the
alternatives were discarded.

---

## Context

The sentiment service currently uses a hardcoded map literal of 170 words with
`float64` valence scores ranging from `−0.9` to `+0.9`. The algorithm — 3-token
negation lookahead and intensity multipliers — is correct and is not being
replaced. The problem is data coverage:

- Most real-world words are not in the 170-word set and return a score of zero.
- Everything trends toward neutral regardless of the actual sentiment of the
  text, making the output unreliable for any text beyond trivially simple
  inputs.
- The lexicon is hardcoded in `service.go`, coupling vocabulary updates to API
  deploys.

For context: AFINN-165 has 2,477 words and VADER has 4,000+. The current lexicon
covers approximately 7% of what AFINN would cover. The algorithm is
data-starved, not broken.

The decision was taken to replace only the lexicon — not the algorithm — and to
extract it into a dedicated `pkg/sentiment` package with its own test suite,
consistent with the bobadilla-tech OSS pattern.

---

## Alternatives Evaluated

### 1. Hardcoded 170-word lexicon (current, replaced)

A `map[string]float64` literal embedded directly in `service.go` with 170
manually curated entries.

**Rejected because:** vocabulary coverage is too narrow for commercial use. The
vast majority of real-world words score zero, making the output statistically
neutral regardless of actual sentiment. Updating the lexicon requires an API
redeploy. The lexicon has no independent test coverage.

### 2. VADER

A sentiment analysis tool designed specifically for informal text — social
media, reviews, and comments. It includes rules for capitalization, punctuation
intensity, and emojis in addition to a lexicon of 4,000+ words.

**Not selected as primary, but noted as upgrade path:** VADER is a Python
library. Its lexicon can be exported to JSON and loaded in Go, but requires a
manual conversion step. Its internal rules (capitalization handling, emoji
scoring) do not transfer — only the raw word scores do. For formally structured
text, AFINN provides equivalent coverage with no conversion overhead. VADER
remains the recommended upgrade path if the service needs to handle informal or
social-media text.

### 3. Oyemi

A deterministic semantic encoding system that assigns structured codes to words.
The last digit of each code encodes valence: `0` = neutral, `1` = positive, `2`
= negative. Coverage exceeds 145,000 words.

**Rejected because:** Oyemi returns a direction (positive / negative / neutral),
not a numeric magnitude. The algorithm requires `float64` scores to apply
intensity multipliers (`valence *= intensifierMult`) and accumulate signed sums
(`posSum`, `negSum`). Without numeric magnitude, both operations are undefined.
Additionally, Oyemi is a Python library with no Go binding, requiring an
external service call on every analysis request.

### 4. KeyNeg

A Python library specialised in extracting negative sentiment and keywords from
workforce-context text — employee feedback, surveys, and complaints. It detects
signals such as departure intent, escalation risk, and burnout, and returns
labelled categories rather than numeric scores.

**Rejected because:** KeyNeg is domain-specific (workforce) and does not provide
numeric valence scores compatible with the existing algorithm. It is not a
general-purpose lexicon replacement. It is a Python library with no Go binding.

### 5. spaCy

An industrial-strength NLP framework for Python with support for 75+ languages.
Its capabilities include named entity recognition, part-of-speech tagging,
dependency parsing, word vectors, and text classification pipelines.

**Rejected because:** spaCy is a framework for building NLP pipelines, not a
sentiment lexicon. It does not include native sentiment scoring. Integrating it
would require building a custom sentiment component on top of it, in Python,
served as a microservice. This is significant over-engineering for a problem
that is purely about vocabulary coverage.

### 6. Python microservice (Docker)

Packaging any Python sentiment library (VADER, spaCy, KeyNeg) into a small
FastAPI service, containerising it as a Docker image, and calling it from the Go
API via HTTP on every analysis request.

**Rejected for current scope:** introduces a network hop on every sentiment
call, adds a service dependency that must be operated and monitored, and creates
a failure mode where the sentiment endpoint returns 500 if the container is
unreachable. The quality gain over AFINN does not justify this operational cost
for a problem that is purely vocabulary coverage. This option remains valid if
transformer-based or ML-based sentiment becomes a future requirement.

---

## Decision: AFINN-165 via `//go:embed`

AFINN-165 is an open-source sentiment lexicon assigning integer scores from `−5`
to `+5` to 2,477 English words, covering a broad range of emotional vocabulary
including negations, intensifiers, profanity, and domain-specific terms.

It was selected because it occupies the optimal point on the
complexity-vs-quality curve for the platform's current needs:

| Criterion                        | 170-word lexicon | Oyemi   | VADER          | AFINN-165  | Microservice (Python) |
| -------------------------------- | ---------------- | ------- | -------------- | ---------- | --------------------- |
| Numeric scores (float64)         | ✅               | ❌      | ✅             | ✅         | ✅                    |
| Vocabulary coverage              | ❌ (~170 words)  | ✅ 145k | ✅ 4,000+      | ✅ 2,477   | ✅ varies             |
| Native Go — no external service  | ✅               | ❌      | ❌ (export)    | ✅         | ❌                    |
| No infrastructure overhead       | ✅               | ❌      | ❌             | ✅         | ❌ (Docker + network) |
| Independent lexicon update cycle | ❌               | —       | —              | ✅         | ✅                    |
| Algorithm compatibility          | ✅               | ❌      | ✅ (with work) | ✅         | ✅                    |
| License                          | —                | OSS     | MIT            | Apache 2.0 | varies                |

**Key reasons for selecting AFINN-165:**

- **Direct algorithm compatibility.** AFINN scores are integers in `[−5, +5]`.
  Each token's raw score is normalised individually — `(raw + 5) / 10` — before
  being accumulated into `posSum` or `negSum`. The final `Score` is the
  probability of the dominant class, derived from those sums. The negation
  lookahead, intensity multipliers, and signed accumulation logic inside
  `scoreTokens` are unchanged. The lexicon is a drop-in replacement.

- **14× vocabulary increase with zero new infrastructure.** The word map grows
  from 170 to 2,477 entries. No Docker container, no network call, no external
  process. The binary remains self-contained.

- **Lexicon decoupled from API deploys.** Moving the lexicon to
  `pkg/sentiment/lexicon.json` and embedding it via `//go:embed` gives the
  vocabulary its own release cycle. Updating AFINN entries, adding words, or
  swapping to VADER in the future means rebuilding only `pkg/sentiment` — not
  the full API.

- **Independent test surface.** Extracting the algorithm and lexicon into
  `pkg/sentiment` allows the negation lookahead, intensity multipliers, and
  score normalisation to be tested in isolation, without API routing,
  middleware, or serialisation noise.

- **License clarity.** AFINN-165 is released under the Apache 2.0 licence. It is
  compatible with commercial use and imposes no per-request or per-seat cost.

- **Consistent with the bobadilla-tech OSS pattern.** A standalone
  `pkg/sentiment` package with embedded data and a focused public API follows
  the same structure as other discrete packages in the repository.

---

## Score Normalisation

AFINN raw scores span `[−5, +5]`. The service response shape requires
`Score
float64` in `[0.0, 1.0]`. Normalisation is applied **per token**, before
accumulation, so that multi-word input never pushes the accumulated sum outside
a bounded range:

```text
Per token:
  normalised = (afinnRaw + 5) / 10        → maps [−5, +5] to [0.0, 1.0]
  posSum    += normalised  (if valence > 0)
  negSum    += normalised  (if valence < 0)

Output probabilities:
  denom  = posSum + negSum + 0.1          → neutralityWeight prevents any class reaching 1.0
  pos    = posSum / denom
  neg    = negSum / denom
  neu    = 1 − pos − neg

Score  = probability of the dominant class
Label  = "positive" | "neutral" | "negative"  (thresholds unchanged)
```

Per-token normalisation examples:

| AFINN raw | Per-token normalised | Accumulates into |
| --------- | -------------------- | ---------------- |
| +5        | 1.0                  | posSum           |
| +3        | 0.8                  | posSum           |
| 0         | 0.5                  | (skipped)        |
| −3        | 0.2                  | negSum           |
| −5        | 0.0                  | negSum           |

When no sentiment words are found, the service returns `Score = 1.0` and
`Sentiment = "neutral"` directly, without going through the probability
calculation. The response shape (`Score`, `Breakdown`, `Sentiment`) does not
change.

---

## Architecture

```text
Client
  ↓  HTTP  POST /v1/text/sentiment
Go API (requiems-api)
  ↓  calls
pkg/sentiment
  ├── Analyzer        (negation lookahead, intensity multipliers, score accumulation)
  ├── lexicon.json    (AFINN-165, embedded via //go:embed)
  └── Normalise()     (raw [−5,+5] → response [0.0,1.0])
  ↓
SentimentResult { Sentiment string, Score float64, Breakdown Breakdown }
  ↑
Go API  →  JSON response  →  Client
```

No external process, no network call, no environment variable required. The
`pkg/sentiment` package is self-contained: it embeds its own data and exposes a
single `Analyze(text string) Result` function.

---

## Package Structure

```text
pkg/
  sentiment/
    analyzer.go      ← tokenize(), scoreTokens(), Analyze() — ported verbatim
    lexicon.go       ← //go:embed lexicon.json + loadLexicon()
    normalise.go     ← raw score → [0.0, 1.0] + Label assignment
    analyzer_test.go ← unit tests independent of API context
    lexicon.json     ← AFINN-165 word list (2,477 entries)

apps/api/services/text/sentiment/
    service.go       ← thin wrapper: calls pkg/sentiment.Analyze(), returns HTTP response
```

The existing `service.go` is reduced to an HTTP handler. All logic moves to
`pkg/sentiment` where it can be tested and versioned independently.

---

## Known Trade-offs and Limitations

- **Recompilation required to update the lexicon.** `//go:embed` bakes
  `lexicon.json` into the binary at compile time. Adding or modifying words
  requires rebuilding `pkg/sentiment` and redeploying the API. This is a
  deliberate trade-off: it eliminates runtime file-path management and removes
  any possibility of the service starting with a missing or corrupt lexicon.
  Because the lexicon lives in its own package, the rebuild scope is narrow.

- **English-only.** AFINN-165 covers English vocabulary only. Extending the
  service to other languages would require a multilingual lexicon (e.g., a
  combined AFINN + NRC Emotion Lexicon dataset) or a different approach
  entirely.

- **No contextual understanding.** AFINN is a bag-of-words lexicon. It scores
  individual tokens without awareness of sentence structure beyond what the
  negation lookahead provides. Nuanced sentiment in complex sentences may be
  scored incorrectly.

- **Informal text coverage gaps.** AFINN does not include emojis, slang, or
  internet-specific vocabulary. Text from social media or chat interfaces will
  have higher zero-score rates than formal text, though significantly lower than
  with the 170-word lexicon.

---

## Future Upgrade Path

If sentiment quality becomes a business priority, the recommended next step is
replacing `lexicon.json` with **VADER's exported lexicon** (4,000+ words, better
informal text coverage). This is a file swap inside `pkg/sentiment` — the
algorithm, normalisation, and API surface are unchanged.

If contextual or transformer-based sentiment is required in the future, the
recommended path is a **Python FastAPI microservice** (e.g., using a fine-tuned
`distilbert-base-uncased-finetuned-sst-2-english` model), reachable via an
environment variable (e.g., `SENTIMENT_ML_URL`). The Go service layer would
remain unchanged; only the call target inside `service.go` would switch.

This migration path is clean because all sentiment logic is encapsulated inside
`pkg/sentiment`. The transport layer, validation, and documentation are
engine-agnostic.

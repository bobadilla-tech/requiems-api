# Spell Check Engine — Trade-off Analysis and Decision

This document justifies the selection of LanguageTool as the spell-check engine
for the `spell-check` API (REQAPI-350), replacing the previous `sajari/fuzzy`
implementation, and explains why the alternatives were discarded.

---

## Context

The spell-check API previously used `sajari/fuzzy`, a pure-Go fuzzy matching
library based on Levenshtein distance. While trivially easy to integrate (no
external dependencies, embedded in the Go binary), its correction quality was
inadequate for a commercial product:

- No contextual awareness: "tset" could be corrected to "test" or "set" with no
  way to choose the better option from surrounding words.
- English-only: no path to multilingual support without a complete rewrite.
- Single suggestion per word: the library returns one candidate, not a ranked
  list of alternatives.
- Misleading corrections on short or ambiguous inputs.

REQAPI-350 required adding a batch endpoint. Rather than scaling a low-quality
engine, the team took the opportunity to evaluate better alternatives and
migrate the underlying engine in the same PR.

---

## Alternatives Evaluated

### 1. sajari/fuzzy (current, replaced)

A pure-Go fuzzy matching library. Correction is based on Levenshtein distance
against a static word list embedded at compile time.

**Rejected because:** no contextual understanding, English-only, single
suggestion, poor correction quality on real-world inputs. Its only strength —
zero infrastructure overhead — does not justify the quality gap for a commercial
spelling API.

### 2. SymSpell

A correction algorithm optimised for speed using pre-generated delete
variations. Requires a word-frequency dictionary (a file mapping each word to
its corpus frequency) to rank candidates.

**Rejected because:** still no contextual understanding, so ambiguous typos are
resolved arbitrarily. Additionally, the most complete Go implementation carries
a restrictive license that is incompatible with commercial use.

### 3. Hunspell (via CGO)

The correction engine used by LibreOffice and Firefox. Mature, precise for
classical spellchecking, and available in many languages via `.dic`/`.aff`
dictionary pairs.

**Rejected because:** Go has no pure-Go Hunspell binding with suggestion
support. The only option was CGO (a bridge between Go and C), which
significantly complicates compilation, Docker images, and CI pipelines. The one
pure-Go wrapper found only exposed word validation (is this word correct?), not
suggestion generation, making it unusable for this API.

### 4. Transformer models — ByT5 / mT5 / T5

Neural sequence-to-sequence models trained on large multilingual corpora. They
understand full-sentence context, tolerate severe typos, and produce
high-quality corrections across many languages without dictionaries.

**Rejected (for now) because:** Go has no mature ML inference ecosystem. Serving
a transformer requires a separate Python microservice (FastAPI + PyTorch),
adding significant operational complexity. The smallest viable model
(ByT5-small, ~250 MB) takes 2–5 seconds per inference on CPU, which is
unacceptable for a synchronous API without dedicated GPU hardware. The quality
gain is real but the infrastructure cost is disproportionate for the current
stage of the platform.

This option remains the recommended upgrade path if correction quality becomes a
competitive differentiator in the future.

### 5. Hugging Face Inference API

Running transformer models via Hugging Face's hosted endpoints, removing the
need for local inference infrastructure.

**Rejected because:** the free tier imposes per-day request limits that are
incompatible with a multi-tenant commercial product. The paid tier introduces a
per-request cost and a hard dependency on a third-party service, both of which
are undesirable for a core text utility.

### 6. LLM APIs (OpenAI, Groq, etc.)

Using a large language model to perform spell correction via prompt.

**Rejected because:** the cost per token accumulates rapidly at volume, the
latency is higher than purpose-built correction tools, and the approach
introduces a proprietary external dependency for a task that does not require
general language understanding. Architecturally, LLMs are significantly
over-engineered for spellchecking.

---

## Decision: LanguageTool

LanguageTool is an open-source grammar, style, and spell-checking engine that
combines rule-based NLP, linguistic heuristics, and statistical language models.
It exposes an HTTP API and runs as a standalone Java service.

It was selected because it occupies the optimal point on the
complexity-vs-quality curve for the platform's current needs:

| Criterion                       | sajari/fuzzy | SymSpell | Hunspell | LanguageTool   | Transformers      |
| ------------------------------- | ------------ | -------- | -------- | -------------- | ----------------- |
| Contextual understanding        | ❌           | ❌       | ❌       | Partial ✅     | ✅                |
| No external dictionary required | ✅           | ❌       | ❌       | ✅             | ✅                |
| No infrastructure overhead      | ✅           | ✅       | ❌ (CGO) | ❌ (Docker)    | ❌ (Python + GPU) |
| Multilingual support            | ❌           | ❌       | ✅       | ✅ (30+ langs) | ✅                |
| Multiple suggestions per word   | ❌           | ✅       | ✅       | ✅             | ✅                |
| Compatible with existing infra  | ✅           | ✅       | ❌       | ✅             | ❌                |
| Correction quality              | Poor         | Fair     | Good     | Good           | Excellent         |

**Key reasons for selecting LanguageTool:**

- **Contextual correction without ML infrastructure.** LanguageTool uses n-gram
  language models to consider surrounding words when ranking candidates. This
  eliminates the most common class of absurd corrections produced by pure
  edit-distance approaches.

- **No licensing cost for self-hosted deployment.** The Community edition is
  released under the LGPL licence. Running it on Bobadilla's infrastructure
  incurs zero per-request or per-seat cost regardless of how many API clients
  use the spell-check endpoint.

- **Client transparency.** LanguageTool runs entirely on Bobadilla's servers.
  Clients call `/v1/text/spellcheck` and receive results — they have no
  dependency on, or knowledge of, LanguageTool. Licensing and infrastructure
  concerns are fully encapsulated on the server side.

- **Incremental infrastructure.** Adding a Docker container is a smaller step
  than introducing a Python ML microservice. The team already runs Docker in all
  environments and the existing CI/CD pipeline requires no changes.

- **Multilingual path.** LanguageTool supports 30+ languages out of the box via
  the same API. Extending the spell-check service to additional languages is a
  configuration change, not a rewrite.

- **Multiple suggestions.** LanguageTool returns a ranked list of replacement
  candidates. The implementation surfaces the top three, giving frontend clients
  enough options to present a correction picker to end users.

---

## Architecture

```
Client
  ↓  HTTP  POST /v1/text/spellcheck[/batch]
Go API (requiems-api)
  ↓  HTTP  POST /v2/check  (form-encoded)
LanguageTool Server  (Java, Docker container)
  ↓
Correction response (JSON)
  ↑
Go API  →  structured Result / batch Results  →  Client
```

The LanguageTool base URL is injected via the `LANGUAGETOOL_URL` environment
variable (default: `http://localhost:8010`). This allows each environment to
point to its own instance without code changes.

---

## Infrastructure Requirements

### Development (local)

```bash
docker run -d --name languagetool -p 8010:8010 erikvl87/languagetool
```

Requirements: ~2 GB RAM, 1 vCPU. Any developer machine running Docker can run
this without additional setup.

### Production

A single LanguageTool container added to the existing Docker Compose or
Kubernetes deployment. Recommended sizing for moderate load:

| Resource | Minimum      | Recommended  |
| -------- | ------------ | ------------ |
| RAM      | 2 GB         | 4 GB         |
| vCPU     | 1            | 2            |
| GPU      | Not required | Not required |

No external network access is required — LanguageTool runs fully air-gapped on
the server's internal network.

---

## Known Trade-offs and Limitations

- **Requires a running Java process.** Unlike `sajari/fuzzy`, which was embedded
  in the Go binary, LanguageTool is a separate service. If it is unreachable,
  the spell-check endpoint returns 500. Startup and health-check dependencies
  must be configured in production.

- **Correction quality on very short or isolated tokens.** LanguageTool performs
  best on complete sentences. Single-word or very short inputs (e.g., a single
  misspelled word with no surrounding context) produce lower-quality suggestions
  than full sentences, similar to the weakness of edit-distance approaches.

- **Not fully neural.** LanguageTool's language models are statistical n-grams,
  not transformer-based. It does not "reason" about text the way a model like
  ByT5 would. Contextually ambiguous corrections may still be wrong.

- **Java runtime overhead.** Cold-start time for the LanguageTool container is
  ~10 seconds. In production this is absorbed at deployment time, but local
  development environments must ensure the container is running before starting
  the API.

---

## Future Upgrade Path

If correction quality becomes a business priority, the recommended next step is
a **ByT5-small** (or similar encoder-decoder transformer) served as a Python
FastAPI microservice, reachable via a second environment variable (e.g.,
`SPELLCHECK_ML_URL`). The Go service layer would remain unchanged; only the HTTP
call target would switch.

This migration path is clean because the engine is fully encapsulated inside
`services/text/spellcheck/service.go`. The transport layer, batch logic,
validation, and documentation are engine-agnostic.

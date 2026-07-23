# Thesaurus — Trade-off Analysis and Decision

This document justifies the selection of **Open English WordNet (OEWN)** for
synonyms combined with a **curated antonym override layer** as the data source
for the `thesaurus` service
(`apps/api/services/text/thesaurus/service.go`), replacing the current
hardcoded 100-word lexicon, and explains why the alternatives were discarded.

---

## Context

The thesaurus service currently uses a hardcoded map literal of ~100 words in
`data.go`, with a `Synonyms []string` / `Antonyms []string` structure per word.
The problem is data coverage:

- Any word outside the ~100-word set returns an empty response.
- Adding a word requires editing `data.go` and redeploying the API.
- There is no independent test coverage for the lexicon itself.

The API response shape (`Word`, `Synonyms []string`, `Antonyms []string`) is a
hard backward-compatibility constraint and does not change under any option
considered below.

The decision was taken to source **synonyms** from a large, well-licensed
external dataset, while continuing to source **antonyms** from a small curated
override layer — because, as detailed below, no evaluated open dataset
provides antonym coverage cheap enough to justify replacing the curated layer.
Grammatical tagging (POS tagging) and multilingual support are out of scope.

---

## Alternatives Evaluated

### 1. Hardcoded 100-word lexicon (current, replaced)

A `map[string]struct{Synonyms []string; Antonyms []string}` literal embedded
directly in `data.go`.

**Rejected because:** coverage is far too narrow for production use — the vast
majority of real-world words return an empty result. Updating the lexicon
requires an API redeploy. No independent test coverage.

### 2. Moby Thesaurus

A public-domain dataset from Project Gutenberg with ~30,000 root words and
2.5M synonym entries, distributed as plain text requiring parsing to JSON.

**Not selected as primary, but viable as a fallback:** Moby has strong
synonym volume and is straightforward to parse into a flat
`word → []string` map. However:

- It does not include antonyms at all — the curated override layer would
  still be required regardless of which synonym source is chosen.
- It is a 1990s dataset with no active maintenance; entries may be arcane,
  noisy, or inconsistent by modern usage.
- Its license provenance is public domain, but documented less rigorously
  than WordNet's — a real diligence gap for anything shipped as attributed
  OSS.

Moby remains a reasonable fallback if OEWN parsing effort turns out to be
prohibitive, but it does not outperform OEWN on any axis except parsing
simplicity.

### 3. Open English WordNet (OEWN) — selected for synonyms

A public dataset maintained by the Global WordNet Association, distributed in
several formats:

- **GWN-LMF (XML)** — the official, complete format. Models
  `LexicalEntry → Sense → Synset`, including both `SynsetRelation`
  (synset-to-synset, e.g. hypernym/hyponym) and `SenseRelation`
  (sense-to-sense, e.g. antonym).
- **WNDB** — the classic Princeton flat-file format, compatible with legacy
  tooling (JWI, NLTK), but with weaker native Go tooling support.
- **Flattened JSON derivatives** (e.g. community datasets published on
  Hugging Face) — easiest to consume, but see the antonym finding below for
  why this format is insufficient on its own.

**Selected because:**

- Synonym extraction is a straightforward inverted index: for a given word,
  find every synset where it appears as a `member`, and collect the other
  members of those synsets as synonyms.
- License and provenance are clearly documented (Global WordNet Association),
  unlike Moby.
- Actively maintained, reducing the arcane/noisy-entry risk inherent to a
  1990s dataset.

**Rejected as the antonym source** — see the dedicated finding below.

### 4. ConceptNet

A broader semantic network (not English-only, not lexical-only) that includes
explicit `Antonym` and `SimilarTo` relations, under a CC/Apache license.

**Rejected because:** significantly heavier than needed for a simple
synonym/antonym lookup. Its multilingual scope is explicitly out of this
service's scope. It would be reasonable to revisit if the service's ambitions
expand beyond lexical synonym/antonym lookup (e.g. broader semantic
relatedness).

### 5. Roget's Thesaurus (Project Gutenberg)

Same era and public-domain provenance as Moby, but organized by conceptual
category rather than word → word.

**Rejected because:** requires more transformation work to reach the target
`Entry{Synonyms, Antonyms}` shape than either Moby or OEWN, with no coverage
or licensing advantage over OEWN to justify the extra effort.

### 6. Database-backed storage (instead of embedding data in the binary)

Storing the parsed dataset in a relational or document database instead of
embedding it via `//go:embed`, with the service querying it at request time.

**Rejected for current scope:** the dataset is read-only and changes
infrequently (occasional OEWN version bumps, occasional antonym corrections).
A database converts a nanosecond in-memory map lookup into a network
round-trip, and introduces an operational dependency (provisioning, backups,
monitoring) disproportionate to the problem. It would become justified if the
service later needs richer queries the embedded map cannot support — for
example, full-text/fuzzy search for "did you mean…" suggestions, where a
database with trigram or full-text indexing (Postgres `pg_trgm`,
Elasticsearch/Meilisearch) earns its complexity on its own merits, not merely
as storage for the thesaurus data.

---

## Key Finding: WordNet Antonyms Are Not a Drop-in Replacement

Antonymy in WordNet is a **lexical relation** (specific word ↔ specific word),
modeled almost exclusively as `SenseRelation` in the official LMF schema — not
as `SynsetRelation`. The format documentation is explicit that antonymy cannot
be correctly modeled at the synset level without violating the Wordnet-LMF
DTD, because a synset's members do not all share the same antonym (e.g. "good"
the adjective is the antonym of "bad" specifically, not of every member of its
synset).

Practical consequence: **flattened, synset-level JSON exports of OEWN (e.g.
community datasets that expose only `SynsetRelation`) do not carry usable
antonym data**, regardless of how many rows are inspected — the granularity
needed for antonyms was discarded during that flattening, not merely sparse.
Extracting antonyms correctly requires parsing the official LMF XML (or WNDB)
at the `Sense` level — a materially more complex three-level parse
(`entry → sense → relation`) than the flat `word → []string` shape used for
synonyms.

Given that antonym relations are also inherently far less numerous than
synonym/hypernym relations in WordNet, the added parsing complexity is not
justified by the yield. **The existing curated antonym layer is retained.**

---

## Decision Summary

| Criterion                         | 100-word lexicon | Moby       | OEWN (synonyms) | ConceptNet | Database-backed |
| ---------------------------------- | ----------------- | ---------- | ---------------- | ---------- | ---------------- |
| Vocabulary coverage                | ❌ (~100 words)   | ✅ ~30k    | ✅ ~155k         | ✅ broad   | — (storage only) |
| Antonym coverage (native)          | ✅ (curated)      | ❌ none    | ⚠️ costly to extract | ✅ | — (storage only) |
| Parsing complexity                 | —                 | ✅ low     | ⚠️ moderate (synsets) | ❌ high | — |
| License clarity                    | —                 | ⚠️ weak    | ✅ documented    | ✅ documented | — |
| Native Go, no external service     | ✅                | ✅         | ✅               | ✅ | ❌ (DB dependency) |
| No infrastructure overhead         | ✅                | ✅         | ✅               | ✅ | ❌ (DB ops) |
| Independent update cycle           | ❌                | ✅         | ✅               | ✅ | ✅ |
| In scope for this service          | —                 | fallback   | ✅ primary       | ❌ out of scope | ❌ rejected |

**Final decision:**

- **Synonyms:** Open English WordNet (OEWN), parsed from the official GWN-LMF
  XML into a flat `word → []string` inverted index built from synset
  membership.
- **Antonyms:** the existing ~100 curated entries, retained as an override
  layer applied at runtime over the OEWN-derived synonyms — unchanged in
  spirit from the original Moby-based proposal.
- **Storage:** embedded in the binary via `//go:embed`, **gzip-compressed**,
  not a database (see storage rationale below).

---

## Storage: `//go:embed` with gzip compression

The dataset is read-only, changes infrequently, and sits in the hot path of
every `Lookup()` call — a profile that favors in-memory embedding over a
database. To address the binary-size concern raised against embedding a
dataset of this scale, the JSON is gzip-compressed at build time and
decompressed once at process startup:

```text
Build time (once):
  synonyms.json  →  gzip  →  synonyms.json.gz   (embedded via //go:embed)

Process startup (once):
  synonyms.json.gz  →  gzip.NewReader  →  json.Unmarshal  →  in-memory map

Request time (every Lookup call):
  in-memory map access — no decompression, no I/O, no network
```

**What gzip compression improves:**

- Compiled binary size — JSON's repeated structural tokens (keys, brackets,
  quotes) compress well, typically 70–85% smaller.
- Container image size and registry pull time.
- CI/CD artifact transfer time.

**What gzip compression does not improve:**

- Runtime memory usage — once decompressed at startup, the in-memory map
  occupies the same space it would without compression. If per-replica memory
  becomes a concern, the mitigation is a more compact binary serialization
  format (e.g. FlatBuffers/msgpack) or an out-of-process store, not gzip.
- `Lookup()` latency — decompression happens once at startup, not per
  request.
- Startup time is affected marginally (a few milliseconds to decompress),
  which is acceptable for this service's deployment profile.

---

## Architecture

```text
Client
  ↓  HTTP  GET /v1/text/thesaurus?word=...
Go API (requiems-api)
  ↓  calls
pkg/thesaurus
  ├── Lookup(word string) (Entry, bool)
  ├── synonyms.json.gz    (OEWN-derived, embedded via //go:embed)
  ├── antonyms.json       (curated override layer, embedded via //go:embed)
  └── loadData()           (gzip decompress + JSON unmarshal, once at startup)
  ↓
Entry { Synonyms []string, Antonyms []string }
  ↑
Go API  →  JSON response  →  Client
```

No external process, no network call, no database connection. The
`pkg/thesaurus` package is self-contained: it embeds its own data and exposes
a single `Lookup(word string) (Entry, bool)` function, combining OEWN
synonyms with the curated antonym override at runtime.

---

## Package Structure

```text
pkg/
  thesaurus/
    lookup.go          ← Lookup(), applies antonym override over synonym index
    data.go             ← //go:embed synonyms.json.gz, antonyms.json + loadData()
    normalise.go        ← basic word normalisation (lowercase, trim)
    lookup_test.go      ← known word, unknown word, override application
    synonyms.json.gz    ← OEWN-derived synonym index (gzip-compressed)
    antonyms.json        ← curated antonym overrides (~100 entries)

apps/api/services/text/thesaurus/
    service.go          ← thin wrapper: calls pkg/thesaurus.Lookup(), returns HTTP response
```

The existing `service.go` is reduced to an HTTP handler; `data.go` is removed
from `apps/api` entirely. All lexicon logic and data move to `pkg/thesaurus`,
where they can be tested and versioned independently.

---

## Known Trade-offs and Limitations

- **Recompilation required to update the dataset.** `//go:embed` bakes both
  `synonyms.json.gz` and `antonyms.json` into the binary at compile time.
  Updating to a newer OEWN release or correcting a curated antonym requires
  rebuilding `pkg/thesaurus` and redeploying the API. This is a deliberate
  trade-off, consistent with the equivalent decision in `pkg/sentiment`: it
  removes any possibility of the service starting with a missing or corrupt
  dataset, and the rebuild scope is narrow because the data lives in its own
  package.

- **Antonym coverage remains capped at the curated set (~100 words).**
  Extending antonym coverage beyond the current curated layer requires manual
  curation effort or a future investment in parsing OEWN's `SenseRelation`
  data at the XML level — not a quick dataset swap.

- **English-only.** OEWN and the curated antonym layer cover English
  vocabulary only. Multilingual support would require a different dataset
  entirely and is out of scope for this service.

- **No POS-aware disambiguation.** `Lookup()` returns synonyms aggregated
  across all synsets a word belongs to, regardless of part of speech. A word
  with unrelated meanings across POS categories (e.g. a noun/verb pair) will
  return synonyms from both senses undifferentiated. POS tagging is
  explicitly out of scope.

- **1990s-dataset risk avoided, but not eliminated.** OEWN is actively
  maintained, unlike Moby, but as with any large lexical resource it may still
  contain archaic or low-frequency entries; no filtering or result-quality
  thresholding is applied in this scope.

---

## Future Upgrade Path

If antonym coverage becomes a priority, the recommended next step is parsing
OEWN's official **GWN-LMF XML** at the `Sense`/`SenseRelation` level to extract
native antonym pairs, replacing (or supplementing) the curated override layer.
This is isolated to `pkg/thesaurus/data.go` and the build-time data
generation step — the `Lookup()` signature and API response shape are
unchanged.

If fuzzy matching or "did you mean…" suggestions become a requirement, the
recommended path is introducing a database with trigram/full-text indexing
(Postgres `pg_trgm`, or a dedicated search engine) specifically for that
query pattern — not as a replacement for the embedded lookup path, which
should remain in place for exact-match lookups given its latency advantage.

This migration path is clean because all thesaurus logic is encapsulated
inside `pkg/thesaurus`. The transport layer, validation, and documentation
remain engine-agnostic.

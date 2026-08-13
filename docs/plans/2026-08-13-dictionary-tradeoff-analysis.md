# Dictionary — Trade-off Analysis and Decision

This document justifies the selection of **Wiktionary (via wiktextract raw
data, distributed by kaikki.org)** as the broad-coverage source for
`etymology`, `phonetic` (UK/US/Other), `partOfSpeech`, `definition`, and
`example` data in the `go-dictionary` package, and explains why the
alternatives — including Open English WordNet (OEWN) — were discarded for
this specific package.

The pre-existing hand-curated dataset of ~30 words is retained alongside
this decision, not replaced by it: it remains embedded separately (as
`dataset/curated.json`) and exposed through its own getter, `GetCurated`,
independent of and unmerged with the Wiktionary-derived dataset exposed
through `Get`. OEWN remains in the codebase, but its scope is limited
entirely to `pkg/thesaurus` (synonyms), which is unaffected by this document
and covered by its own companion trade-off analysis.

---

## Context

The `words` service needs to look up English words and return phonetics,
part of speech, definitions, and examples. A small hand-curated map of ~30
words (`dataset/curated.json`, exposed via `GetCurated`) already covers a
handful of common words with high-confidence, hand-authored content, but its
coverage is far too narrow for general use:

- Any word outside that ~30-word set returns nothing from the curated path.
- Growing the curated set requires hand-authoring new entries and
  redeploying.
- It has no independent, generalizable pronunciation or definition source
  behind it — it's a fixed, manually maintained list.

The package needs a second, much broader source to sit alongside the
curated set for the majority of real-world lookups. This document evaluates
candidates for that broader source.

The resulting package exposes two independent getters with two different
shapes, and does not merge them:

- `GetCurated(word) (CuratedEntry, bool)` →
  `{ Phonetic, Definitions[]{PartOfSpeech, Definition, Example}, Synonyms[] }`
  for the ~30 curated words.
- `Get(word) (Entry, bool)` →
  `{ Word, Variants[]{ Etymology, PhoneticUK, PhoneticUS, PhoneticOther,
  Definitions[]{PartOfSpeech, Definition, Examples[]}, SenseCount } }` for
  the Wiktionary-derived broad-coverage set.

Deciding source precedence for a given word, and reconciling the two shapes
into a single API response, is left to the calling service
(`apps/api/services/text/words/service.go`) — this package's own
responsibility is scoped to data access only, with no merge or selection
logic (see Architecture, below).

The `synonyms` field is out of scope for this document: it is served by
`pkg/thesaurus`, backed by OEWN, and is not re-sourced here.

---

## Alternatives Evaluated

### 1. Hand-curated word list alone

A small, hand-authored `map[string]entry` (shipped as `dataset/curated.json`,
exposed via `GetCurated`).

**Not sufficient as the sole source** — coverage is far too narrow for
production use on its own, and growing it requires hand-authoring and a
redeploy for every new word. It is kept as a fast, high-confidence path for
the words it does cover, used at the calling service's discretion alongside
a broader source.

### 2. Free Dictionary API (dictionaryapi.dev)

A free, keyless JSON API returning phonetics, definitions, examples, and
synonyms per word.

**Rejected because:**

- It is a live API, not a downloadable dataset. This does not fit the
  `//go:embed` pattern already established for `pkg/thesaurus` and
  `pkg/sentiment` — embedding requires a static artifact generated at build
  time, not a runtime HTTP dependency.
- No formally published rate limit or documented license/terms for the
  underlying data; provenance of individual entries is not rigorously
  documented, which is a real diligence gap for anything shipped as an
  attributed, embedded dataset.
- Quality is inconsistent across words: some entries are rich, others sparse
  or missing fields entirely, with no way to bulk-audit coverage ahead of a
  build.

### 3. Merriam-Webster Dictionary API

A commercial dictionary API with authoritative definitions, etymologies,
audio pronunciations, and specialized (medical, Spanish, ESL) vocabularies.

**Rejected because:** the free tier is licensed for non-commercial use only;
commercial use — which this service is — requires a paid license agreement.
Adopting it would introduce an ongoing per-word or per-call cost and a vendor
dependency the embedded-dataset approach is specifically designed to avoid.

### 4. Open English WordNet (OEWN) — rejected for this package

Already embedded and integrated for `pkg/thesaurus` via the `oewn` `Provider`
that parses the official GWN-LMF XML.

**Rejected for the broad-coverage path because, once phonetics and etymology
were scoped in as requirements, consolidating on OEWN for
definition/example/partOfSpeech while sourcing phonetics from elsewhere
would have meant maintaining two XML/JSON parsers, two source-update cycles,
and two stylistically distinct sets of prose for the same fields** — a cost
a single-pipeline design avoids entirely. Specific findings that informed
this, verified empirically against the real 2024-edition GWA XML
(`english-wordnet-2024.xml.gz`, 120,630 synsets, 161,705 `LexicalEntry`
elements):

- `Definition`: **100% coverage** — every synset carries one, no exceptions.
- `Example`: **27.9% coverage** at the synset level (33,604 of 120,630
  synsets carry at least one usage example) — meaning roughly 3 in 4 OEWN
  synsets would ship with an empty `example` field.
- `Pronunciation`: **27.6% coverage** (44,669 of 161,705 `LexicalEntry`
  elements), a single unlabeled transcription with no US/UK distinction —
  insufficient on its own to satisfy the UK/US phonetics requirement, which
  is precisely why an external phonetics source was needed in the first
  place.
- A direct comparison of a curated entry against its real OEWN synset found
  the curated definition, example, and 3 of 5 synonyms did not match OEWN's
  actual content at all — confirming the existing curated dataset was
  independently authored, not OEWN-derived, and that grafting fresh OEWN
  extractions onto it would produce visibly inconsistent entries.

Given `example` coverage under 28% and phonetics under 28% with no
accent/dialect data, OEWN alone cannot satisfy the broad-coverage path's
requirements. OEWN is not discarded from the codebase; it remains the
correct, already-validated source for `pkg/thesaurus`, which has no
phonetics requirement and where its ~28% pronunciation gap is irrelevant.

### 5. Wiktionary (raw wiktextract data via kaikki.org) — selected, broad-coverage source

Structured JSON Lines (JSONL) data extracted from the English Wiktionary
edition using the `wiktextract` tool, distributed at
`kaikki.org/dictionary/rawdata.html`.

**Selected because it is the only evaluated source that plausibly covers
etymology, `phonetic` (UK/US), `partOfSpeech`, `definition`, and `example`
for English words from a single pipeline:**

- **Format**: raw data is JSONL — one JSON object per line, per word +
  language + part-of-speech entry, filtered to `lang_code == "en"` by the
  generator, since the raw feed is not pre-split by language.
- **Phonetics field**: pronunciation data lives under a `sounds` key — a
  *list* of dictionaries, each of which may include `ipa`, `enpr`,
  `audio-ipa`, `tags` (region/dialect labels), and `text`. Because it is a
  list, a word commonly carries multiple pronunciations, each tagged by
  dialect — this is what makes UK/US/Other separation possible at all,
  unlike OEWN's single unlabeled transcription.
- **Part of speech, definitions, and etymology**: each top-level word entry
  carries its own `pos` field, a `senses` list (each sense carrying its own
  gloss), and etymology text — structurally analogous to what would be
  extracted from OEWN's `Sense`/`Synset` tree, but from one already-parsed
  JSON object.
- **License**: made available under the same licenses as Wiktionary itself —
  CC-BY-SA and GFDL. This is a cleaner chain than the postprocessed dataset,
  whose own page discloses that additional data is merged in from other
  sources, introducing license-provenance uncertainty the raw feed does not
  carry.
- **Operational status**: the postprocessed English-only dataset's download
  page is deprecated and being actively removed per an open, tracked project
  issue. The raw feed carries no such deprecation notice and is documented
  as updated regularly (at least weekly), which additionally supports
  pinning a specific dump date for reproducible builds.

**Licensing decision:** the ShareAlike obligation of CC-BY-SA was evaluated
and the risk was accepted on the understanding that the resulting dataset is
served only through the package's own API and never distributed as a
standalone downloadable artifact, dataset export, or open-source repository
independent of attribution. This is recorded in the package's `Credits`
section, citing CC-BY-SA, GFDL, and the wiktextract tool (MIT-licensed,
separate from the license of the data it extracts).

**Key implementation decisions:**

- **Dialect is exposed, not selected.** Rather than collapsing a word's
  pronunciations down to a single `phonetic` value, the design keeps three
  separate fields — `PhoneticUK`, `PhoneticUS`, `PhoneticOther` — and makes
  no selection between them; the calling service decides what to surface.
  An inventory pass over `sounds[].tags` across the wordlist-filtered corpus
  found that the bare `US`/`UK` tags alone undercounted real coverage: the
  tags `General-American` and `Received-Pronunciation`/`British` carry
  substantial additional volume and are recognized alongside `US`/`UK`.
  Everything else (Australian, Canadian, Scottish, register qualifiers like
  "dialectal" or "archaic") collapses into `PhoneticOther` as a single
  value — a multi-dialect map beyond UK/US/Other was considered and
  explicitly rejected as out of scope.
- **Senses are exposed as a full list, not selected down to one.** The
  design returns the full `Definitions []Definition` list per `Variant`
  (each with up to 5 deduplicated `Examples`), plus a `SenseCount` field as
  a signal for callers that want to choose one. Definitions are
  deduplicated by exact gloss text (some entries repeat identical glosses
  once per citation), and `"example"`-type sentences are preferred over
  `"quotation"`-type ones when both exist for the same sense.
- **Entries are grouped by word and etymology, not by word and part of
  speech.** A single etymology commonly spans multiple parts of speech
  (flattened into one `Variant`), while a genuinely different etymology for
  the same spelling (e.g. "name" the identifier vs. "name" the Caribbean
  yam) produces a separate `Variant`, never merged. Etymology prose is
  cleaned of two kinds of machine-generated noise ("Etymology tree" and
  "Cognates" blocks) as a best-effort heuristic, not a guaranteed parse of
  every format.
- **Sensitive senses are filtered at generation time.** Senses tagged
  `vulgar`, `derogatory`, `offensive`, or `slur` are excluded from the
  generated dataset, chosen from real tag-frequency data as the tags
  Wiktionary itself uses to flag content needing special handling —
  deliberately not `slang` or `euphemistic`, which are common and mostly
  not sensitive. This filter only catches what the source tagged
  explicitly.
- **Field-coverage completeness was empirically verified, in two passes**:
  a global pass over the full raw dump, and a second pass restricted to a
  common-English wordlist, to test whether the global numbers were being
  dragged down by long-tail vocabulary this service is unlikely to be
  queried for.

  Global pass (10,736,399 total lines read, 1,481,704 English entries):

  | Field | Coverage |
  | --- | --- |
  | `definition` (`senses[].glosses`) | 99.9% (1,480,504 / 1,481,704) |
  | `example` | 19.4% (287,234 / 1,481,704) |
  | any `ipa` in `sounds[]` | 9.3% (138,189 / 1,481,704) |
  | `ipa` tagged `US` | 4.0% (58,555 / 1,481,704) |
  | `ipa` tagged `UK` | 4.0% (59,484 / 1,481,704) |
  | entries with multiple senses | 9.0% (133,772 / 1,481,704) |

  Restricted to a 50,000-word frequency list (`hermitdave/FrequencyWords`,
  `en_50k.txt`), combined coverage across the fields measured rose to
  **60.6%** — well above OEWN's equivalents (27.6% phonetics, 27.9%
  example). The low global numbers are a long-tail artifact: Wiktionary
  documents roughly 1.48M English entries versus OEWN's 157,513 lemmas, and
  that ~9x larger vocabulary is overwhelmingly rare, technical, or archaic
  terms that Wiktionary's editors have defined but rarely transcribed
  phonetically or illustrated with a usage example. The service being
  built here will only ever be queried for common vocabulary, so the global
  figure measures the wrong population.

  One caveat surfaced by the field-level breakdown: `US`-tagged coverage
  plus `UK`-tagged coverage does not sum to the "any `ipa`" figure — there
  is overlap (words with both tags) and a meaningful share of entries carry
  a populated `ipa` with **neither** tag. This is exactly why the design
  adds `PhoneticOther` as a third field rather than treating dialect as a
  binary.

### 6. CMU Pronouncing Dictionary

A permissively licensed pronunciation dictionary for American English, using
ARPAbet transcription.

**Discarded** given Wiktionary covers etymology, `phonetic`, `partOfSpeech`,
`definition`, and `example` from a single source. Worth recording why it was
never a serious contender for the full scope: it carries no semantic data
(no definitions, no examples), only American English (no UK variant, which
the curated dataset's style depends on), and ARPAbet transcription would
have required a conversion layer Wiktionary's native IPA output does not.

### 7. Database-backed storage (instead of embedding data in the binary)

Storing the parsed dataset in a relational or document database instead of
embedding it via `//go:embed`, with the service querying it at request time.

**Rejected**, for the same reasons already established for `pkg/thesaurus`:
the dataset is read-only and changes infrequently, sits in the hot path of
every lookup, and a database converts a nanosecond in-memory map lookup into
a network round-trip while adding an operational dependency disproportionate
to the problem.

---

## Key Finding: Global Coverage Understated the Real Picture — Verification Against a Common-Vocabulary Filter Was Required

**First pass (global, all 1,481,704 English entries): the numbers looked
worse than OEWN, not better.** `ipa` coverage was 9.3%, roughly a third of
OEWN's 27.6%, and `example` coverage was 19.4%, below OEWN's 27.9% — the
opposite of what would justify moving off OEWN. Taken at face value, this
would have called the whole decision into question.

**Second pass (restricted to a 50,000-word common-English list): the
numbers inverted.** Combined field coverage rose to 60.6% — well above
OEWN's equivalents. The explanation is straightforward: Wiktionary's ~9x
larger vocabulary is overwhelmingly long-tail, and the service being built
here will only ever be queried for common vocabulary, so the global figure
was measuring the wrong population entirely.

**Practical consequence for the generator:** this is not just a
documentation footnote, it is an architectural input. Filtering the source
data to common vocabulary before generating `dictionary.json.gz` — rather
than embedding all 1.48M entries — improves both data quality and binary
size. See Future Upgrade Path for how this affects dataset generation scope.

---

## Decision Summary

| Criterion                              | Curated word list | Free Dictionary API | Merriam-Webster | OEWN | Wiktionary (raw, selected) | CMU | Database-backed |
| --------------------------------------- | ---------------- | -------------------- | ---------------- | ---- | --------------------------- | --- | ---------------- |
| Vocabulary coverage                     | ❌ (~30 words)   | ✅ broad             | ✅ broad         | ✅ 157,513 lemmas (verified) | ✅ broad (millions of entries) | ⚠️ US-only | — (storage only) |
| Fits `//go:embed` (static artifact)     | ✅               | ❌ live API          | ❌ live API/paid | ✅   | ✅ (dump-based)              | ✅  | ❌ (DB dependency) |
| License clarity                         | —                 | ⚠️ undocumented      | ⚠️ paid for prod | ✅ documented | ✅ documented (CC-BY-SA + GFDL), no "merged from other sources" caveat | ✅ documented, permissive | — |
| Role in the shipped package              | ✅ kept, unmerged, via `GetCurated` | rejected | rejected | scoped to `pkg/thesaurus` only | ✅ selected, via `Get` | rejected | rejected |
| Covers etymology + all phonetic/definition/example fields from one source | —            | ✅ (but live API)    | ✅ (but paid)    | ❌ (no phonetic UK/US, no etymology) | ✅ selected | ❌ (no definitions, no etymology) | — |
| `phonetic` UK/US/Other coverage          | —                 | ⚠️ inconsistent      | ✅               | ❌ single unlabeled form, 27.6% coverage | ✅ 60.6% combined on 50k common-vocab list (verified; 9.3% unrestricted) | ❌ US only | — |
| `definition` coverage                    | —                 | ⚠️ inconsistent      | ✅               | ✅ 100% (verified) | ✅ 99.9% (verified) | ❌ none | — |
| `example` coverage                       | —                 | ⚠️ inconsistent      | ✅               | ⚠️ 27.9% (verified) | ✅ well above OEWN on common vocab (verified; 19.4% unrestricted) | ❌ none | — |
| Reuses existing pipeline/`Provider`      | —                 | ❌                    | ❌               | ✅ (but insufficient alone) | ❌ new provider needed | ❌ | — |
| No infrastructure overhead               | ✅               | ✅                    | ✅               | ✅   | ✅                            | ✅  | ❌ (DB ops)       |
| Independent update cycle                 | ❌               | ✅                    | ✅               | ✅   | ✅ (weekly dumps, documented) | ✅  | ✅                |
| In scope for this package                | ✅ retained as a second source | rejected | rejected | ❌ out of scope, remains sole source for `pkg/thesaurus` | ✅ selected as the broad-coverage source | rejected | rejected |

**Final decision:**

- **Broad-coverage path (`etymology`, `phonetic` UK/US/Other,
  `partOfSpeech`, `definition`, `example`):** Wiktionary, via raw
  `wiktextract` JSONL data distributed by kaikki.org, filtered to
  `lang_code == "en"` and to a 50,000-word common-vocabulary list. A single
  getter, `Get`.
- **Curated path (`Phonetic`, `Definitions[]`, `Synonyms[]` for ~30 words):**
  a separate, unmerged getter, `GetCurated`.
- **`synonyms`:** served by `pkg/thesaurus`, backed by OEWN, out of scope
  here.
- **Storage:** both datasets embedded in the binary via `//go:embed` — the
  Wiktionary dataset gzip-compressed, the small curated dataset not — not a
  database, consistent with `pkg/thesaurus` and `pkg/sentiment`.
- **No merge logic in this package.** Which source wins for a given word,
  and how the two response shapes are reconciled, is a calling-service
  decision, not one this package makes.

---

## Storage: `//go:embed` with gzip compression

The dataset is read-only, changes infrequently, and sits in the hot path of
every lookup call. Two artifacts are embedded side by side:

```text
Build time (once):
  dictionary.json  →  gzip  →  dataset/dictionary.json.gz   (Wiktionary, embedded via //go:embed)
  curated.json     →  (no compression; small)               (hand-curated, embedded via //go:embed)

Process startup (once):
  dataset/dictionary.json.gz  →  gzip.NewReader  →  json.Unmarshal  →  in-memory map
  dataset/curated.json        →  json.Unmarshal                    →  in-memory map

Request time (every lookup):
  in-memory map access — no decompression, no I/O, no network
```

---

## Architecture

This package's responsibility is scoped to **dataset availability only** —
it embeds and exposes both the Wiktionary-derived and hand-curated data, and
performs no lookup, normalization, selection, or merge logic itself. That
includes not resolving dialect (UK/US/Other), not picking a sense when a
word has several, and not deciding whether a given word should be served
from the curated set, the Wiktionary set, both, or neither for a particular
API response — those decisions belong to the calling service:

```text
Client
  ↓  HTTP  GET /v1/text/words?word=...
Go API (requiems-api)
  ↓  service.go: normalizes input, queries all relevant packages, composes response
  │
  ├──  calls  go-dictionary (this package — data access only, two unmerged getters)
  │      ├── Get(word string) (Entry, bool)
  │      │      → { Word, Variants[]{ Etymology, PhoneticUK, PhoneticUS,
  │      │           PhoneticOther, Definitions[]{PartOfSpeech, Definition,
  │      │           Examples[]}, SenseCount } }
  │      ├── GetCurated(word string) (CuratedEntry, bool)
  │      │      → { Phonetic, Definitions[]{PartOfSpeech, Definition, Example}, Synonyms[] }
  │      ├── dataset/dictionary.json.gz   (Wiktionary-derived, embedded via //go:embed)
  │      ├── dataset/curated.json         (hand-curated, embedded via //go:embed)
  │      └── loadData()                    (decompress/unmarshal both, once at startup)
  │
  └──  calls  pkg/thesaurus (existing, unchanged, backed by OEWN)
         └── Lookup(word string) (Entry, bool)   ← synonyms only
  ↓
service.go decides source precedence and shape, assembles the final response
  ↑
Go API  →  JSON response  →  Client
```

The package's public surface is intentionally two independent getters
(`Get`, `GetCurated`), each with its own shape, plus a richer nested
structure (`Variant`/`Definition`) inside `Get` to accommodate multiple
etymologies, multiple dialects, and multiple examples per sense without
collapsing any of them inside the package itself.

---

## Package Structure

```text
go-dictionary/
  data.go                    ← package source, go:embed directives live here
  go.mod
  dataset/                   ← EMBEDDED via go:embed — consumed at runtime
    curated.json             ← hand-curated dataset
    dictionary.json.gz       ← generated Wiktionary-derived dataset
  data/                       ← NOT embedded — build-time inputs to the generator only
    en_50k.txt                 (hermitdave/FrequencyWords wordlist)
    raw-wiktextract-data.jsonl (decompressed raw wiktextract dump, ~23GB — not committed)
  cmd/
    datasetbuild/             ← the generator that produces dataset/dictionary.json.gz
      main.go
      main_test.go

pkg/thesaurus/                ← existing, unchanged, backed by OEWN, see companion document

apps/api/services/text/words/
    service.go                ← owns Lookup(): normalizes input, calls
                                 go-dictionary.Get() and/or .GetCurated() and
                                 pkg/thesaurus.Lookup(), decides source precedence,
                                 not-found/partial-data behavior, and composes the
                                 HTTP response
```

The package intentionally exposes no unified `Lookup()` and no merge logic
of its own — only two independent per-word getters. This keeps the package a
pure data-access layer, consistent with the responsibility split already
established for this service.

---

## Known Trade-offs and Limitations

- **Recompilation required to update the dataset.** `//go:embed` bakes both
  indexes into the binary at compile time. Updating to a newer Wiktionary
  dump requires rebuilding `go-dictionary` and redeploying the API.
  Deliberate, for the same reasons already accepted for `pkg/thesaurus`.

- **Dialect (UK/US/Other) is exposed as three fields, not resolved to
  one.** `PhoneticUK`, `PhoneticUS`, and `PhoneticOther` are independent,
  possibly-empty fields with no package-internal fallback rule between
  them; any UK-preference or fallback behavior is the calling service's
  decision. `PhoneticOther` does not distinguish "no dialect tag at all"
  from "a real third dialect (e.g. Australian) that wasn't specifically
  requested" — both land in the same field.

- **Sense selection is left to the caller.** The package returns the full
  `Definitions[]` list (deduplicated by gloss text, up to 5 examples each)
  per `Variant`, plus `SenseCount`, and does not pick one for the caller —
  this affects a meaningful share of common-vocabulary entries, since
  multiple senses per word are common, not an edge case.

- **Wiktionary entries and the hand-curated words are likely to diverge
  stylistically.** Because the two sources are exposed unmerged rather than
  reconciled inside the package, any consumer that wants a single
  consistent voice per word needs to apply its own normalization or
  precedence rule when combining `Get` and `GetCurated` results.

- **`sounds[]` can bleed across etymologies in the source data.** For words
  with multiple etymologies, wiktextract sometimes replicates the full
  pronunciation list across etymology sections rather than scoping each
  dialect tag to the section it belongs to — observed with `"name"`, where
  the yam etymology's `Variant` shows a UK pronunciation that actually
  belongs to the identifier etymology. Not corrected — an accepted
  data-source limitation rather than an unverified heuristic fix.

- **Community-edited content risk.** Unlike OEWN, which has an editorial
  process, Wiktionary is open to public edits. Coverage is far larger, but
  individual entries carry a higher risk of being arcane, inconsistent, or
  occasionally vandalized. A sensitive-content tag filter (`vulgar`,
  `derogatory`, `offensive`, `slur`) is applied at generation time, but only
  catches what the source tagged explicitly; a broader spot-check on a
  sample of generated entries is still recommended over a full moderation
  pipeline, given the scope of this ticket.

- **English-only**, by explicit requirement — no multilingual scope is
  implied or supported by this design.

- **Raw feed requires a language filter Wiktionary itself does not apply.**
  The raw JSONL spans hundreds of languages; the generator filters on
  `lang_code == "en"` explicitly.

- **Etymology cleanup is a heuristic**, not a guaranteed parse — only
  "Etymology tree" and "Cognates" noise blocks were confirmed and handled;
  other formats may still leak through unclean.

---

## Future Upgrade Path

`dictionary.json.gz` is generated from entries filtered to a 50,000-word
common-vocabulary list (`hermitdave/FrequencyWords`, `en_50k.txt`), not from
the full 1.48M-entry raw feed. This is not just a size optimization — it
directly improves data quality, since coverage for `ipa` and `example` is
several times better within common vocabulary than across the full
long-tail dataset.

If real-world lookups against the deployed service reveal words outside
that filtered vocabulary are being requested and missed, the isolated
fallback point is `cmd/datasetbuild/main.go` — the wordlist can be widened
further, or a supplemental source layered in, without touching `data.go`'s
getter signatures, the calling service's `Lookup()`, or the public API
shape. OEWN remains a candidate for such a supplemental role given its 100%
`Definition` coverage, should a hybrid approach be revisited later — but
that is explicitly not the design adopted here.

If broader multilingual support becomes a requirement, Wiktionary itself
already covers other languages in the same raw feed (filtering on a
different `lang_code`), which is a materially smaller lift than OEWN's
English-only scope would have allowed.

This migration path is clean because dataset access is encapsulated inside
`go-dictionary`, kept separate from both `pkg/thesaurus` and from the
lookup/composition logic in `service.go`. The transport layer, validation,
and documentation remain agnostic to the underlying data engine.

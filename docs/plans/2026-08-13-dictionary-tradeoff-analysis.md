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
  `{ Word, PhoneticUK, PhoneticUS, PhoneticOther, Variants[]{ Etymology,
  Definitions[]{PartOfSpeech, Definition, Examples[]}, SenseCount } }` for
  the Wiktionary-derived broad-coverage set. Phonetics are word-scoped
  (one set per `Entry`), not etymology-scoped (per `Variant`) — see the
  note under Key implementation decisions for why.

Deciding source precedence for a given word, and reconciling the two shapes
into a single API response, is left to the calling service
(`apps/api/services/text/words/service.go`) — this package's own
responsibility is scoped to data access only, with no merge or selection
logic (see Architecture, below).

**Mapping to the existing service.** The calling service currently exposes
`Service.Define` behind `/dictionary/{word}`, returning the existing
`DictionaryEntry` shape. This document's `Entry`/`CuratedEntry` types are
not a drop-in replacement for `DictionaryEntry` — `Service.Define` (or its
replacement) is responsible for: normalizing the input word to lowercase
before calling `Get`/`GetCurated` (both getters expect already-lowercased
input); deciding not-found behavior when neither getter returns a result;
choosing which `Variant` (if `Get` returns more than one) and which
`Definitions[]` entries to surface within `DictionaryEntry`'s existing
fields; and deciding how `PhoneticUK`/`PhoneticUS`/`PhoneticOther` map onto
`DictionaryEntry`'s single `Phonetic` field. Whether this mapping is added
directly inside `Service.Define`, or `Service.Define` is renamed/replaced
as part of a broader `/dictionary/{word}` → `/v1/text/words?word=...`
migration, is an open implementation question this document does not
resolve — it is called out here so the generated dataset isn't shipped
without a defined path to becoming reachable through the API.

The `synonyms` field is out of scope for this document: it is served by
`pkg/thesaurus`, backed by OEWN, and is not re-sourced here.

**Synonyms precedence, clarified.** `CuratedEntry.Synonyms` is not
redundant with `pkg/thesaurus` — the two are not automatically
reconciled, and `Service.Define`'s current behavior (returning the
curated word's own `e.synonyms` when the word is in the curated set) is
preserved by this design, not changed by it. `Get`'s `Entry`/`Variant`
types carry no `Synonyms` field at all — Wiktionary-derived words rely
entirely on `pkg/thesaurus.Lookup()` for synonyms. The precedence is: for
curated words, `Service.Define` uses `CuratedEntry.Synonyms` as-is (matching
existing test expectations, e.g. `TestDictionary_KnownWord`); for
Wiktionary-only words, `Service.Define` falls back to
`pkg/thesaurus.Lookup()`. No merge between the two sources is performed for
a single word — this avoids silently changing the existing `Synonyms`
response field's compatibility contract for the ~30 curated words already
covered.

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
  underlying *data* it returns. Note: the project's own repository is
  GPL-3.0-licensed, but that covers the API server's source code, not the
  dictionary content it serves — the two are legally distinct, and the
  data's own provenance and license remain undocumented. This distinction,
  and the resulting uncertainty, is itself the diligence gap that rules
  this option out for an attributed, embedded dataset.
- Quality is inconsistent across words: some entries are rich, others sparse
  or missing fields entirely, with no way to bulk-audit coverage ahead of a
  build.

### 3. Merriam-Webster Dictionary API

A commercial dictionary API with authoritative definitions, etymologies,
audio pronunciations, and specialized (medical, Spanish, ESL) vocabularies.

**Rejected because:** the free tier is licensed for non-commercial use only
and capped at 1,000 queries/day, with restrictions on storing or
redistributing returned content beyond the immediate request — none of
which fit an embedded, offline dataset serving unlimited internal lookups.
Commercial use — which this service is — requires a paid license
agreement, introducing an ongoing per-word or per-call cost and a vendor
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
  unlike OEWN's single unlabeled transcription. When multiple pronunciations
  share the same dialect tag, the generator keeps the first one encountered
  in source order and discards the rest — a deliberate simplification, not
  an omission. Source ordering is stable for a given pinned dump (see
  `source.lock`), so this does not introduce nondeterminism across rebuilds
  of the same input, though it does mean two equally valid same-dialect
  transcriptions are not both preserved.
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
  as updated regularly (at least weekly), which is exactly why a pinning
  mechanism is required rather than optional: without one, two builds run
  on different days could silently pull different dumps and produce
  different datasets. This is handled by a `source.lock` file, generated
  via `cmd/datasetbuild -generate-lock` and committed at the repo root
  alongside `go.mod`. It records the dump's source URL, the fetch
  timestamp, and its SHA-256 hash. Every normal build run
  (`cmd/datasetbuild -in ... -wordlist ... -out ...`) verifies the input
  dump's hash against `source.lock` before generating anything, and aborts
  with a descriptive error on mismatch — the same role `go.sum` plays for
  dependencies, applied to the data input instead. `-skip-verify` bypasses
  this for local experimentation but is not intended for reproducible or
  production builds.
- **Dump distribution strategy.** The raw dump (~23GB) is deliberately not
  committed to the repository, but it also isn't left to each developer to
  independently discover and download from kaikki.org. It is mirrored to an
  internal S3 bucket (`s3://<bucket>/wiktextract/raw-wiktextract-data-<date>.jsonl.gz`)
  at the same cadence `source.lock` is refreshed, so the URL recorded in
  `source.lock` and the internal mirror always point at the same pinned
  version. `cmd/datasetbuild` does not fetch this automatically — developers
  pull it via the team's existing S3 tooling (or `aws s3 cp`) before running
  the generator (`go run ./cmd/datasetbuild -in ... -wordlist ... -out ...`).
  This keeps onboarding to "sync the bucket, run the generator" instead of
  "find and re-download a 23GB file from a public mirror that updates
  weekly," and keeps the source used across the team consistent with what
  `source.lock` pins.

**Licensing decision:** the ShareAlike obligation of CC-BY-SA was evaluated
and the risk was accepted on the understanding that the resulting dataset is
served only through the package's own API and never distributed as a
standalone downloadable artifact, dataset export, or open-source repository
independent of attribution. This is recorded in the package's `Credits`
section, citing CC-BY-SA, GFDL, and the wiktextract tool (MIT-licensed,
separate from the license of the data it extracts).

**Key implementation decisions:**

- **Dialect is exposed, not selected — and modeled at word scope, not
  etymology scope.** Rather than collapsing a word's pronunciations down
  to a single `phonetic` value, the design keeps three separate fields —
  `PhoneticUK`, `PhoneticUS`, `PhoneticOther` — and makes no selection
  between them; the calling service decides what to surface. These fields
  live on `Entry` (one set per word), not on `Variant` (one set per
  etymology): an empirical check against the 50k-word wordlist found that
  wiktextract duplicates the identical `sounds[]` array across etymology
  sections in 85.2% of multi-etymology words (3,507 of 4,115), rather than
  genuinely scoping pronunciation per meaning. Modeling phonetics per
  `Variant` would therefore misattribute a dialect transcription to the
  wrong sense in the large majority of cases it applies to at all; modeling
  it per `Entry` reflects what the source data actually supports. See
  Known Trade-offs for the corresponding limitation. An inventory pass
  over `sounds[].tags` across the wordlist-filtered corpus found that the
  bare `US`/`UK` tags alone undercounted real coverage: the tags
  `General-American` and `Received-Pronunciation`/`British` carry
  substantial additional volume and are recognized alongside `US`/`UK`.
  Everything else (Australian, Canadian, Scottish, register qualifiers like
  "dialectal" or "archaic") collapses into `PhoneticOther` as a single
  value — a multi-dialect map beyond UK/US/Other was considered and
  explicitly rejected as out of scope.
- **Senses are exposed as a full list, not selected down to one.** The
  design returns the full `Definitions []Definition` list per `Variant`
  (each with up to 5 deduplicated `Examples`), plus a `SenseCount` field as
  a signal for callers that want to choose one. Definitions are
  deduplicated by `partOfSpeech` + exact gloss text together, not gloss
  alone — a `Variant` can mix multiple parts of speech (it is grouped by
  etymology, not by pos), so two senses can legitimately share identical
  gloss text while being genuinely different senses (e.g. "3rd" as an
  abbreviated adjective vs. verb). Deduplicating on gloss alone would
  silently drop one. This was verified empirically against the 50k-word
  wordlist: 434 word/etymology groups (out of 37,357 matched words) had
  identical gloss text across different parts of speech — confirming this
  is a real, if infrequent (~1.2%), case worth correcting rather than a
  theoretical one. Within each gloss group, examples of type `"example"`
  are preferred over type `"quotation"`, deduplicated by exact text.

  **Known limitation, deliberately not addressed:** deduplication does not
  additionally key on sense-level tags. The same empirical check found
  1,798 cases of identical gloss + partOfSpeech with differing tags — a
  higher count than the pos case above, but inspection of a sample showed
  the large majority are grammatical-usage tags (`countable`/`uncountable`,
  `transitive`/`intransitive`) that wiktextract records once per valid
  usage combination rather than genuinely distinct senses, where collapsing
  is the correct behavior. A smaller subset (e.g. regional tags like `UK`
  vs. `Canada`/`US`, or register tags like `dialectal`/`obsolete`/`archaic`)
  does represent real semantic distinctions that get merged away as a
  result. Given this signal-to-noise ratio, a tags-aware dedup key was
  rejected: it would reintroduce far more false "duplicates" (from
  grammatical-usage tag combinations) than genuine distinctions it would
  preserve, and doing it properly would additionally require a new
  `Labels` field on `Definition` and a corresponding change to how
  `Service.Define` maps onto `DictionaryEntry` — out of scope for this
  ticket.
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
  `en_50k.txt`, 37,357 words matched — this is the actual denominator, not
  the raw dump's millions of entries), coverage was re-measured using the
  **same classifier the real generator uses** for dialect (`US`/
  `General-American` → US; `UK`/`Received-Pronunciation`/`British` → UK),
  and each field kept **separate** rather than averaged into one figure:

  | Field | Coverage (50k wordlist, 37,357 words) | OEWN equivalent |
  | --- | --- | --- |
  | `definition` | 100.0% (37,357 / 37,357) | 100% |
  | `example` | 67.6% (25,263 / 37,357) | 27.9% |
  | any `ipa` present (any tag or untagged) | 66.2% (24,727 / 37,357) | — |
  | `ipa` tagged with a dialect the real classifier recognizes (UK or US) | 37.6% (14,054 / 37,357) | 27.6% |

  Wiktionary outperforms OEWN on every field measured, though by a smaller
  margin on phonetics than an earlier, methodologically flawed pass of
  this analysis reported. That earlier pass combined `definition`,
  `example`, and phonetic coverage into a single "60.6%" figure and
  compared it directly against OEWN's phonetic-only rate — comparing a
  blended metric to a single-field one. It also counted only the bare
  literal `"US"`/`"UK"` tags (13.8% of matched words) rather than the
  classifier the generator actually uses, undercounting real phonetic
  coverage before the blending error compounded it further upward. The
  conclusion (Wiktionary beats OEWN on every measured field, for the
  vocabulary this service actually serves) is unchanged; the number
  supporting the phonetics claim specifically is not 60.6% — it is 37.6%,
  still comfortably above OEWN's 27.6%.

  The low global numbers seen in the first pass above are a long-tail
  artifact: Wiktionary documents roughly 1.48M English entries versus
  OEWN's 157,513 lemmas, and that ~9x larger vocabulary is overwhelmingly
  rare, technical, or archaic terms that Wiktionary's editors have defined
  but rarely transcribed phonetically or illustrated with a usage example.
  The service being built here will only ever be queried for common
  vocabulary, so the global figure measures the wrong population.

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
to the problem. This comparison is about *runtime* overhead specifically —
the embedded approach still carries its own build-time cost (see the note
under Decision Summary), just not one that is paid on every request or
requires a running service to maintain.

---

## Key Finding: Global Coverage Understated the Real Picture — Verification Against a Common-Vocabulary Filter Was Required

**First pass (global, all 1,481,704 English entries): the numbers looked
worse than OEWN, not better.** `ipa` coverage was 9.3%, roughly a third of
OEWN's 27.6%, and `example` coverage was 19.4%, below OEWN's 27.9% — the
opposite of what would justify moving off OEWN. Taken at face value, this
would have called the whole decision into question.

**Second pass (restricted to a 50,000-word common-English list): the
numbers inverted.** Measured separately per field with the classifier the
generator actually uses, `example` coverage rose to 67.6% and phonetic
(UK/US) coverage to 37.6% — both above OEWN's equivalents (27.9% and
27.6% respectively). The explanation is straightforward: Wiktionary's ~9x
larger vocabulary is overwhelmingly long-tail, and the service being built
here will only ever be queried for common vocabulary, so the global figure
was measuring the wrong population entirely.

**A correction to this finding's own first draft:** an earlier version of
this section reported a single "60.6%" figure, produced by averaging
`definition`, `example`, and phonetic coverage together and comparing that
blend against OEWN's phonetic-only rate — not a valid comparison, and it
also undercounted phonetics by classifying only the literal `US`/`UK`
tags rather than the full set the generator recognizes
(`General-American`, `Received-Pronunciation`, `British`). Re-measured
correctly, the underlying conclusion holds — Wiktionary still beats OEWN
on every field, for the vocabulary this service serves — but the
phonetics margin is real and moderate (37.6% vs. 27.6%), not the
inflated one the blended figure implied. See the Field-coverage
completeness bullet under Alternative 5 for the full corrected table.

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
| `phonetic` UK/US/Other coverage          | —                 | ⚠️ inconsistent      | ✅               | ❌ single unlabeled form, 27.6% coverage | ✅ 37.6% on 50k common-vocab list, real dialect classifier (verified; 9.3% unrestricted, any-tag) | ❌ US only | — |
| `definition` coverage                    | —                 | ⚠️ inconsistent      | ✅               | ✅ 100% (verified) | ✅ 99.9% (verified) | ❌ none | — |
| `example` coverage                       | —                 | ⚠️ inconsistent      | ✅               | ⚠️ 27.9% (verified) | ✅ 67.6% on 50k common-vocab list (verified; 19.4% unrestricted) | ❌ none | — |
| Reuses existing pipeline/`Provider`      | —                 | ❌                    | ❌               | ✅ (but insufficient alone) | ❌ new provider needed | ❌ | — |
| No *runtime* infrastructure overhead     | ✅               | ✅                    | ✅               | ✅   | ✅ (runtime); ⚠️ build-time: ~23GB disk, generator run required | ✅  | ❌ (DB ops)       |
| Independent update cycle                 | ❌               | ✅                    | ✅               | ✅   | ✅ (weekly dumps, documented) | ✅  | ✅                |
| In scope for this package                | ✅ retained as a second source | rejected | rejected | ❌ out of scope, remains sole source for `pkg/thesaurus` | ✅ selected as the broad-coverage source | rejected | rejected |

**A note on "no infrastructure overhead."** This criterion means no
*runtime* infrastructure — no database, no external API call, no network
dependency once the binary is built. It does not mean the Wiktionary path
is free. Build-time requirements are real and worth stating explicitly:
~23GB of local disk space for the decompressed raw dump (see Dump
distribution strategy), a full pass over ~10.7M JSONL lines by
`cmd/datasetbuild` (single-pass streaming, no significant memory beyond
the in-progress word/etymology maps), and enough CPU/time to complete that
pass before `dictionary.json.gz` can be regenerated. None of this affects
the running API — it is a one-time (or per-dump-update) cost paid by
whoever runs the generator, not by every request.

**Final decision:**

- **Broad-coverage path (`etymology`, `phonetic` UK/US/Other,
  `partOfSpeech`, `definition`, `example`):** Wiktionary, via raw
  `wiktextract` JSONL data distributed by kaikki.org, filtered to
  `lang_code == "en"` and to a 50,000-word common-vocabulary list. A single
  getter, `Get`.
- **Curated path (`Phonetic`, `Definitions[]`, `Synonyms[]` for ~30 words):**
  a separate, unmerged getter, `GetCurated`. For these words,
  `Service.Define` uses `CuratedEntry.Synonyms` as-is — see the precedence
  rule in Context.
- **`synonyms` for Wiktionary-only words:** served by `pkg/thesaurus`,
  backed by OEWN. `Get`'s `Entry`/`Variant` types carry no `Synonyms` field;
  `Service.Define` falls back to `pkg/thesaurus.Lookup()` for words outside
  the curated set. See the precedence rule in Context.
- **Storage:** both datasets embedded in the binary via `//go:embed` — the
  Wiktionary dataset gzip-compressed, the small curated dataset not — not a
  database, consistent with `pkg/thesaurus` and `pkg/sentiment`.
- **No merge logic in this package.** Which source wins for a given word,
  and how the two response shapes are reconciled, is a calling-service
  decision, not one this package makes.

---

## Storage: `//go:embed` with gzip compression

The dataset is read-only, changes infrequently, and sits in the hot path of
every lookup call. Two artifacts are embedded side by side. This is where
the "no runtime infrastructure overhead" criterion from Decision Summary is
implemented concretely — the cost shown in the "Build time (once)" row
below is real, but it is paid once by whoever runs the generator, not on
every request:

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
  ↓  HTTP  GET /dictionary/{word}
Go API (requiems-api)
  ↓  Service.Define: normalizes input, queries all relevant packages, composes response
  │
  ├──  calls  go-dictionary (this package — data access only, two unmerged getters)
  │      ├── Get(word string) (Entry, bool)
  │      │      → { Word, PhoneticUK, PhoneticUS, PhoneticOther,
  │      │           Variants[]{ Etymology, Definitions[]{PartOfSpeech,
  │      │           Definition, Examples[]}, SenseCount } }
  │      ├── GetCurated(word string) (CuratedEntry, bool)
  │      │      → { Phonetic, Definitions[]{PartOfSpeech, Definition, Example}, Synonyms[] }
  │      ├── dataset/dictionary.json.gz   (Wiktionary-derived, embedded via //go:embed)
  │      ├── dataset/curated.json         (hand-curated, embedded via //go:embed)
  │      └── loadData()                    (decompress/unmarshal both, once at startup)
  │
   └──  calls  pkg/thesaurus (existing, unchanged, backed by OEWN)
         └── Lookup(word string) (Entry, bool)   ← synonyms for Wiktionary-only
                                                     words; curated words use
                                                     CuratedEntry.Synonyms instead
                                                     — see precedence rule in Context
  ↓
Service.Define decides source precedence and shape, assembles the final
DictionaryEntry response
  ↑
Go API  →  JSON response  →  Client
```

The package's public surface is intentionally two independent getters
(`Get`, `GetCurated`), each with its own shape, plus a richer nested
structure (`Variant`/`Definition`) inside `Get` to accommodate multiple
etymologies and multiple examples per sense without collapsing any of
them inside the package itself. Dialects (`PhoneticUK`/`PhoneticUS`/
`PhoneticOther`) are word-scoped on `Entry` rather than per-`Variant` —
see Key implementation decisions for why.

---

## Package Structure

```text
go-dictionary/
  data.go                    ← package source, go:embed directives live here
  go.mod
  source.lock                ← pins the dump used to generate dataset/dictionary.json.gz:
                                 source URL, fetch timestamp, SHA-256 — see Operational
                                 status under Alternative 5
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
    service.go                ← owns Service.Define, serving /dictionary/{word}:
                                 normalizes input to lowercase, calls
                                 go-dictionary.Get() and/or .GetCurated() and
                                 pkg/thesaurus.Lookup(), decides source precedence,
                                 not-found/partial-data behavior, and maps the
                                 result onto the existing DictionaryEntry shape
                                 (see mapping note in Context)
```

The package intentionally exposes no unified lookup function and no merge
logic of its own — only two independent per-word getters (`Get`,
`GetCurated`). This keeps the package a pure data-access layer; combining
them into a single result, as `Service.Define` currently does for
`DictionaryEntry`, is the calling service's responsibility, not this
package's.

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

- **Phonetics are word-scoped, not etymology-scoped — a source limitation,
  not a package one.** `PhoneticUK`/`PhoneticUS`/`PhoneticOther` live on
  `Entry`, describing the word's pronunciation as a whole, not on `Variant`
  per etymology. This was a deliberate correction, first prompted by
  observing `"name"` (where the yam etymology showed a UK pronunciation
  that actually belongs to the identifier etymology) and then confirmed
  empirically: wiktextract duplicates the identical `sounds[]` array
  across etymology sections in 85.2% of multi-etymology words (3,507 of
  4,115) — the source itself does not reliably scope pronunciation per
  meaning, so a genuinely different etymology with its own distinct
  pronunciation (where the source does provide that distinction cleanly)
  is not preserved separately. A caller cannot infer which etymology/sense
  a given `PhoneticUK`/`US`/`Other` value "belongs to," because the source
  data does not support that distinction in the large majority of cases.
  Extraction now scans every raw entry for the word (across all
  `etymology_number` groups, not just one), which widens the "first in
  source order wins" rule from the Phonetics field note under Alternative
  5 to the whole word: the determinism guarantee still holds — stable for
  a given pinned dump (see `source.lock`) — but the pool of candidate
  sounds a first match is drawn from is now the word's entire raw entry
  set, not one etymology group's.

- **Sense selection is left to the caller.** The package returns the full
  `Definitions[]` list (deduplicated by `partOfSpeech` + gloss text, up to
  5 examples each) per `Variant`, plus `SenseCount`, and does not pick one
  for the caller — this affects a meaningful share of common-vocabulary
  entries, since multiple senses per word are common, not an edge case.

- **Wiktionary entries and the hand-curated words are likely to diverge
  stylistically.** Because the two sources are exposed unmerged rather than
  reconciled inside the package, any consumer that wants a single
  consistent voice per word needs to apply its own normalization or
  precedence rule when combining `Get` and `GetCurated` results.

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

## Credits / Source Manifest

| Source | License | Notes |
| --- | --- | --- |
| Wiktionary content (via wiktextract raw data) | CC BY-SA 4.0 and GFDL (dual) | ShareAlike applies to any redistribution of the data itself; served only through this package's own API, not redistributed as a standalone dataset. |
| `wiktextract` tool | MIT | Separate from the license of the data it extracts — the tool's code license does not extend to Wiktionary content. |
| `hermitdave/FrequencyWords` (`en_50k.txt`) | [VERIFY: confirm whether this is MIT (tooling) or CC BY-SA 4.0 (underlying OpenSubtitles-derived word/count data) before publishing] | Frequency counts are derived from OpenSubtitles data; provenance and license should be confirmed at the data level, not just the repository's top-level license badge. |
| `dataset/dictionary.json.gz` (generated artifact) | Inherits CC BY-SA 4.0 / GFDL from Wiktionary | Distributed only as an embedded binary asset via this package's API, not as a standalone downloadable dataset — see Licensing decision under Alternative 5. |
| This repository's own code | [BSL — confirm exact version/date-based conversion terms] | Applies to the Go source only, not to the embedded third-party datasets above. |
| `source.lock` | N/A (metadata only) | Records the pinned dump URL, fetch date, and SHA-256 — not a license record, but supports the provenance/reproducibility obligations above. |

Any future redistribution of `dictionary.json.gz` outside this package's
own API (e.g. as a public dataset export) would need to satisfy CC BY-SA's
ShareAlike terms independently — this has not been evaluated and is out of
scope for the current embedded-only usage.

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
getter signatures, `Service.Define`, or the public API shape. OEWN remains
a candidate for such a supplemental role given its 100% `Definition`
coverage, should a hybrid approach be revisited later — but that is
explicitly not the design adopted here.

If broader multilingual support becomes a requirement, Wiktionary itself
already covers other languages in the same raw feed (filtering on a
different `lang_code`), which is a materially smaller lift than OEWN's
English-only scope would have allowed.

This migration path is clean because dataset access is encapsulated inside
`go-dictionary`, kept separate from both `pkg/thesaurus` and from the
lookup/composition logic in `Service.Define`. The transport layer,
validation, and documentation remain agnostic to the underlying data
engine.
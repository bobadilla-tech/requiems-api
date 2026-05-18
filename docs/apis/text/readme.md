# Text & Language APIs

## Overview

Text processing, language analysis, and generation endpoints. Covers everything
from spell checking and vocabulary tools to NLP (sentiment, language detection,
text similarity) and placeholder content generation.

## Endpoints

### [Spell Check](./spell-check.md) - ✅ MVP

Check spelling and get correction suggestions for misspelled words.

- **Status:** mvp
- **Endpoint:** `POST /v1/text/spellcheck`
- **Credit Cost:** 1

### [Thesaurus](./thesaurus.md) - ✅ MVP

Find synonyms and antonyms for any word.

- **Status:** mvp
- **Endpoint:** `GET /v1/text/thesaurus/{word}`
- **Credit Cost:** 1

### [Dictionary](./dictionary.md) - ✅ MVP

Get word definitions, phonetics, usage examples, and synonyms.

- **Status:** mvp
- **Endpoint:** `GET /v1/text/dictionary/{word}`
- **Credit Cost:** 1

### [Sentiment Analysis](../ai-computer-vision/sentiment.md) - ✅ MVP

Analyze sentiment (positive/negative/neutral) with confidence score.

- **Status:** mvp
- **Endpoint:** `POST /v1/text/sentiment`
- **Credit Cost:** 1

### [Language Detection](../ai-computer-vision/text-similarity.md) - ✅ MVP

Detect the language of any text with confidence scoring.

- **Status:** mvp
- **Endpoint:** `POST /v1/text/language`
- **Credit Cost:** 1

### [Text Similarity](../ai-computer-vision/text-similarity.md) - ✅ MVP

Compare two texts and get a cosine similarity score between 0 and 1.

- **Status:** mvp
- **Endpoint:** `POST /v1/text/similarity`
- **Credit Cost:** 1

### [Markdown to HTML](./markdown.md) - ✅ MVP

Convert Markdown to HTML. Optionally sanitize output to prevent XSS.

- **Status:** mvp
- **Endpoint:** `POST /v1/technology/markdown`
- **Credit Cost:** 1

### [Lorem Ipsum Generator](./lorem-ipsum.md) - ✅ MVP

Generate placeholder text for design mockups and prototypes.

- **Status:** mvp
- **Endpoint:** `GET /v1/text/lorem`
- **Credit Cost:** 1

### [Random Words](./random-word.md) - ✅ MVP

Generate random words for testing, games, and creative projects.

- **Status:** mvp
- **Endpoints:** `GET /v1/text/words/random`, `GET /v1/text/words`
- **Credit Cost:** 1

### [Email Normalizer](../email/normalize.md) - ✅ MVP

Normalize email addresses to canonical form with provider-specific rules.

- **Status:** mvp
- **Endpoints:** `POST /v1/text/normalize`, `POST /v1/text/normalize/batch`
- **Credit Cost:** 1 per email (batch charges per item)

## Category Statistics

- Total Endpoints: 10
- Live: 10

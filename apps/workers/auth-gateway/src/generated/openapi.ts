// AUTO-GENERATED — do not edit. Run `pnpm generate:openapi` to regenerate.
export const openApiSpec = {
  "openapi": "3.0.3",
  "info": {
    "title": "Requiems API",
    "version": "1.0.0",
    "description":
      "Unified access to enterprise-grade APIs — email validation, text utilities, and more. Authenticate with the `requiems-api-key` header.",
  },
  "servers": [
    {
      "url": "https://api.requiems.xyz",
      "description": "Production",
    },
  ],
  "components": {
    "securitySchemes": {
      "requiems-api-key": {
        "type": "apiKey",
        "in": "header",
        "name": "requiems-api-key",
        "description": "Your Requiems API key",
      },
    },
  },
  "security": [
    {
      "requiems-api-key": [],
    },
  ],
  "tags": [
    {
      "name": "advice",
      "description": "Get random pieces of advice and wisdom for inspiration",
    },
    {
      "name": "barcode",
      "description":
        "Generate barcodes in multiple formats (Code 128, Code 93, Code 39, EAN-8, EAN-13), returned as a PNG image or base64-encoded JSON",
    },
    {
      "name": "base64",
      "description":
        "Encode strings to Base64 and decode Base64 back to plain text. Supports standard and URL-safe (base64url) variants.",
    },
    {
      "name": "bin-lookup",
      "description":
        "Pass the first 6–8 digits of any payment card and get back the issuing bank, card network, type, and country",
    },
    {
      "name": "chuck-norris",
      "description":
        "Get a random Chuck Norris fact from a curated built-in database. Every call returns a different fact with a stable unique ID.",
    },
    {
      "name": "cities",
      "description":
        "Look up city metadata including population, timezone, and coordinates",
    },
    {
      "name": "color-conversion",
      "description":
        "Convert color values between HEX, RGB, HSL, and CMYK. Every response includes all four representations.",
    },
    {
      "name": "commodities",
      "description":
        "Historical and current annual average prices for 16 major commodities — precious metals, energy, and agricultural goods. Sourced from FRED (Federal Reserve Economic Data).",
    },
    {
      "name": "counter",
      "description": "Atomic, namespace-isolated hit counter",
    },
    {
      "name": "crypto",
      "description":
        "Get live cryptocurrency prices, 24h change, market cap, and trading volume for 20+ major coins including BTC, ETH, SOL, and more.",
    },
    {
      "name": "dad-jokes",
      "description":
        "Get a random dad joke. Classic groan-worthy puns and wholesome humor, served one at a time.",
    },
    {
      "name": "data-format-conversion",
      "description":
        "Convert structured data between JSON, YAML, CSV, XML, and TOML in a single API call.",
    },
    {
      "name": "detect-language",
      "description": "Detect the language of any text with confidence scoring",
    },
    {
      "name": "dictionary",
      "description":
        "Get word definitions, phonetics, usage examples, and synonyms",
    },
    {
      "name": "disposable_email",
      "description":
        "Detect disposable and temporary email addresses to prevent fraud and improve data quality. Our comprehensive blocklist is continuously updated to catch the latest disposable email providers.",
    },
    {
      "name": "domain-info",
      "description":
        "Look up DNS records and check domain availability. Returns A, AAAA, MX, NS, TXT, and CNAME records alongside a registration availability flag.",
    },
    {
      "name": "email-normalize",
      "description":
        "Normalize email addresses to their canonical form. Lowercased, trimmed, and canonicalized with provider-specific rules including alias-domain resolution. Also available as part of the full Email Validator.",
    },
    {
      "name": "email-validate",
      "description":
        "Full email validation in one call. Syntax check, MX record lookup, disposable domain detection, normalization, and typo suggestions. Includes everything from the Disposable Domain Checker and Email Normalizer in a single request.",
    },
    {
      "name": "emoji",
      "description":
        "Look up emoji by name, search by keyword, or get a random emoji with full Unicode metadata.",
    },
    {
      "name": "exchange-rate",
      "description":
        "Get live currency exchange rates and convert amounts between currencies. Rates are sourced from the ECB and cached for up to one hour.",
    },
    {
      "name": "facts",
      "description":
        "Get random interesting facts from a curated database. Filter by category — science, history, technology, nature, space, or food. Includes source attribution.",
    },
    {
      "name": "fitness-exercises",
      "description":
        "Browse 1,500+ exercises with step-by-step instructions, target muscles, secondary muscles, equipment requirements, and body part filters.",
    },
    {
      "name": "geocode",
      "description":
        "Convert addresses to coordinates and coordinates back to addresses",
    },
    {
      "name": "holidays",
      "description": "Get a list of holidays for a specific country and year",
    },
    {
      "name": "horoscope",
      "description": "Get daily horoscope readings for all 12 zodiac signs",
    },
    {
      "name": "iban",
      "description":
        "Validate IBAN numbers and extract the bank code and account number. Supports all countries in the official SWIFT IBAN Registry (~80 countries).",
    },
    {
      "name": "inflation",
      "description":
        "Historical and current CPI inflation rates for 241 countries, sourced from the World Bank. Includes up to 30 years of annual data.",
    },
    {
      "name": "ip-asn",
      "description":
        "Look up Autonomous System Number (ASN), organization, ISP, and network route information for any IP address.",
    },
    {
      "name": "ip-info",
      "description":
        "Get geolocation data for any IP address including country, city, ISP, and VPN detection.",
    },
    {
      "name": "lorem-ipsum",
      "description":
        "Generate placeholder text for design mockups and prototypes",
    },
    {
      "name": "markdown",
      "description":
        "Convert Markdown to HTML in a single API call. Optionally sanitize the output to strip unsafe tags and prevent XSS.",
    },
    {
      "name": "mortgage",
      "description":
        "Calculate monthly mortgage payments and full amortization schedules. Pass principal, annual interest rate, and loan term in years.",
    },
    {
      "name": "mx-lookup",
      "description":
        "Look up MX (Mail Exchange) records for any domain. Returns all mail server hostnames with their priorities, sorted from highest to lowest priority.",
    },
    {
      "name": "number-base-conversion",
      "description":
        "Convert integers between binary, octal, decimal, and hexadecimal. Accepts optional 0x, 0b, and 0o prefixes.",
    },
    {
      "name": "password-generator",
      "description":
        "Generate cryptographically secure random passwords with customizable length and character sets",
    },
    {
      "name": "phone-validation",
      "description":
        "Validate phone numbers globally. Detect carrier, country, type, and VOIP or virtual risk using only phone metadata.",
    },
    {
      "name": "postal-code",
      "description":
        "Look up city, state, and coordinates for any postal code worldwide",
    },
    {
      "name": "profanity",
      "description":
        "Detect and censor profanity in text for content moderation",
    },
    {
      "name": "qr-code",
      "description":
        "Generate QR codes from any text or URL, returned as a PNG image or base64-encoded JSON",
    },
    {
      "name": "quotes",
      "description": "Access a database of inspirational and famous quotes",
    },
    {
      "name": "random-user",
      "description":
        "Generate random fake user profiles for testing and prototyping — names, emails, phone numbers, addresses, and avatars",
    },
    {
      "name": "sentiment",
      "description":
        "Analyze the sentiment of any text and get a positive, negative, or neutral classification with a confidence score and full class breakdown.",
    },
    {
      "name": "spell-check",
      "description":
        "Check spelling and get correction suggestions for misspelled words",
    },
    {
      "name": "sudoku",
      "description":
        "Generate Sudoku puzzles with solutions across multiple difficulty levels",
    },
    {
      "name": "swift-code",
      "description":
        "Validate and look up bank information by SWIFT/BIC code, including institution, location, and branch metadata.",
    },
    {
      "name": "text-similarity",
      "description":
        "Compare two texts and get a cosine similarity score between 0 and 1",
    },
    {
      "name": "thesaurus",
      "description":
        "Find synonyms and antonyms for any word to enhance vocabulary and writing",
    },
    {
      "name": "timezone",
      "description":
        "Get timezone information for any location by coordinates or city name",
    },
    {
      "name": "trivia",
      "description":
        "Get random trivia questions with multiple-choice answers. Filter by category (science, history, geography, sports, music, movies, literature, math, technology, nature) and difficulty (easy, medium, hard).",
    },
    {
      "name": "unit-conversion",
      "description":
        "Convert between units of measurement — length, weight, volume, temperature, area, and speed",
    },
    {
      "name": "useragent",
      "description":
        "Parse user agent strings to extract browser, OS, device type, and bot detection",
    },
    {
      "name": "vpn-detection",
      "description":
        "Detect if an IP address belongs to a VPN, proxy, Tor exit node, or hosting provider. Returns threat scores and fraud indicators for fraud prevention and risk assessment.",
    },
    {
      "name": "whois",
      "description":
        "Get domain registration details including registrar, name servers, status, creation and expiry dates, and DNSSEC information.",
    },
    {
      "name": "random-word",
      "description":
        "Get random words with definitions and parts of speech. Perfect for vocabulary builders, educational apps, word games, or content inspiration.",
    },
    {
      "name": "working-days",
      "description":
        "Calculate the number of working days between two dates with optional country-specific holidays",
    },
    {
      "name": "world-time",
      "description": "Get the current time for any IANA timezone by name",
    },
  ],
  "paths": {
    "/v1/text/advice": {
      "get": {
        "summary": "Get Random Advice",
        "tags": [
          "advice",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description": "Returns a random piece of advice",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "id": {
                          "type": "integer",
                          "description": "Unique identifier for the advice",
                        },
                        "advice": {
                          "type": "string",
                          "description": "A random piece of advice",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "id": 42,
                    "advice":
                      "Don't compare yourself to others. Compare yourself to the person you were yesterday.",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
        },
      },
    },
    "/v1/technology/barcode": {
      "get": {
        "summary": "Generate Barcode (PNG)",
        "tags": [
          "barcode",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a raw PNG image of the barcode. Ideal for direct embedding or file download.",
        "parameters": [
          {
            "name": "data",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": 123456789,
            },
            "description":
              "The text or numeric string to encode in the barcode",
          },
          {
            "name": "type",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "code128",
            },
            "description":
              "Barcode format: code128, code93, code39, ean8, ean13",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
          },
          "400": {
            "description":
              "Missing or invalid parameters (e.g. data not provided, unsupported type)",
          },
          "422": {
            "description":
              "Data is invalid for the chosen barcode type (e.g. wrong digit count for EAN-8/EAN-13, non-numeric EAN data)",
          },
        },
      },
    },
    "/v1/technology/barcode/base64": {
      "get": {
        "summary": "Generate Barcode (Base64 JSON)",
        "tags": [
          "barcode",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a JSON envelope containing the barcode as a base64-encoded PNG string, along with its type and dimensions.",
        "parameters": [
          {
            "name": "data",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": 123456789,
            },
            "description":
              "The text or numeric string to encode in the barcode",
          },
          {
            "name": "type",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "code128",
            },
            "description":
              "Barcode format: code128, code93, code39, ean8, ean13",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "image": {
                          "type": "string",
                          "description": "Base64-encoded PNG image data",
                        },
                        "type": {
                          "type": "string",
                          "description": "The barcode format that was used",
                        },
                        "width": {
                          "type": "integer",
                          "description":
                            "Width of the generated image in pixels",
                        },
                        "height": {
                          "type": "integer",
                          "description":
                            "Height of the generated image in pixels",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "image": "<base64-encoded PNG data>",
                    "type": "code128",
                    "width": 300,
                    "height": 100,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "Missing or invalid parameters (e.g. data not provided, unsupported type)",
          },
          "422": {
            "description":
              "Data is invalid for the chosen barcode type (e.g. wrong digit count for EAN-8/EAN-13, non-numeric EAN data)",
          },
        },
      },
    },
    "/v1/technology/base64/encode": {
      "post": {
        "summary": "Encode",
        "tags": [
          "base64",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description": "Encode a plain-text string to Base64",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "value": {
                    "type": "string",
                    "description": "The string to encode",
                    "example": "Hello, world!",
                  },
                  "variant": {
                    "type": "string",
                    "description":
                      "Encoding variant: standard (default) or url (URL-safe base64url)",
                    "example": "standard",
                  },
                },
                "required": [
                  "value",
                ],
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "result": {
                          "type": "string",
                          "description": "The Base64-encoded output",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "result": "SGVsbG8sIHdvcmxkIQ==",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "Missing or empty value field",
          },
          "422": {
            "description":
              "Validation constraint on the variant field (must be standard or url)",
          },
        },
      },
    },
    "/v1/technology/base64/decode": {
      "post": {
        "summary": "Decode",
        "tags": [
          "base64",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description": "Decode a Base64-encoded string back to plain text",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "value": {
                    "type": "string",
                    "description": "The Base64-encoded string to decode",
                    "example": "SGVsbG8sIHdvcmxkIQ==",
                  },
                  "variant": {
                    "type": "string",
                    "description":
                      "Encoding variant: standard (default) or url (URL-safe base64url)",
                    "example": "standard",
                  },
                },
                "required": [
                  "value",
                ],
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "result": {
                          "type": "string",
                          "description": "The decoded plain-text output",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "result": "Hello, world!",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "Missing or empty value field",
          },
          "422": {
            "description":
              "The value is not valid Base64 and cannot be decoded",
          },
        },
      },
    },
    "/v1/finance/bin/{bin}": {
      "get": {
        "summary": "BIN Lookup",
        "tags": [
          "bin-lookup",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns card metadata for the given 6–8 digit BIN prefix.",
        "parameters": [
          {
            "name": "bin",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "424242",
            },
            "description":
              "6–8 digit Bank Identification Number. Dashes and spaces are stripped automatically.",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "bin": {
                          "type": "string",
                          "description":
                            "The normalised BIN prefix used for the lookup",
                        },
                        "scheme": {
                          "type": "string",
                          "description":
                            "Card network: visa, mastercard, amex, discover, jcb, diners, unionpay, maestro, mir, rupay, private_label",
                        },
                        "card_type": {
                          "type": "string",
                          "description": "credit, debit, prepaid, or charge",
                        },
                        "card_level": {
                          "type": "string",
                          "description":
                            "classic, gold, platinum, infinite, business, signature, or standard",
                        },
                        "issuer_name": {
                          "type": "string",
                          "description": "Name of the card-issuing bank",
                        },
                        "issuer_url": {
                          "type": "string",
                          "description": "Bank website URL",
                        },
                        "issuer_phone": {
                          "type": "string",
                          "description": "Bank customer service phone number",
                        },
                        "country_code": {
                          "type": "string",
                          "description":
                            "ISO 3166-1 alpha-2 country code of the issuing bank (e.g. US, GB, DE)",
                        },
                        "country_name": {
                          "type": "string",
                          "description":
                            "Full country name of the issuing bank",
                        },
                        "prepaid": {
                          "type": "boolean",
                          "description": "Whether the card is a prepaid card",
                        },
                        "luhn": {
                          "type": "boolean",
                          "description":
                            "Whether the BIN prefix passes the Luhn algorithm check",
                        },
                        "confidence": {
                          "type": "number",
                          "description":
                            "Data quality score (0.00–1.00). Multi-source confirmed records score higher.",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "bin": "424242",
                    "scheme": "visa",
                    "card_type": "credit",
                    "card_level": "classic",
                    "issuer_name": "Chase",
                    "issuer_url": "www.chase.com",
                    "issuer_phone": "+18002324000",
                    "country_code": "US",
                    "country_name": "United States",
                    "prepaid": false,
                    "luhn": true,
                    "confidence": 0.92,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "BIN is not 6–8 digits or contains non-digit characters.",
          },
          "404": {
            "description": "BIN prefix not found in the database.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/entertainment/chuck-norris": {
      "get": {
        "summary": "Get Random Chuck Norris Fact",
        "tags": [
          "chuck-norris",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a randomly selected Chuck Norris fact from the built-in database.",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "id": {
                          "type": "string",
                          "description":
                            "Unique fact identifier in the format cn_<index> (e.g. cn_0, cn_7)",
                        },
                        "fact": {
                          "type": "string",
                          "description": "The Chuck Norris fact text",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "id": "cn_0",
                    "fact": "Chuck Norris can divide by zero.",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
        },
      },
    },
    "/v1/places/cities/{city}": {
      "get": {
        "summary": "Get City Info",
        "tags": [
          "cities",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns metadata for a city by name. Lookup is case-insensitive.",
        "parameters": [
          {
            "name": "city",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "london",
            },
            "description":
              "City name to look up (e.g. london, tokyo, new york city)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "name": {
                          "type": "string",
                          "description":
                            "Official city name as listed in the GeoNames dataset",
                        },
                        "country": {
                          "type": "string",
                          "description":
                            "ISO 3166-1 alpha-2 country code (uppercase)",
                        },
                        "population": {
                          "type": "integer",
                          "description":
                            "City population from the GeoNames dataset",
                        },
                        "timezone": {
                          "type": "string",
                          "description":
                            'IANA timezone identifier for the city (e.g. "America/New_York")',
                        },
                        "lat": {
                          "type": "number",
                          "description": "Latitude of the city centre",
                        },
                        "lon": {
                          "type": "number",
                          "description": "Longitude of the city centre",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "name": "London",
                    "country": "GB",
                    "population": 7556900,
                    "timezone": "Europe/London",
                    "lat": 51.5085,
                    "lon": -0.1257,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "404": {
            "description": "No city with that name was found in the dataset.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/technology/color": {
      "get": {
        "summary": "Convert Color",
        "tags": [
          "color-conversion",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Convert a color value from one format to another. The response always includes all four formats.",
        "parameters": [
          {
            "name": "from",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "hex",
            },
            "description": "Source color format: hex, rgb, hsl, or cmyk",
          },
          {
            "name": "to",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "hsl",
            },
            "description": "Target color format: hex, rgb, hsl, or cmyk",
          },
          {
            "name": "value",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "#ff5733",
            },
            "description":
              "Color value in the source format (e.g. #ff5733, rgb(255,87,51), hsl(11,100%,60%), cmyk(0%,66%,80%,0%))",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "input": {
                          "type": "string",
                          "description":
                            "The original value passed in the value parameter",
                        },
                        "result": {
                          "type": "string",
                          "description":
                            "The color expressed in the requested to format",
                        },
                        "formats.hex": {
                          "type": "string",
                          "description": "HEX representation (#rrggbb)",
                        },
                        "formats.rgb": {
                          "type": "string",
                          "description": "RGB representation (rgb(r, g, b))",
                        },
                        "formats.hsl": {
                          "type": "string",
                          "description": "HSL representation (hsl(h, s%, l%))",
                        },
                        "formats.cmyk": {
                          "type": "string",
                          "description":
                            "CMYK representation (cmyk(c%, m%, y%, k%))",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "input": "#ff5733",
                    "result": "hsl(11, 100%, 60%)",
                    "formats": {
                      "hex": "#ff5733",
                      "rgb": "rgb(255, 87, 51)",
                      "hsl": "hsl(11, 100%, 60%)",
                      "cmyk": "cmyk(0%, 66%, 80%, 0%)",
                    },
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "One or more of from, to, or value parameters is missing or the from/to value is not one of: hex, rgb, hsl, cmyk.",
          },
          "422": {
            "description":
              "The value cannot be parsed in the specified from format.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/finance/commodities/{commodity}": {
      "get": {
        "summary": "Get Commodity Price",
        "tags": [
          "commodities",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns the latest annual average price and up to 10 years of historical data for the requested commodity slug.",
        "parameters": [
          {
            "name": "commodity",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "gold",
            },
            "description":
              "Commodity slug (e.g. gold, silver, oil). See supported slugs below.",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "commodity": {
                          "type": "string",
                          "description":
                            "The commodity slug as provided in the request path",
                        },
                        "name": {
                          "type": "string",
                          "description": "Human-readable commodity name",
                        },
                        "price": {
                          "type": "number",
                          "description":
                            "Latest annual average price in the commodity's display unit",
                        },
                        "unit": {
                          "type": "string",
                          "description":
                            "Price unit (oz, barrel, mmbtu, lb, or metric_ton)",
                        },
                        "currency": {
                          "type": "string",
                          "description": "Currency code — always USD",
                        },
                        "change_24h": {
                          "type": "number",
                          "description":
                            "Year-over-year percentage change from the prior year's annual average (positive = price increased)",
                        },
                        "historical": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Up to 10 prior years of annual average prices, ordered newest to oldest",
                        },
                        "historical[].period": {
                          "type": "string",
                          "description": "Year of the historical data point",
                        },
                        "historical[].price": {
                          "type": "number",
                          "description": "Annual average price for that year",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "commodity": "gold",
                    "name": "Gold",
                    "price": 2386.33,
                    "unit": "oz",
                    "currency": "USD",
                    "change_24h": 23.01,
                    "historical": [
                      {
                        "period": "2023",
                        "price": 1940.54,
                      },
                      {
                        "period": "2022",
                        "price": 1800.12,
                      },
                      {
                        "period": "2021",
                        "price": 1798.52,
                      },
                      {
                        "period": "2020",
                        "price": 1769.64,
                      },
                      {
                        "period": "2019",
                        "price": 1392.6,
                      },
                      {
                        "period": "2018",
                        "price": 1268.93,
                      },
                      {
                        "period": "2017",
                        "price": 1257.15,
                      },
                      {
                        "period": "2016",
                        "price": 1251.65,
                      },
                      {
                        "period": "2015",
                        "price": 1160.06,
                      },
                      {
                        "period": "2014",
                        "price": 1266.4,
                      },
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "404": {
            "description":
              "No data found for the given commodity slug. Check the supported slugs list.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/technology/counter/{namespace}": {
      "post": {
        "summary": "Increment Counter",
        "tags": [
          "counter",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Atomically increment a counter in the specified namespace and return the new value",
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "page-views",
            },
            "description":
              "Counter namespace (1-64 chars: alphanumeric, hyphen, underscore)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "namespace": {
                          "type": "string",
                          "description": "The counter namespace",
                        },
                        "value": {
                          "type": "integer",
                          "description":
                            "The new counter value after increment",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "namespace": "page-views",
                    "value": 42,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "Invalid namespace: must be 1–64 chars, alphanumeric, hyphen or underscore only",
          },
          "500": {
            "description": "Internal server error",
          },
        },
      },
      "get": {
        "summary": "Get Counter Value",
        "tags": [
          "counter",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Get the current value of a counter without incrementing it",
        "parameters": [
          {
            "name": "namespace",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "page-views",
            },
            "description":
              "Counter namespace (1-64 chars: alphanumeric, hyphen, underscore)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "namespace": {
                          "type": "string",
                          "description": "The counter namespace",
                        },
                        "value": {
                          "type": "integer",
                          "description":
                            "The current counter value (returns 0 if counter doesn't exist)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "namespace": "page-views",
                    "value": 42,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "Invalid namespace: must be 1–64 chars, alphanumeric, hyphen or underscore only",
          },
          "500": {
            "description": "Internal server error",
          },
        },
      },
    },
    "/v1/finance/crypto/{symbol}": {
      "get": {
        "summary": "Get Crypto Price",
        "tags": [
          "crypto",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns current price data for the given cryptocurrency symbol.",
        "parameters": [
          {
            "name": "symbol",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "BTC",
            },
            "description": "Uppercase ticker symbol (e.g. BTC, ETH, SOL)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "symbol": {
                          "type": "string",
                          "description": "Uppercase ticker symbol",
                        },
                        "name": {
                          "type": "string",
                          "description": "Full coin name",
                        },
                        "price_usd": {
                          "type": "number",
                          "description": "Current price in USD",
                        },
                        "change_24h": {
                          "type": "number",
                          "description":
                            "Price change over the last 24 hours as a percentage",
                        },
                        "market_cap": {
                          "type": "number",
                          "description": "Total market capitalisation in USD",
                        },
                        "volume_24h": {
                          "type": "number",
                          "description":
                            "Total trading volume over the last 24 hours in USD",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "symbol": "BTC",
                    "name": "Bitcoin",
                    "price_usd": 42000.5,
                    "change_24h": 2.5,
                    "market_cap": 820000000000,
                    "volume_24h": 25000000000,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "422": {
            "description": "The symbol is not in the supported coin list.",
          },
          "503": {
            "description":
              "CoinGecko is unavailable or returned an unexpected response.",
          },
        },
      },
    },
    "/v1/entertainment/jokes/dad": {
      "get": {
        "summary": "Get Random Dad Joke",
        "tags": [
          "dad-jokes",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a randomly selected dad joke from the collection.",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "id": {
                          "type": "string",
                          "description":
                            'Stable identifier for the joke (e.g. "joke_7")',
                        },
                        "joke": {
                          "type": "string",
                          "description": "The full text of the dad joke",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "id": "joke_7",
                    "joke":
                      "Why don't scientists trust atoms? Because they make up everything!",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "401": {
            "description": "Missing API key",
          },
          "403": {
            "description": "Invalid or revoked API key",
          },
        },
      },
    },
    "/v1/technology/format": {
      "post": {
        "summary": "Convert Format",
        "tags": [
          "data-format-conversion",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Convert content from one structured data format to another. Supported formats: json, yaml, csv, xml, toml.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "from": {
                    "type": "string",
                    "description":
                      "Source format. One of: json, yaml, csv, xml, toml",
                    "example": "json",
                  },
                  "to": {
                    "type": "string",
                    "description":
                      "Target format. One of: json, yaml, csv, xml, toml",
                    "example": "yaml",
                  },
                  "content": {
                    "type": "string",
                    "description":
                      "The content to convert, serialized as a string in the source format.",
                    "example": '{"name":"Alice","age":30}',
                  },
                },
                "required": [
                  "from",
                  "to",
                  "content",
                ],
                "example": {
                  "from": "json",
                  "to": "yaml",
                  "content": '{"name": "Alice", "age": 30}',
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "result": {
                          "type": "string",
                          "description":
                            "The converted content serialized in the target format.",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "result": "age: 30\nname: Alice\n",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "413": {
            "description": "The content field exceeds the 512 KB limit.",
          },
          "422": {
            "description":
              "One of from, to, or content is missing, or from/to is not one of the supported format values.; The content field is not valid JSON (when from is json).; The content field is not valid YAML (when from is yaml).; The content field is not valid CSV, or a row has more columns than the header (when from is csv).; The content field is not valid XML (when from is xml).; The content field is not valid TOML (when from is toml).; The data structure is incompatible with the target format (e.g. converting a JSON array to TOML, which requires a top-level object).",
          },
        },
      },
    },
    "/v1/text/detect-language": {
      "post": {
        "summary": "Detect Language",
        "tags": [
          "detect-language",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Identifies the language of the provided text and returns the language name, ISO 639-1 code, and confidence score.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "text": {
                    "type": "string",
                    "description":
                      "The text whose language should be detected.",
                    "example": "Bonjour, comment ça va?",
                  },
                },
                "required": [
                  "text",
                ],
                "example": {
                  "text": "Bonjour, comment ça va?",
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "language": {
                          "type": "string",
                          "description":
                            "Full name of the detected language (e.g. French, English, Spanish)",
                        },
                        "code": {
                          "type": "string",
                          "description":
                            "ISO 639-1 two-letter language code (e.g. fr, en, es). Empty string when detection is unreliable.",
                        },
                        "confidence": {
                          "type": "string",
                          "description":
                            "Confidence score between 0.0 and 1.0. 0.0 is returned when the language cannot be reliably detected.",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "language": "French",
                    "code": "fr",
                    "confidence": 0.98,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The request body is missing or malformed.",
          },
          "422": {
            "description": "The text field is missing or empty.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/text/dictionary/{word}": {
      "get": {
        "summary": "Dictionary Lookup",
        "tags": [
          "dictionary",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns the definition, phonetics, examples, and synonyms for the given word.",
        "parameters": [
          {
            "name": "word",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "ephemeral",
            },
            "description": "The word to look up in the dictionary",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "word": {
                          "type": "string",
                          "description":
                            "The normalized (lowercased) word that was looked up",
                        },
                        "phonetic": {
                          "type": "string",
                          "description":
                            "IPA phonetic transcription of the word (may be omitted if unavailable)",
                        },
                        "definitions": {
                          "type": "string",
                          "description":
                            "One or more definitions for the word, each with partOfSpeech, definition, and an optional example",
                        },
                        "definitions[].partOfSpeech": {
                          "type": "string",
                          "description":
                            "Grammatical category (e.g. noun, verb, adjective)",
                        },
                        "definitions[].definition": {
                          "type": "string",
                          "description": "Plain-text definition of the word",
                        },
                        "definitions[].example": {
                          "type": "string",
                          "description":
                            "Example sentence using the word (may be omitted)",
                        },
                        "synonyms": {
                          "type": "string",
                          "description": "List of words with similar meaning",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "word": "ephemeral",
                    "phonetic": "/ɪˈfɛm(ə)rəl/",
                    "definitions": [
                      {
                        "partOfSpeech": "adjective",
                        "definition": "lasting for a very short time",
                        "example": "ephemeral pleasures",
                      },
                    ],
                    "synonyms": [
                      "transient",
                      "fleeting",
                      "momentary",
                      "brief",
                      "short-lived",
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The word path parameter is missing.",
          },
          "404": {
            "description": "The word was not found in the dictionary dataset.",
          },
        },
      },
    },
    "/v1/networking/disposable/check": {
      "post": {
        "summary": "Check Single Email",
        "tags": [
          "disposable_email",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Validate whether an email address uses a disposable domain",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "email": {
                    "type": "string",
                    "description": "The email address to check",
                    "example": "test@example.com",
                  },
                },
                "required": [
                  "email",
                ],
                "example": {
                  "email": "user@tempmail.com",
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "email": {
                          "type": "string",
                          "description": "The email address that was checked",
                        },
                        "is_disposable": {
                          "type": "boolean",
                          "description":
                            "Whether the email uses a disposable domain",
                        },
                        "domain": {
                          "type": "string",
                          "description": "The domain part of the email address",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "email": "user@tempmail.com",
                    "is_disposable": true,
                    "domain": "tempmail.com",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "The request body is missing or malformed; The email address format is invalid",
          },
        },
      },
    },
    "/v1/networking/disposable/batch": {
      "post": {
        "summary": "Check Batch Emails",
        "tags": [
          "disposable_email",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Validate multiple email addresses in a single request (max 100 emails)",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "emails": {
                    "type": "array",
                    "items": {},
                    "description":
                      "Array of email addresses to check (max 100)",
                    "example": [
                      "user1@example.com",
                      "user2@tempmail.com",
                    ],
                  },
                },
                "required": [
                  "emails",
                ],
                "example": {
                  "emails": [
                    "user1@gmail.com",
                    "user2@tempmail.com",
                    "user3@guerrillamail.com",
                  ],
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "results": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Array of check results for each email",
                        },
                        "total": {
                          "type": "integer",
                          "description": "Total number of emails checked",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "results": [
                      {
                        "email": "user1@gmail.com",
                        "is_disposable": false,
                        "domain": "gmail.com",
                      },
                      {
                        "email": "user2@tempmail.com",
                        "is_disposable": true,
                        "domain": "tempmail.com",
                      },
                      {
                        "email": "user3@guerrillamail.com",
                        "is_disposable": true,
                        "domain": "guerrillamail.com",
                      },
                    ],
                    "total": 3,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "The request body is missing or malformed; The emails field is missing; Too many emails in the request",
          },
        },
      },
    },
    "/v1/networking/disposable/domain/{domain}": {
      "get": {
        "summary": "Check Domain",
        "tags": [
          "disposable_email",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Check if a specific domain is in the disposable blocklist",
        "parameters": [
          {
            "name": "domain",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "tempmail.com",
            },
            "description": "The domain to check",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "domain": {
                          "type": "string",
                          "description": "The domain that was checked",
                        },
                        "is_disposable": {
                          "type": "boolean",
                          "description":
                            "Whether the domain is in the disposable blocklist",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "domain": "tempmail.com",
                    "is_disposable": true,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The domain parameter is missing",
          },
        },
      },
    },
    "/v1/networking/disposable/stats": {
      "get": {
        "summary": "Get Statistics",
        "tags": [
          "disposable_email",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description": "Get statistics about the disposable email blocklist",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "total_domains": {
                          "type": "integer",
                          "description":
                            "Total number of disposable domains in the blocklist",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "total_domains": 10500,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
        },
      },
    },
    "/v1/networking/disposable/domains": {
      "get": {
        "summary": "List Domains (Paginated)",
        "tags": [
          "disposable_email",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Get a paginated list of all disposable domains in the blocklist",
        "parameters": [
          {
            "name": "page",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 1,
            },
            "description": "Page number (default: 1)",
          },
          {
            "name": "per_page",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 100,
            },
            "description": "Items per page (default: 100, max: 1000)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "domains": {
                          "type": "array",
                          "items": {},
                          "description": "Array of domain names",
                        },
                        "total": {
                          "type": "integer",
                          "description": "Total number of domains in blocklist",
                        },
                        "page": {
                          "type": "integer",
                          "description": "Current page number",
                        },
                        "per_page": {
                          "type": "integer",
                          "description": "Number of items per page",
                        },
                        "has_more": {
                          "type": "boolean",
                          "description":
                            "Whether there are more pages available",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "domains": [
                      "tempmail.com",
                      "guerrillamail.com",
                      "10minutemail.com",
                    ],
                    "total": 10500,
                    "page": 1,
                    "per_page": 100,
                    "has_more": true,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
        },
      },
    },
    "/v1/networking/domain/{domain}": {
      "get": {
        "summary": "Get Domain Info",
        "tags": [
          "domain-info",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns DNS records and availability status for the given domain.",
        "parameters": [
          {
            "name": "domain",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "example.com",
            },
            "description": "The domain to look up (e.g. example.com)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "domain": {
                          "type": "string",
                          "description": "The domain that was looked up",
                        },
                        "available": {
                          "type": "boolean",
                          "description":
                            "True when the domain appears to be unregistered (NS lookup returns NXDOMAIN). False when name servers are delegated.",
                        },
                        "dns.a": {
                          "type": "string",
                          "description": "IPv4 addresses (A records)",
                        },
                        "dns.aaaa": {
                          "type": "string",
                          "description": "IPv6 addresses (AAAA records)",
                        },
                        "dns.mx": {
                          "type": "string",
                          "description":
                            "Mail exchange records, each with host and priority fields",
                        },
                        "dns.ns": {
                          "type": "string",
                          "description": "Authoritative name server hostnames",
                        },
                        "dns.txt": {
                          "type": "string",
                          "description":
                            "TXT record values (SPF, DKIM, verification tokens, etc.)",
                        },
                        "dns.cname": {
                          "type": "string",
                          "description":
                            "CNAME alias target, if the domain is an alias. Empty string when no alias exists.",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "domain": "example.com",
                    "available": false,
                    "dns": {
                      "a": [
                        "93.184.216.34",
                      ],
                      "aaaa": [
                        "2606:2800:220:1:248:1893:25c8:1946",
                      ],
                      "mx": [],
                      "ns": [
                        "a.iana-servers.net.",
                        "b.iana-servers.net.",
                      ],
                      "txt": [
                        "v=spf1 -all",
                      ],
                      "cname": "",
                    },
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "The domain parameter is not a valid hostname (e.g. missing TLD, invalid characters, or leading/trailing hyphens).",
          },
        },
      },
    },
    "/v1/text/normalize": {
      "post": {
        "summary": "Normalize Email",
        "tags": [
          "email-normalize",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Normalizes a single email address and returns the canonical form together with a breakdown of all transformations applied.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "email": {
                    "type": "string",
                    "description":
                      "The email address to normalize. Must be a syntactically valid address.",
                    "example": "Te.st.User+spam@Googlemail.com",
                  },
                },
                "required": [
                  "email",
                ],
                "example": {
                  "email": "Te.st.User+spam@Googlemail.com",
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "original": {
                          "type": "string",
                          "description":
                            "The email address exactly as supplied in the request body",
                        },
                        "normalized": {
                          "type": "string",
                          "description":
                            "The canonical form of the address after all transformations",
                        },
                        "local": {
                          "type": "string",
                          "description":
                            "The local part (before @) of the normalized address",
                        },
                        "domain": {
                          "type": "string",
                          "description":
                            "The domain part (after @) of the normalized address",
                        },
                        "changes": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Ordered list of transformations applied. Possible values: lowercased, trimmed_whitespace, removed_dots, removed_plus_tag, canonicalised_domain. Empty array when no changes were needed.",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "original": "Te.st.User+spam@Googlemail.com",
                    "normalized": "testuser@gmail.com",
                    "local": "testuser",
                    "domain": "gmail.com",
                    "changes": [
                      "lowercased",
                      "removed_dots",
                      "removed_plus_tag",
                      "canonicalised_domain",
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "The request body is missing, not valid JSON, or contains unknown fields.",
          },
          "422": {
            "description":
              "The email field is missing or not a valid email address format.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/text/normalize/batch": {
      "post": {
        "summary": "Normalize Email Batch",
        "tags": [
          "email-normalize",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Normalizes up to 100 email addresses in one request. Results are in the same order as the input. Each item includes valid (boolean); when false, only original and message are set. Usage is billed per email processed (see gateway usage headers).",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "emails": {
                    "type": "array",
                    "items": {},
                    "description":
                      "Array of addresses to normalize (min 1, max 100; each entry non-empty)",
                    "example": [
                      "user@example.com",
                      "not-an-email",
                      "te.st@gmail.com",
                    ],
                  },
                },
                "required": [
                  "emails",
                ],
                "example": {
                  "emails": [
                    "user@example.com",
                    "not-an-email",
                    "te.st@gmail.com",
                  ],
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "results": {
                          "type": "array",
                          "items": {},
                          "description":
                            "One normalization result per input email, in order",
                        },
                        "total": {
                          "type": "integer",
                          "description":
                            "Number of emails in the batch (same as results length)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "results": [
                      {
                        "original": "user@example.com",
                        "normalized": "user@example.com",
                        "local": "user",
                        "domain": "example.com",
                        "changes": [],
                        "valid": true,
                      },
                      {
                        "original": "not-an-email",
                        "valid": false,
                        "message": "invalid email address",
                      },
                      {
                        "original": "te.st@gmail.com",
                        "normalized": "test@gmail.com",
                        "local": "test",
                        "domain": "gmail.com",
                        "changes": [
                          "lowercased",
                          "removed_dots",
                        ],
                        "valid": true,
                      },
                    ],
                    "total": 3,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "Invalid JSON or unknown fields in the body",
          },
          "422": {
            "description":
              "Missing emails, empty array, too many items, or empty string in the array",
          },
        },
      },
    },
    "/v1/validation/email": {
      "post": {
        "summary": "Validate Email",
        "tags": [
          "email-validate",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Validates a single email address and returns a full breakdown of syntax validity, MX record status, disposable domain check, normalized form, and any typo suggestion.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "email": {
                    "type": "string",
                    "description": "The email address to validate.",
                    "example": "user@gmial.com",
                  },
                },
                "required": [
                  "email",
                ],
                "example": {
                  "email": "user@gmial.com",
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "email": {
                          "type": "string",
                          "description":
                            "The email address exactly as supplied in the request body; null when syntax is invalid",
                        },
                        "valid": {
                          "type": "boolean",
                          "description":
                            "Overall validity. True only when the address passes syntax validation and the domain has at least one MX record",
                        },
                        "syntax_valid": {
                          "type": "boolean",
                          "description":
                            "Whether the address is syntactically valid according to RFC 5322",
                        },
                        "mx_valid": {
                          "type": "boolean",
                          "description":
                            "Whether the domain has at least one MX record, meaning it can receive email",
                        },
                        "disposable": {
                          "type": "boolean",
                          "description":
                            "Whether the address uses a known disposable or temporary email domain",
                        },
                        "normalized": {
                          "type": "string",
                          "description":
                            "The canonical form of the address after normalization (lowercase, plus-tag removal, alias-domain resolution). Null when syntax is invalid",
                        },
                        "domain": {
                          "type": "string",
                          "description":
                            "The domain part of the address (after @). Null when syntax is invalid",
                        },
                        "suggestion": {
                          "type": "string",
                          "description":
                            "A corrected domain name when the supplied domain looks like a typo of a well-known provider (e.g. gmial.com → gmail.com). Null when no close match is found or the domain is already correct",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "email": "user@gmial.com",
                    "valid": false,
                    "syntax_valid": true,
                    "mx_valid": false,
                    "disposable": false,
                    "normalized": "user@gmial.com",
                    "domain": "gmial.com",
                    "suggestion": "gmail.com",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "The request body is missing, not valid JSON, or contains unknown fields.",
          },
          "422": {
            "description": "The email field is missing from the request body.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/validation/email/batch": {
      "post": {
        "summary": "Validate Emails (Batch)",
        "tags": [
          "email-validate",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Validates up to 50 email addresses in a single request. Each email is processed independently and returns a full validation breakdown (syntax, MX record, disposable check, normalization, and typo suggestion). Invalid emails do not fail the request. Billing: 1 credit per email.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "emails": {
                    "type": "array",
                    "items": {},
                    "description":
                      "Array of email addresses to validate. Min: 1, Max: 50.",
                    "example": '["user@gmail.com", "user@gmial.com"]',
                  },
                },
                "required": [
                  "emails",
                ],
                "example": {
                  "emails": [
                    "user@gmail.com",
                    "user@gmial.com",
                  ],
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "results": {
                          "type": "array",
                          "items": {},
                          "description":
                            "List of validation results for each email, preserving input order",
                        },
                        "results[].email": {
                          "type": "string",
                          "description":
                            "Original email input (null if invalid syntax)",
                        },
                        "results[].valid": {
                          "type": "boolean",
                          "description":
                            "Overall validity (syntax + MX record)",
                        },
                        "results[].syntax_valid": {
                          "type": "boolean",
                          "description":
                            "Whether the email is syntactically valid (RFC 5322)",
                        },
                        "results[].mx_valid": {
                          "type": "boolean",
                          "description":
                            "Whether the domain has valid MX records",
                        },
                        "results[].disposable": {
                          "type": "boolean",
                          "description":
                            "Whether the email comes from a disposable domain",
                        },
                        "results[].normalized": {
                          "type": "string",
                          "description":
                            "Canonical normalized email (lowercase, alias handling, etc.)",
                        },
                        "results[].domain": {
                          "type": "string",
                          "description": "Extracted domain from email address",
                        },
                        "results[].suggestion": {
                          "type": "string",
                          "description":
                            "Suggested correction for common domain typos",
                        },
                        "total": {
                          "type": "integer",
                          "description":
                            "Number of emails processed in the batch",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "results": [
                      {
                        "email": "user@gmail.com",
                        "valid": true,
                        "syntax_valid": true,
                        "mx_valid": true,
                        "disposable": false,
                        "normalized": "user@gmail.com",
                        "domain": "gmail.com",
                        "suggestion": null,
                      },
                      {
                        "email": "user@gmial.com",
                        "valid": false,
                        "syntax_valid": true,
                        "mx_valid": false,
                        "disposable": false,
                        "normalized": "user@gmial.com",
                        "domain": "gmial.com",
                        "suggestion": "gmail.com",
                      },
                    ],
                    "total": 2,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "Invalid JSON, malformed request body, or unexpected field types.",
          },
          "422": {
            "description":
              "Valid JSON body that fails field validation (empty array or more than 50 emails).",
          },
          "500": {
            "description": "Unexpected server error",
          },
        },
      },
    },
    "/v1/entertainment/emoji/random": {
      "get": {
        "summary": "Get Random Emoji",
        "tags": [
          "emoji",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a randomly selected emoji with its full metadata.",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "emoji": {
                          "type": "string",
                          "description": "The rendered emoji glyph",
                        },
                        "name": {
                          "type": "string",
                          "description":
                            "CLDR short name in snake_case (e.g. grinning_face)",
                        },
                        "category": {
                          "type": "string",
                          "description":
                            "Unicode category (e.g. Smileys & Emotion, Animals & Nature)",
                        },
                        "unicode": {
                          "type": "string",
                          "description":
                            "Unicode code-point in U+XXXX notation (e.g. U+1F600)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "emoji": "😀",
                    "name": "grinning_face",
                    "category": "Smileys & Emotion",
                    "unicode": "U+1F600",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
        },
      },
    },
    "/v1/entertainment/emoji/search": {
      "get": {
        "summary": "Search Emoji",
        "tags": [
          "emoji",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Search for emojis whose name or category contains the given query string (case-insensitive). Returns a list of all matches.",
        "parameters": [
          {
            "name": "q",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "happy",
            },
            "description":
              "Search term to match against emoji names and categories (e.g. happy, heart, food)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "items": {
                          "type": "array",
                          "items": {},
                          "description": "List of matching emoji objects",
                        },
                        "total": {
                          "type": "integer",
                          "description": "Total number of matches",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "items": [
                      {
                        "emoji": "😄",
                        "name": "grinning_face_with_smiling_eyes",
                        "category": "Smileys & Emotion",
                        "unicode": "U+1F604",
                      },
                    ],
                    "total": 1,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The q query parameter is missing or empty.",
          },
        },
      },
    },
    "/v1/entertainment/emoji/{name}": {
      "get": {
        "summary": "Get Emoji by Name",
        "tags": [
          "emoji",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a specific emoji by its CLDR snake_case name. The name is case-insensitive.",
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "grinning_face",
            },
            "description":
              "CLDR snake_case emoji name (e.g. grinning_face, thumbs_up)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "emoji": {
                          "type": "string",
                          "description": "The rendered emoji glyph",
                        },
                        "name": {
                          "type": "string",
                          "description": "CLDR short name in snake_case",
                        },
                        "category": {
                          "type": "string",
                          "description": "Unicode category",
                        },
                        "unicode": {
                          "type": "string",
                          "description":
                            "Unicode code-point in U+XXXX notation",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "emoji": "😀",
                    "name": "grinning_face",
                    "category": "Smileys & Emotion",
                    "unicode": "U+1F600",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "404": {
            "description": "No emoji found with the given name.",
          },
        },
      },
    },
    "/v1/finance/exchange-rate": {
      "get": {
        "summary": "Get Exchange Rate",
        "tags": [
          "exchange-rate",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns the current exchange rate between two currencies.",
        "parameters": [
          {
            "name": "from",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "USD",
            },
            "description":
              "ISO 4217 source currency code (3 letters, e.g. USD)",
          },
          {
            "name": "to",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "EUR",
            },
            "description":
              "ISO 4217 target currency code (3 letters, e.g. EUR)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "from": {
                          "type": "string",
                          "description": "Source currency code (uppercased)",
                        },
                        "to": {
                          "type": "string",
                          "description": "Target currency code (uppercased)",
                        },
                        "rate": {
                          "type": "number",
                          "description":
                            "Exchange rate — how many units of `to` equal 1 unit of `from`",
                        },
                        "timestamp": {
                          "type": "string",
                          "description":
                            "Date the rate was published by the ECB (ISO 8601)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "from": "USD",
                    "to": "EUR",
                    "rate": 0.92,
                    "timestamp": "2024-12-15T00:00:00Z",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "A required parameter is missing or the currency code is not exactly 3 alphabetic characters.",
          },
          "422": {
            "description":
              "One or both currency codes are not recognised by the upstream data source.",
          },
          "503": {
            "description":
              "The exchange rate data source is temporarily unavailable.",
          },
        },
      },
    },
    "/v1/finance/convert": {
      "get": {
        "summary": "Convert Currency",
        "tags": [
          "exchange-rate",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Converts an amount from one currency to another and returns the rate alongside the converted value.",
        "parameters": [
          {
            "name": "from",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "USD",
            },
            "description":
              "ISO 4217 source currency code (3 letters, e.g. USD)",
          },
          {
            "name": "to",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "EUR",
            },
            "description":
              "ISO 4217 target currency code (3 letters, e.g. EUR)",
          },
          {
            "name": "amount",
            "in": "query",
            "required": true,
            "schema": {
              "type": "number",
              "example": 100,
            },
            "description": "Amount to convert. Must be greater than 0.",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "from": {
                          "type": "string",
                          "description": "Source currency code (uppercased)",
                        },
                        "to": {
                          "type": "string",
                          "description": "Target currency code (uppercased)",
                        },
                        "rate": {
                          "type": "number",
                          "description":
                            "Exchange rate used for the conversion",
                        },
                        "amount": {
                          "type": "number",
                          "description":
                            "The original amount passed in the request",
                        },
                        "converted": {
                          "type": "number",
                          "description":
                            "Result of amount × rate, rounded to 2 decimal places",
                        },
                        "timestamp": {
                          "type": "string",
                          "description":
                            "Date the rate was published by the ECB (ISO 8601)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "from": "USD",
                    "to": "EUR",
                    "rate": 0.92,
                    "amount": 100,
                    "converted": 92,
                    "timestamp": "2024-12-15T00:00:00Z",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "A required parameter is missing, the currency code is not 3 alphabetic characters, or the amount is 0 or negative.",
          },
          "422": {
            "description":
              "One or both currency codes are not recognised by the upstream data source.",
          },
          "503": {
            "description":
              "The exchange rate data source is temporarily unavailable.",
          },
        },
      },
    },
    "/v1/entertainment/facts": {
      "get": {
        "summary": "Get Random Fact",
        "tags": [
          "facts",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a randomly selected fact, optionally filtered by category.",
        "parameters": [
          {
            "name": "category",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "science",
            },
            "description":
              "Filter by category. Valid values: science, history, technology, nature, space, food",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "fact": {
                          "type": "string",
                          "description": "The fact text",
                        },
                        "category": {
                          "type": "string",
                          "description": "The category the fact belongs to",
                        },
                        "source": {
                          "type": "string",
                          "description":
                            "The source or publication the fact is attributed to",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "fact":
                      "Honey never spoils. Archaeologists have found 3,000-year-old honey in Egyptian tombs that was still perfectly edible.",
                    "category": "science",
                    "source": "National Geographic",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "bad_request",
          },
        },
      },
    },
    "/v1/health/exercises": {
      "get": {
        "summary": "List Exercises",
        "tags": [
          "fitness-exercises",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a paginated list of exercises. All filter parameters are optional and combinable.",
        "parameters": [
          {
            "name": "body_part",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "chest",
            },
            "description":
              "Filter by body part (e.g. chest, back, upper legs). Use /v1/health/body-parts for valid values.",
          },
          {
            "name": "equipment",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "barbell",
            },
            "description":
              "Filter by equipment type (e.g. barbell, dumbbell, body weight). Use /v1/health/equipment for valid values.",
          },
          {
            "name": "muscle",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "biceps",
            },
            "description":
              "Filter by target or secondary muscle (e.g. biceps, glutes). Use /v1/health/muscles for valid values.",
          },
          {
            "name": "search",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "bench press",
            },
            "description": "Full-text search on exercise name.",
          },
          {
            "name": "page",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 1,
            },
            "description": "Page number (default: 1)",
          },
          {
            "name": "per_page",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 20,
            },
            "description": "Results per page, 1–100 (default: 20)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "items": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Array of exercise objects for the current page",
                        },
                        "items[].id": {
                          "type": "integer",
                          "description": "Unique exercise identifier",
                        },
                        "items[].name": {
                          "type": "string",
                          "description": "Exercise name",
                        },
                        "items[].body_parts": {
                          "type": "array",
                          "items": {},
                          "description": "Body part categories involved",
                        },
                        "items[].equipment": {
                          "type": "array",
                          "items": {},
                          "description": "Equipment required",
                        },
                        "items[].target_muscles": {
                          "type": "array",
                          "items": {},
                          "description": "Primary muscles targeted",
                        },
                        "items[].secondary_muscles": {
                          "type": "array",
                          "items": {},
                          "description": "Secondary muscles engaged",
                        },
                        "items[].instructions": {
                          "type": "array",
                          "items": {},
                          "description": "Ordered step-by-step instructions",
                        },
                        "total": {
                          "type": "integer",
                          "description":
                            "Total number of exercises matching the filters",
                        },
                        "page": {
                          "type": "integer",
                          "description": "Current page number",
                        },
                        "per_page": {
                          "type": "integer",
                          "description": "Number of results per page",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "items": [
                      {
                        "id": 1,
                        "name": "barbell bench press",
                        "body_parts": [
                          "chest",
                        ],
                        "equipment": [
                          "barbell",
                        ],
                        "target_muscles": [
                          "pectorals",
                        ],
                        "secondary_muscles": [
                          "triceps",
                          "deltoids",
                        ],
                        "instructions": [
                          "Lie flat on a bench and grip the barbell slightly wider than shoulder-width.",
                          "Lower the bar to your chest under control.",
                          "Press the bar back up to full arm extension.",
                          "Repeat for the desired number of repetitions.",
                        ],
                      },
                    ],
                    "total": 312,
                    "page": 1,
                    "per_page": 20,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "A query parameter has an invalid value (e.g. per_page out of range).",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/health/exercises/{id}": {
      "get": {
        "summary": "Get Exercise by ID",
        "tags": [
          "fitness-exercises",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description": "Returns a single exercise by its numeric ID.",
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "example": 1,
            },
            "description": "Numeric exercise ID",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "id": {
                          "type": "integer",
                          "description": "Unique exercise identifier",
                        },
                        "name": {
                          "type": "string",
                          "description": "Exercise name",
                        },
                        "body_parts": {
                          "type": "array",
                          "items": {},
                          "description": "Body part categories involved",
                        },
                        "equipment": {
                          "type": "array",
                          "items": {},
                          "description": "Equipment required",
                        },
                        "target_muscles": {
                          "type": "array",
                          "items": {},
                          "description": "Primary muscles targeted",
                        },
                        "secondary_muscles": {
                          "type": "array",
                          "items": {},
                          "description": "Secondary muscles engaged",
                        },
                        "instructions": {
                          "type": "array",
                          "items": {},
                          "description": "Ordered step-by-step instructions",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "id": 1,
                    "name": "barbell bench press",
                    "body_parts": [
                      "chest",
                    ],
                    "equipment": [
                      "barbell",
                    ],
                    "target_muscles": [
                      "pectorals",
                    ],
                    "secondary_muscles": [
                      "triceps",
                      "deltoids",
                    ],
                    "instructions": [
                      "Lie flat on a bench and grip the barbell slightly wider than shoulder-width.",
                      "Lower the bar to your chest under control.",
                      "Press the bar back up to full arm extension.",
                      "Repeat for the desired number of repetitions.",
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The id parameter is not a positive integer.",
          },
          "404": {
            "description": "No exercise exists with the given ID.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/health/exercises/random": {
      "get": {
        "summary": "Random Exercise",
        "tags": [
          "fitness-exercises",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a single randomly selected exercise. Accepts the same filter parameters as the list endpoint, so you can get a random chest exercise, a random bodyweight exercise, etc.",
        "parameters": [
          {
            "name": "body_part",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "back",
            },
            "description": "Restrict random selection to this body part.",
          },
          {
            "name": "equipment",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "body weight",
            },
            "description": "Restrict random selection to this equipment type.",
          },
          {
            "name": "muscle",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "glutes",
            },
            "description":
              "Restrict random selection to exercises targeting this muscle.",
          },
          {
            "name": "search",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "curl",
            },
            "description":
              "Restrict random selection to exercises matching this search term.",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {},
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "id": 42,
                    "name": "dumbbell curl",
                    "body_parts": [
                      "upper arms",
                    ],
                    "equipment": [
                      "dumbbell",
                    ],
                    "target_muscles": [
                      "biceps",
                    ],
                    "secondary_muscles": [
                      "brachialis",
                    ],
                    "instructions": [
                      "Stand with a dumbbell in each hand, arms fully extended.",
                      "Curl the weights toward your shoulders, keeping your elbows close to your torso.",
                      "Squeeze at the top, then lower the weights slowly.",
                      "Repeat for the desired number of repetitions.",
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "404": {
            "description": "No exercises match the given filters.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/health/body-parts": {
      "get": {
        "summary": "List Body Parts",
        "tags": [
          "fitness-exercises",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a sorted list of all distinct body part values present in the dataset. Use these as valid values for the body_part filter.",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "items": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Sorted list of all distinct body part names",
                        },
                        "total": {
                          "type": "integer",
                          "description": "Total number of distinct body parts",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "items": [
                      "back",
                      "cardio",
                      "chest",
                      "lower arms",
                      "lower legs",
                      "neck",
                      "shoulders",
                      "upper arms",
                      "upper legs",
                      "waist",
                    ],
                    "total": 10,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/health/equipment": {
      "get": {
        "summary": "List Equipment",
        "tags": [
          "fitness-exercises",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a sorted list of all distinct equipment types. Use these as valid values for the equipment filter.",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "items": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Sorted list of all distinct equipment names",
                        },
                        "total": {
                          "type": "integer",
                          "description":
                            "Total number of distinct equipment types",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "items": [
                      "band",
                      "barbell",
                      "body weight",
                      "cable",
                      "dumbbell",
                      "ez barbell",
                      "kettlebell",
                      "medicine ball",
                      "stability ball",
                      "tire",
                    ],
                    "total": 28,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/health/muscles": {
      "get": {
        "summary": "List Muscles",
        "tags": [
          "fitness-exercises",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a sorted list of all distinct muscle names (combining target and secondary muscles). Use these as valid values for the muscle filter.",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "items": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Sorted list of all distinct muscle names",
                        },
                        "total": {
                          "type": "integer",
                          "description": "Total number of distinct muscles",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "items": [
                      "biceps",
                      "calves",
                      "deltoids",
                      "glutes",
                      "hamstrings",
                      "lats",
                      "pectorals",
                      "quadriceps",
                      "traps",
                      "triceps",
                    ],
                    "total": 51,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/places/geocode": {
      "get": {
        "summary": "Geocode Address",
        "tags": [
          "geocode",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Converts a free-text address into latitude and longitude coordinates.",
        "parameters": [
          {
            "name": "address",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "1600 Pennsylvania Ave NW, Washington DC",
            },
            "description":
              "The address to geocode (street, city, country, or any combination)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "address": {
                          "type": "string",
                          "description":
                            "Full display name of the matched location",
                        },
                        "city": {
                          "type": "string",
                          "description": "City or town of the matched location",
                        },
                        "country": {
                          "type": "string",
                          "description":
                            "ISO 3166-1 alpha-2 country code (uppercase)",
                        },
                        "lat": {
                          "type": "number",
                          "description": "Latitude of the matched location",
                        },
                        "lon": {
                          "type": "number",
                          "description": "Longitude of the matched location",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "address":
                      "White House, 1600, Pennsylvania Avenue Northwest, Ward 2, Washington, District of Columbia, 20500, United States",
                    "city": "Washington",
                    "country": "US",
                    "lat": 38.8976763,
                    "lon": -77.0365298,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The address parameter is missing.",
          },
          "404": {
            "description": "No results found for the given address.",
          },
          "503": {
            "description": "The geocoding service is temporarily unavailable.",
          },
        },
      },
    },
    "/v1/places/reverse-geocode": {
      "get": {
        "summary": "Reverse Geocode",
        "tags": [
          "geocode",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Converts geographic coordinates into a human-readable address.",
        "parameters": [
          {
            "name": "lat",
            "in": "query",
            "required": true,
            "schema": {
              "type": "number",
              "example": 38.8977,
            },
            "description": "Latitude of the location (-90 to 90)",
          },
          {
            "name": "lon",
            "in": "query",
            "required": true,
            "schema": {
              "type": "number",
              "example": -77.0365,
            },
            "description": "Longitude of the location (-180 to 180)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "lat": {
                          "type": "number",
                          "description": "Latitude as provided in the request",
                        },
                        "lon": {
                          "type": "number",
                          "description": "Longitude as provided in the request",
                        },
                        "address": {
                          "type": "string",
                          "description":
                            "Full display name of the location at the given coordinates",
                        },
                        "city": {
                          "type": "string",
                          "description":
                            "City or town at the given coordinates",
                        },
                        "country": {
                          "type": "string",
                          "description":
                            "ISO 3166-1 alpha-2 country code (uppercase)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "lat": 38.8977,
                    "lon": -77.0365,
                    "address":
                      "White House, 1600, Pennsylvania Avenue Northwest, Ward 2, Washington, District of Columbia, 20500, United States",
                    "city": "Washington",
                    "country": "US",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "lat or lon is missing or out of range (lat: -90..90, lon: -180..180).",
          },
          "404": {
            "description": "No address found for the given coordinates.",
          },
          "503": {
            "description": "The geocoding service is temporarily unavailable.",
          },
        },
      },
    },
    "/v1/places/holidays": {
      "get": {
        "summary": "Get Holidays",
        "tags": [
          "holidays",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a list of public holidays for the specified country and year",
        "parameters": [
          {
            "name": "country",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "US",
            },
            "description":
              'ISO 3166-1 alpha-2 country code (e.g., "US", "GB", "DE")',
          },
          {
            "name": "year",
            "in": "query",
            "required": true,
            "schema": {
              "type": "integer",
              "example": 2025,
            },
            "description": "Year for which to retrieve holidays (e.g., 2025)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "country": {
                          "type": "string",
                          "description": "ISO 3166-1 alpha-2 country code",
                        },
                        "year": {
                          "type": "integer",
                          "description": "Year for which holidays are returned",
                        },
                        "holidays": {
                          "type": "array",
                          "items": {},
                          "description": "Array of holiday objects",
                        },
                        "holidays[].date": {
                          "type": "string",
                          "description": "Holiday date in YYYY-MM-DD format",
                        },
                        "holidays[].name": {
                          "type": "string",
                          "description": "Name of the holiday",
                        },
                        "total": {
                          "type": "integer",
                          "description":
                            "Total number of holidays for the country/year",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "country": "US",
                    "year": 2025,
                    "holidays": [
                      {
                        "date": "2025-01-01",
                        "name": "New Year's Day",
                      },
                      {
                        "date": "2025-07-04",
                        "name": "Independence Day",
                      },
                    ],
                    "total": 11,
                  },
                  "metadata": {
                    "timestamp": "2025-01-15T10:30:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "Missing or invalid country code or year parameter",
          },
          "404": {
            "description":
              "No holidays found for the specified country and year",
          },
        },
      },
    },
    "/v1/places/holidays/batch": {
      "post": {
        "summary": "Batch Get Holidays",
        "tags": [
          "holidays",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns holidays for up to 50 (country, year) pairs in a single request. Each pair is processed independently — if one combination has no data, it returns found:false without failing the entire batch.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "queries": {
                    "type": "array",
                    "items": {},
                    "description":
                      "Array of (country, year) pairs. Min: 1, Max: 50.",
                    "example":
                      '[{"country":"US","year":2025},{"country":"AR","year":2024}]',
                  },
                },
                "required": [
                  "queries",
                ],
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "results": {
                          "type": "array",
                          "items": {},
                          "description":
                            "One result per query, in the same order as the request",
                        },
                        "results[].country": {
                          "type": "string",
                          "description": "ISO 3166-1 alpha-2 country code",
                        },
                        "results[].year": {
                          "type": "integer",
                          "description": "Year queried",
                        },
                        "results[].found": {
                          "type": "boolean",
                          "description":
                            "false when no holidays exist for that country/year combination",
                        },
                        "results[].holidays": {
                          "type": "array",
                          "items": {},
                          "description":
                            "List of holidays. Omitted when found is false.",
                        },
                        "results[].total": {
                          "type": "integer",
                          "description":
                            "Number of holidays. Omitted when found is false.",
                        },
                        "total": {
                          "type": "integer",
                          "description":
                            "Total number of results (equals the number of queries sent)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "results": [
                      {
                        "country": "US",
                        "year": 2025,
                        "found": true,
                        "holidays": [
                          {
                            "date": "2025-01-01",
                            "name": "New Year's Day",
                          },
                          {
                            "date": "2025-07-04",
                            "name": "Independence Day",
                          },
                        ],
                        "total": 11,
                      },
                      {
                        "country": "AR",
                        "year": 2024,
                        "found": true,
                        "holidays": [
                          {
                            "date": "2024-01-01",
                            "name": "Año Nuevo",
                          },
                        ],
                        "total": 19,
                      },
                      {
                        "country": "AQ",
                        "year": 2025,
                        "found": false,
                      },
                    ],
                    "total": 3,
                  },
                  "metadata": {
                    "timestamp": "2025-01-15T10:30:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "Malformed request body",
          },
          "422": {
            "description":
              "queries is missing, empty, exceeds 50 items, or contains invalid country codes or years",
          },
        },
      },
    },
    "/v1/entertainment/horoscope/{sign}": {
      "get": {
        "summary": "Get Daily Horoscope",
        "tags": [
          "horoscope",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a daily horoscope reading for the specified zodiac sign.",
        "parameters": [
          {
            "name": "sign",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
            },
            "description":
              "Zodiac sign (case-insensitive). Supported values: aries, taurus, gemini, cancer, leo, virgo, libra, scorpio, sagittarius, capricorn, aquarius, pisces",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "sign": {
                          "type": "string",
                          "description": "Normalized zodiac sign (lowercase)",
                        },
                        "date": {
                          "type": "string",
                          "description":
                            "Today's date in YYYY-MM-DD format (UTC)",
                        },
                        "horoscope": {
                          "type": "string",
                          "description": "Daily horoscope reading",
                        },
                        "lucky_number": {
                          "type": "integer",
                          "description": "Lucky number for the day (1-99)",
                        },
                        "mood": {
                          "type": "string",
                          "description": "Suggested mood for the day",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "sign": "aries",
                    "date": "2024-12-15",
                    "horoscope":
                      "Today is a great day for new beginnings. Trust your instincts and take that first step toward your goals.",
                    "lucky_number": 7,
                    "mood": "energetic",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
        },
      },
    },
    "/v1/finance/iban/{iban}": {
      "get": {
        "summary": "Validate IBAN",
        "tags": [
          "iban",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Validates an IBAN and returns the country, bank code, and account number. Spaces in the input are stripped automatically. Always returns HTTP 200 — check the valid field to determine whether the IBAN is valid.",
        "parameters": [
          {
            "name": "iban",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "DE89370400440532013000",
            },
            "description":
              "The IBAN to validate. Spaces are stripped. Case-insensitive.",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "iban": {
                          "type": "string",
                          "description":
                            "The normalised IBAN (spaces stripped, uppercased)",
                        },
                        "valid": {
                          "type": "boolean",
                          "description":
                            "true if the IBAN passed length and ISO 13616 checksum validation",
                        },
                        "country": {
                          "type": "string",
                          "description":
                            "Full country name (empty if the country code is not in the registry)",
                        },
                        "bank_code": {
                          "type": "string",
                          "description":
                            "Bank identifier extracted from the BBAN (empty if country not in registry or positions not defined)",
                        },
                        "account": {
                          "type": "string",
                          "description":
                            "Account number extracted from the BBAN (empty if country not in registry or positions not defined)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "iban": "DE89370400440532013000",
                    "valid": true,
                    "country": "Germany",
                    "bank_code": "37040044",
                    "account": "0532013000",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "500": {
            "description":
              "Unexpected server error (e.g. database unreachable).",
          },
        },
      },
    },
    "/v1/finance/iban/batch": {
      "post": {
        "summary": "Batch Validate IBANs",
        "tags": [
          "iban",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Validates up to 50 iban numbers in a single request. Results are returned in the same order as the input.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "numbers": {
                    "type": "array",
                    "items": {},
                    "description":
                      "Array of iban numbers to validate (min: 1, max: 50).",
                    "example":
                      '["GB29NWBK60161331926819", "DE89370400440532013000"]',
                  },
                },
                "required": [
                  "numbers",
                ],
                "example": {
                  "numbers": [
                    "GB29NWBK60161331926819",
                    "DE89370400440532013000",
                    "XX89370400440532013000",
                  ],
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "results": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Validation result for each number in the same order as the input. Each item has the same fields as the single validate endpoint.",
                        },
                        "total": {
                          "type": "integer",
                          "description":
                            "Number of results returned. Matches the length of the input array.",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "results": [
                      {
                        "iban": "GB29NWBK60161331926819",
                        "valid": true,
                        "country": "United Kingdom",
                        "bank_code": "NWBK",
                        "account": "31926819",
                      },
                      {
                        "iban": "DE89370400440532013000",
                        "valid": true,
                        "country": "Germany",
                        "bank_code": "37040044",
                        "account": "0532013000",
                      },
                      {
                        "iban": "XX89370400440532013000",
                        "valid": false,
                        "country": "",
                        "bank_code": "",
                        "account": "",
                      },
                    ],
                    "total": 3,
                  },
                  "metadata": {
                    "timestamp": "2026-05-03T19:25:02Z",
                  },
                },
              },
            },
          },
          "422": {
            "description":
              "The numbers array is missing, empty, or contains more than 50 items.",
          },
        },
      },
    },
    "/v1/finance/inflation": {
      "get": {
        "summary": "Get Inflation Rate",
        "tags": [
          "inflation",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns the latest annual CPI inflation rate for a country plus the previous 10 years of historical data.",
        "parameters": [
          {
            "name": "country",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "US",
            },
            "description":
              "ISO 3166-1 alpha-2 country code (e.g. US, GB, DE). Case-insensitive.",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "country": {
                          "type": "string",
                          "description":
                            "ISO 3166-1 alpha-2 country code, uppercased",
                        },
                        "rate": {
                          "type": "number",
                          "description":
                            "Latest annual CPI inflation rate as a percentage (e.g. 2.9495 means 2.9495%)",
                        },
                        "period": {
                          "type": "string",
                          "description":
                            "Year of the latest data point (e.g. 2024)",
                        },
                        "historical": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Up to 10 previous years of inflation data, ordered newest to oldest",
                        },
                        "historical[].period": {
                          "type": "string",
                          "description": "Year of the historical data point",
                        },
                        "historical[].rate": {
                          "type": "number",
                          "description":
                            "Annual CPI inflation rate for that year",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "country": "US",
                    "rate": 2.9495,
                    "period": "2024",
                    "historical": [
                      {
                        "period": "2023",
                        "rate": 4.1163,
                      },
                      {
                        "period": "2022",
                        "rate": 8.0028,
                      },
                      {
                        "period": "2021",
                        "rate": 4.6979,
                      },
                      {
                        "period": "2020",
                        "rate": 1.2336,
                      },
                      {
                        "period": "2019",
                        "rate": 1.8122,
                      },
                      {
                        "period": "2018",
                        "rate": 2.4426,
                      },
                      {
                        "period": "2017",
                        "rate": 2.1301,
                      },
                      {
                        "period": "2016",
                        "rate": 1.2616,
                      },
                      {
                        "period": "2015",
                        "rate": 0.1186,
                      },
                      {
                        "period": "2014",
                        "rate": 1.6222,
                      },
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "The country parameter is missing or is not a valid ISO 3166-1 alpha-2 code.",
          },
          "404": {
            "description":
              "No inflation data found for the given country code.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/finance/inflation/batch": {
      "post": {
        "summary": "Batch Inflation Rates",
        "tags": [
          "inflation",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns inflation data for up to 50 countries in a single request. Results are in the same order as the input. Countries with no data return found: false instead of failing the whole request. Billing: 1 credit per country (not per HTTP request).",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "countries": {
                    "type": "array",
                    "items": {},
                    "description":
                      "Array of ISO 3166-1 alpha-2 country codes. Min: 1, Max: 50.",
                    "example": '["US", "AR", "DE"]',
                  },
                },
                "required": [
                  "countries",
                ],
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "results": {
                          "type": "array",
                          "items": {},
                          "description":
                            "One entry per country, in the same order as the input array",
                        },
                        "results[].country": {
                          "type": "string",
                          "description":
                            "ISO 3166-1 alpha-2 country code, uppercased",
                        },
                        "results[].found": {
                          "type": "boolean",
                          "description":
                            "false when the country has no data in the World Bank set",
                        },
                        "results[].rate": {
                          "type": "number",
                          "description":
                            "Latest CPI inflation rate. Omitted when found: false",
                        },
                        "results[].period": {
                          "type": "string",
                          "description":
                            "Year of the latest data point. Omitted when found: false",
                        },
                        "results[].historical": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Up to 10 previous years. Omitted when found: false",
                        },
                        "total": {
                          "type": "integer",
                          "description":
                            "Total number of results returned (equals number of countries sent)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "results": [
                      {
                        "country": "US",
                        "found": true,
                        "rate": 2.9495,
                        "period": "2024",
                        "historical": [
                          {
                            "period": "2023",
                            "rate": 4.1163,
                          },
                        ],
                      },
                      {
                        "country": "AR",
                        "found": true,
                        "rate": 211.4,
                        "period": "2024",
                        "historical": [],
                      },
                      {
                        "country": "XX",
                        "found": false,
                      },
                    ],
                    "total": 3,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "422": {
            "description":
              "Body is invalid: empty array, more than 50 items, or a bad country code.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/networking/ip/asn": {
      "get": {
        "summary": "Lookup ASN (Caller IP)",
        "tags": [
          "ip-asn",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Look up ASN, organization, ISP, and network details for the requesting client's IP address. Useful when you want information about the user making the request without specifying an IP explicitly.",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "ip": {
                          "type": "string",
                          "description":
                            "The IP address that was looked up (the requesting client's IP)",
                        },
                        "asn": {
                          "type": "string",
                          "description":
                            'Autonomous System Number in format "ASxxxx" (e.g., "AS15169")',
                        },
                        "org": {
                          "type": "string",
                          "description":
                            "Organization name owning the IP address range",
                        },
                        "isp": {
                          "type": "string",
                          "description":
                            "Internet Service Provider providing the IP",
                        },
                        "domain": {
                          "type": "string",
                          "description":
                            "Domain name associated with the IP or IP range",
                        },
                        "route": {
                          "type": "string",
                          "description":
                            'CIDR notation of the network route (e.g., "8.8.8.0/24")',
                        },
                        "type": {
                          "type": "string",
                          "description":
                            'Type of network (e.g., "hosting", "isp", "business", "cdn")',
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "ip": "8.8.8.8",
                    "asn": "AS15169",
                    "org": "Google LLC",
                    "isp": "Google Public DNS",
                    "domain": "google.com",
                    "route": "8.8.8.0/24",
                    "type": "hosting",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "500": {
            "description": "Unexpected server error",
          },
        },
      },
    },
    "/v1/networking/ip/asn/{ip}": {
      "get": {
        "summary": "Lookup ASN for IP",
        "tags": [
          "ip-asn",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Look up ASN, organization, ISP, and network details for a specific IP address.",
        "parameters": [
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "8.8.8.8",
            },
            "description": "The IP address to look up (supports IPv4 and IPv6)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "ip": {
                          "type": "string",
                          "description": "The IP address that was looked up",
                        },
                        "asn": {
                          "type": "string",
                          "description":
                            'Autonomous System Number in format "ASxxxx" (e.g., "AS15169")',
                        },
                        "org": {
                          "type": "string",
                          "description":
                            "Organization name owning the IP address range",
                        },
                        "isp": {
                          "type": "string",
                          "description":
                            "Internet Service Provider providing the IP",
                        },
                        "domain": {
                          "type": "string",
                          "description":
                            "Domain name associated with the IP or IP range",
                        },
                        "route": {
                          "type": "string",
                          "description":
                            'CIDR notation of the network route (e.g., "8.8.8.0/24")',
                        },
                        "type": {
                          "type": "string",
                          "description":
                            'Type of network (e.g., "hosting", "isp", "business", "cdn")',
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "ip": "8.8.8.8",
                    "asn": "AS15169",
                    "org": "Google LLC",
                    "isp": "Google Public DNS",
                    "domain": "google.com",
                    "route": "8.8.8.0/24",
                    "type": "hosting",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The IP address is invalid",
          },
          "500": {
            "description": "Unexpected server error",
          },
        },
      },
    },
    "/v1/networking/ip": {
      "get": {
        "summary": "Get IP Info (Caller IP)",
        "tags": [
          "ip-info",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Get geolocation and network information for the requesting client's IP address. Useful when you want information about the user making the request without specifying an IP explicitly.",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "ip": {
                          "type": "string",
                          "description":
                            "The IP address that was looked up (the requesting client's IP)",
                        },
                        "country": {
                          "type": "string",
                          "description": "Country name where the IP is located",
                        },
                        "country_code": {
                          "type": "string",
                          "description":
                            'Two-letter ISO country code (e.g., "US", "GB", "DE")',
                        },
                        "city": {
                          "type": "string",
                          "description": "City name where the IP is located",
                        },
                        "isp": {
                          "type": "string",
                          "description":
                            "Internet Service Provider providing the IP",
                        },
                        "is_vpn": {
                          "type": "boolean",
                          "description":
                            "True when the IP belongs to a known VPN",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "ip": "8.8.8.8",
                    "country": "United States",
                    "country_code": "US",
                    "city": "Mountain View",
                    "isp": "Google Public DNS",
                    "is_vpn": false,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "500": {
            "description": "Unexpected server error",
          },
        },
      },
    },
    "/v1/networking/ip/{ip}": {
      "get": {
        "summary": "Get IP Info for IP",
        "tags": [
          "ip-info",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Get geolocation and network information for a specific IP address.",
        "parameters": [
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "8.8.8.8",
            },
            "description": "The IP address to look up (supports IPv4 and IPv6)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "ip": {
                          "type": "string",
                          "description": "The IP address that was looked up",
                        },
                        "country": {
                          "type": "string",
                          "description": "Country name where the IP is located",
                        },
                        "country_code": {
                          "type": "string",
                          "description":
                            'Two-letter ISO country code (e.g., "US", "GB", "DE")',
                        },
                        "city": {
                          "type": "string",
                          "description": "City name where the IP is located",
                        },
                        "isp": {
                          "type": "string",
                          "description":
                            "Internet Service Provider providing the IP",
                        },
                        "is_vpn": {
                          "type": "boolean",
                          "description":
                            "True when the IP belongs to a known VPN",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "ip": "8.8.8.8",
                    "country": "United States",
                    "country_code": "US",
                    "city": "Mountain View",
                    "isp": "Google Public DNS",
                    "is_vpn": false,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The IP address is invalid",
          },
          "500": {
            "description": "Unexpected server error",
          },
        },
      },
    },
    "/v1/text/lorem": {
      "get": {
        "summary": "Generate Lorem Ipsum",
        "tags": [
          "lorem-ipsum",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Generate Lorem Ipsum placeholder text with customizable length and format",
        "parameters": [
          {
            "name": "paragraphs",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 3,
            },
            "description": "Number of paragraphs to generate (1-20)",
          },
          {
            "name": "sentences",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 5,
            },
            "description": "Number of sentences per paragraph (1-20)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "text": {
                          "type": "string",
                          "description": "Generated Lorem Ipsum text",
                        },
                        "paragraphs": {
                          "type": "integer",
                          "description": "Number of paragraphs generated",
                        },
                        "wordCount": {
                          "type": "integer",
                          "description":
                            "Total number of words in generated text",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "text":
                      "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.",
                    "paragraphs": 1,
                    "wordCount": 45,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "The paragraphs parameter is out of valid range; The sentences parameter is out of valid range",
          },
        },
      },
    },
    "/v1/technology/markdown": {
      "post": {
        "summary": "Convert Markdown to HTML",
        "tags": [
          "markdown",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Converts a Markdown string to HTML. Pass sanitize true to strip potentially unsafe tags like script and iframe from the output.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "markdown": {
                    "type": "string",
                    "description": "The Markdown text to convert.",
                    "example": "# Hello\n\nThis is **bold** and _italic_ text.",
                  },
                  "sanitize": {
                    "type": "boolean",
                    "description":
                      "When true, sanitizes the HTML output to remove unsafe tags and attributes.",
                  },
                },
                "required": [
                  "markdown",
                ],
                "example": {
                  "markdown": "# Hello\n\nThis is **bold** and _italic_ text.",
                  "sanitize": true,
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "html": {
                          "type": "string",
                          "description": "The rendered HTML output",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "html":
                      "<h1>Hello</h1>\n<p>This is <strong>bold</strong> and <em>italic</em> text.</p>\n",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "422": {
            "description": "unprocessable_entity",
          },
        },
      },
    },
    "/v1/finance/mortgage": {
      "get": {
        "summary": "Calculate Mortgage",
        "tags": [
          "mortgage",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns the monthly payment, total cost, and full amortization schedule for a fixed-rate mortgage.",
        "parameters": [
          {
            "name": "principal",
            "in": "query",
            "required": true,
            "schema": {
              "type": "number",
              "example": 300000,
            },
            "description":
              "Loan amount in your chosen currency (e.g. 300000 for $300,000)",
          },
          {
            "name": "rate",
            "in": "query",
            "required": true,
            "schema": {
              "type": "number",
              "example": 6.5,
            },
            "description":
              "Annual interest rate as a percentage (e.g. 6.5 for 6.5%). Must be greater than 0.",
          },
          {
            "name": "years",
            "in": "query",
            "required": true,
            "schema": {
              "type": "integer",
              "example": 30,
            },
            "description": "Loan term in years (1–50)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "principal": {
                          "type": "number",
                          "description":
                            "The original loan amount passed in the request",
                        },
                        "rate": {
                          "type": "number",
                          "description":
                            "The annual interest rate passed in the request",
                        },
                        "years": {
                          "type": "integer",
                          "description":
                            "The loan term in years passed in the request",
                        },
                        "monthly_payment": {
                          "type": "number",
                          "description":
                            "Fixed monthly payment amount (rounded to 2 decimal places)",
                        },
                        "total_payment": {
                          "type": "number",
                          "description":
                            "Total amount paid over the life of the loan",
                        },
                        "total_interest": {
                          "type": "number",
                          "description":
                            "Total interest paid (total_payment minus principal)",
                        },
                        "schedule": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Full amortization schedule — one entry per month (years × 12 entries)",
                        },
                        "schedule[].month": {
                          "type": "integer",
                          "description": "Month number (1 to years × 12)",
                        },
                        "schedule[].payment": {
                          "type": "number",
                          "description": "Total payment for this month",
                        },
                        "schedule[].principal": {
                          "type": "number",
                          "description":
                            "Portion of this month's payment applied to principal",
                        },
                        "schedule[].interest": {
                          "type": "number",
                          "description":
                            "Portion of this month's payment applied to interest",
                        },
                        "schedule[].balance": {
                          "type": "number",
                          "description":
                            "Remaining loan balance after this payment",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "principal": 300000,
                    "rate": 6.5,
                    "years": 30,
                    "monthly_payment": 1896.2,
                    "total_payment": 682632,
                    "total_interest": 382632,
                    "schedule": [
                      {
                        "month": 1,
                        "payment": 1896.2,
                        "principal": 271.2,
                        "interest": 1625,
                        "balance": 299728.8,
                      },
                      {
                        "month": 2,
                        "payment": 1896.2,
                        "principal": 272.67,
                        "interest": 1623.53,
                        "balance": 299456.13,
                      },
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "A required parameter is missing, not a valid number, or out of range (e.g. years > 50 or rate <= 0).",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/networking/mx/{domain}": {
      "get": {
        "summary": "MX Lookup",
        "tags": [
          "mx-lookup",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Retrieve all MX records for a domain. Results are sorted by priority ascending (lowest numeric value has highest mail delivery priority per RFC 5321).",
        "parameters": [
          {
            "name": "domain",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "gmail.com",
            },
            "description":
              "The domain name to look up MX records for (e.g. gmail.com)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "domain": {
                          "type": "string",
                          "description": "The domain that was queried",
                        },
                        "records": {
                          "type": "array",
                          "items": {},
                          "description":
                            "List of MX records, sorted by priority ascending (lowest number = highest priority)",
                        },
                        "records[].host": {
                          "type": "string",
                          "description":
                            "Fully-qualified hostname of the mail server (typically ends with a trailing dot)",
                        },
                        "records[].priority": {
                          "type": "integer",
                          "description":
                            "MX priority value. Lower values have higher delivery priority per RFC 5321.",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "domain": "gmail.com",
                    "records": [
                      {
                        "host": "gmail-smtp-in.l.google.com.",
                        "priority": 5,
                      },
                      {
                        "host": "alt1.gmail-smtp-in.l.google.com.",
                        "priority": 10,
                      },
                      {
                        "host": "alt2.gmail-smtp-in.l.google.com.",
                        "priority": 20,
                      },
                      {
                        "host": "alt3.gmail-smtp-in.l.google.com.",
                        "priority": 30,
                      },
                      {
                        "host": "alt4.gmail-smtp-in.l.google.com.",
                        "priority": 40,
                      },
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The domain parameter is not a valid domain name.",
          },
          "404": {
            "description":
              "No MX records were found for the domain (domain may not accept email).",
          },
          "500": {
            "description":
              "DNS lookup failed due to an unexpected server error.",
          },
        },
      },
    },
    "/v1/technology/base": {
      "get": {
        "summary": "Convert Base",
        "tags": [
          "number-base-conversion",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description": "Convert an integer from one number base to another.",
        "parameters": [
          {
            "name": "from",
            "in": "query",
            "required": true,
            "schema": {
              "type": "integer",
              "example": 10,
            },
            "description": "Source base (2, 8, 10, or 16)",
          },
          {
            "name": "to",
            "in": "query",
            "required": true,
            "schema": {
              "type": "integer",
              "example": 16,
            },
            "description": "Target base (2, 8, 10, or 16)",
          },
          {
            "name": "value",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "255",
            },
            "description":
              "The number as a string. Accepts optional prefixes: 0x (hex), 0b (binary), 0o (octal).",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "input": {
                          "type": "string",
                          "description":
                            "The original value as provided in the request",
                        },
                        "from": {
                          "type": "integer",
                          "description": "The source base",
                        },
                        "to": {
                          "type": "integer",
                          "description": "The target base",
                        },
                        "result": {
                          "type": "string",
                          "description":
                            "The converted value in the target base",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "input": "255",
                    "from": 10,
                    "to": 16,
                    "result": "ff",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "A required parameter is missing, the base is not one of 2/8/10/16, or value is not valid for the given base.",
          },
        },
      },
    },
    "/v1/technology/password": {
      "get": {
        "summary": "Generate Password",
        "tags": [
          "password-generator",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Generate a cryptographically secure random password with customizable character sets and length",
        "parameters": [
          {
            "name": "length",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 16,
            },
            "description": "Password length (8-128 characters)",
          },
          {
            "name": "uppercase",
            "in": "query",
            "required": false,
            "schema": {
              "type": "boolean",
              "example": true,
            },
            "description": "Include uppercase letters (A-Z)",
          },
          {
            "name": "numbers",
            "in": "query",
            "required": false,
            "schema": {
              "type": "boolean",
              "example": true,
            },
            "description": "Include numbers (0-9)",
          },
          {
            "name": "symbols",
            "in": "query",
            "required": false,
            "schema": {
              "type": "boolean",
              "example": true,
            },
            "description":
              "Include special characters (!@#$%^&*()-_=+[]{}|;:,.<>?)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "password": {
                          "type": "string",
                          "description": "The generated password",
                        },
                        "length": {
                          "type": "integer",
                          "description": "Length of the generated password",
                        },
                        "strength": {
                          "type": "string",
                          "description":
                            "Password strength assessment (weak, medium, or strong)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "password": "aB3#cDeFgHiJkLmN",
                    "length": 16,
                    "strength": "strong",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The length parameter is out of valid range (8-128)",
          },
          "500": {
            "description":
              "Failed to generate password (rare cryptographic failure)",
          },
        },
      },
    },
    "/v1/validation/phone": {
      "get": {
        "summary": "Validate Phone Number",
        "tags": [
          "phone-validation",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Validates a single phone number and returns its country, type, formatted representation, carrier, and VOIP/virtual risk flags.",
        "parameters": [
          {
            "name": "number",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "+447400123456",
            },
            "description":
              "The phone number to validate. Must include the country calling code (e.g. +12015551234).",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "number": {
                          "type": "string",
                          "description":
                            "The original number as supplied in the request",
                        },
                        "valid": {
                          "type": "boolean",
                          "description":
                            "Whether the number is a valid, dialable phone number",
                        },
                        "country": {
                          "type": "string",
                          "description":
                            "ISO 3166-1 alpha-2 country code (omitted when valid is false)",
                        },
                        "type": {
                          "type": "string",
                          "description":
                            "Number type: mobile, landline, landline_or_mobile, toll_free, voip, premium_rate, shared_cost, personal_number, pager, uan, voicemail, or unknown (omitted when valid is false)",
                        },
                        "formatted": {
                          "type": "string",
                          "description":
                            "International format of the number, e.g. +44 7400 123456 (omitted when valid is false)",
                        },
                        "carrier.name": {
                          "type": "string",
                          "description":
                            "Carrier name from phone prefix metadata (omitted when carrier cannot be determined)",
                        },
                        "carrier.source": {
                          "type": "string",
                          "description":
                            'How the carrier was determined. Always "metadata" when present',
                        },
                        "risk.is_voip": {
                          "type": "boolean",
                          "description": "true when the number type is voip",
                        },
                        "risk.is_virtual": {
                          "type": "boolean",
                          "description":
                            "true when the number is not tied to a physical SIM or fixed line: voip, personal_number, uan, pager, or voicemail",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "number": "+447400123456",
                    "valid": true,
                    "country": "GB",
                    "type": "mobile",
                    "formatted": "+44 7400 123456",
                    "carrier": {
                      "name": "Three",
                      "source": "metadata",
                    },
                    "risk": {
                      "is_voip": false,
                      "is_virtual": false,
                    },
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The number query parameter is missing.",
          },
        },
      },
    },
    "/v1/validation/phone/batch": {
      "post": {
        "summary": "Batch Validate Phone Numbers",
        "tags": [
          "phone-validation",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Validates up to 50 phone numbers in a single request. Results are returned in the same order as the input.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "numbers": {
                    "type": "array",
                    "items": {},
                    "description":
                      "Array of phone numbers to validate (min: 1, max: 50). Each must include the country calling code.",
                    "example": '["+447400123456", "+12015551234"]',
                  },
                },
                "required": [
                  "numbers",
                ],
                "example": {
                  "numbers": [
                    "+447400123456",
                    "+12015551234",
                    "12345",
                  ],
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "results": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Validation result for each number in the same order as the input. Each item has the same fields as the single validate endpoint.",
                        },
                        "total": {
                          "type": "integer",
                          "description":
                            "Number of results returned. Matches the length of the input array.",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "results": [
                      {
                        "number": "+447400123456",
                        "valid": true,
                        "country": "GB",
                        "type": "mobile",
                        "formatted": "+44 7400 123456",
                        "carrier": {
                          "name": "Three",
                          "source": "metadata",
                        },
                        "risk": {
                          "is_voip": false,
                          "is_virtual": false,
                        },
                      },
                      {
                        "number": "+12015551234",
                        "valid": true,
                        "country": "US",
                        "type": "landline_or_mobile",
                        "formatted": "+1 201-555-1234",
                        "risk": {
                          "is_voip": false,
                          "is_virtual": false,
                        },
                      },
                      {
                        "number": "12345",
                        "valid": false,
                      },
                    ],
                    "total": 3,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "422": {
            "description":
              "The numbers array is missing, empty, or contains more than 50 items.",
          },
        },
      },
    },
    "/v1/places/postal/{code}": {
      "get": {
        "summary": "Lookup Postal Code",
        "tags": [
          "postal-code",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns city, state, country, and coordinates for the given postal code.",
        "parameters": [
          {
            "name": "code",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "10001",
            },
            "description":
              "The postal code to look up (e.g. 10001 for New York, SW1A 1AA for London)",
          },
          {
            "name": "country",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "US",
            },
            "description": "ISO 3166-1 alpha-2 country code (default: US)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "postal_code": {
                          "type": "string",
                          "description":
                            "The postal code as stored in the dataset",
                        },
                        "city": {
                          "type": "string",
                          "description":
                            "Primary city or place name for the postal code",
                        },
                        "state": {
                          "type": "string",
                          "description":
                            "State, province, or administrative region name",
                        },
                        "country": {
                          "type": "string",
                          "description":
                            "ISO 3166-1 alpha-2 country code (uppercase)",
                        },
                        "lat": {
                          "type": "number",
                          "description": "Latitude of the postal code centroid",
                        },
                        "lon": {
                          "type": "number",
                          "description":
                            "Longitude of the postal code centroid",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "postal_code": "10001",
                    "city": "New York City",
                    "state": "New York",
                    "country": "US",
                    "lat": 40.7484,
                    "lon": -73.9967,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "404": {
            "description":
              "The postal code was not found for the given country.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/text/profanity": {
      "post": {
        "summary": "Check Profanity",
        "tags": [
          "profanity",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Checks text for profanity, returning a censored version and the list of flagged words.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "text": {
                    "type": "string",
                    "description": "The text to check for profanity.",
                    "example": "What the heck is going on",
                  },
                },
                "required": [
                  "text",
                ],
                "example": {
                  "text": "Some text to check",
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "has_profanity": {
                          "type": "boolean",
                          "description":
                            "Whether any profanity was detected in the text",
                        },
                        "censored": {
                          "type": "string",
                          "description":
                            "The input text with profane words replaced by asterisks",
                        },
                        "flagged_words": {
                          "type": "string",
                          "description":
                            "Deduplicated list of profane words found (lowercase)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "has_profanity": false,
                    "censored": "Some text to check",
                    "flagged_words": [],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The request body is missing or malformed.",
          },
          "422": {
            "description": "The text field is missing or empty.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/technology/qr": {
      "get": {
        "summary": "Generate QR Code (PNG)",
        "tags": [
          "qr-code",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a raw PNG image of the QR code. Ideal for direct embedding or file download.",
        "parameters": [
          {
            "name": "data",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "https://example.com",
            },
            "description": "The text or URL to encode in the QR code",
          },
          {
            "name": "size",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 200,
            },
            "description":
              "Image size in pixels (default: 256, min: 50, max: 1000)",
          },
          {
            "name": "recovery",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "high",
            },
            "description":
              "Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Higher levels are more robust to physical damage but produce larger images. Default: medium",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
          },
          "400": {
            "description":
              "Missing or invalid parameters (e.g. data not provided, size out of range, unknown recovery level)",
          },
          "500": {
            "description": "Failed to generate QR code",
          },
        },
      },
    },
    "/v1/technology/qr/base64": {
      "get": {
        "summary": "Generate QR Code (Base64 JSON)",
        "tags": [
          "qr-code",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a JSON envelope containing the QR code as a base64-encoded PNG string, along with its dimensions.",
        "parameters": [
          {
            "name": "data",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "https://example.com",
            },
            "description": "The text or URL to encode in the QR code",
          },
          {
            "name": "size",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 200,
            },
            "description":
              "Image size in pixels (default: 256, min: 50, max: 1000)",
          },
          {
            "name": "recovery",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "highest",
            },
            "description":
              "Error-correction level: low (7%), medium (15%), high (25%), highest (30%). Default: medium",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "image": {
                          "type": "string",
                          "description": "Base64-encoded PNG image data",
                        },
                        "width": {
                          "type": "integer",
                          "description":
                            "Width of the generated image in pixels",
                        },
                        "height": {
                          "type": "integer",
                          "description":
                            "Height of the generated image in pixels",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "image": "<base64-encoded PNG data>",
                    "width": 256,
                    "height": 256,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "Missing or invalid parameters (e.g. data not provided, size out of range, unknown recovery level)",
          },
          "500": {
            "description": "Failed to generate QR code",
          },
        },
      },
    },
    "/v1/text/quotes/random": {
      "get": {
        "summary": "Get Random Quote",
        "tags": [
          "quotes",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a random inspirational quote with author attribution",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "id": {
                          "type": "integer",
                          "description": "Unique identifier for the quote",
                        },
                        "text": {
                          "type": "string",
                          "description": "The quote text",
                        },
                        "author": {
                          "type": "string",
                          "description":
                            "Name of the person who said or wrote the quote",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "id": 42,
                    "text":
                      "The only way to do great work is to love what you do.",
                    "author": "Steve Jobs",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "503": {
            "description": "No quotes available in the database",
          },
        },
      },
    },
    "/v1/technology/random-user": {
      "get": {
        "summary": "Get Random User",
        "tags": [
          "random-user",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description": "Returns a randomly generated fake user profile.",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "name": {
                          "type": "string",
                          "description": "Full name of the generated user",
                        },
                        "email": {
                          "type": "string",
                          "description": "Email address of the generated user",
                        },
                        "phone": {
                          "type": "string",
                          "description": "Phone number of the generated user",
                        },
                        "address.street": {
                          "type": "string",
                          "description": "Street address",
                        },
                        "address.city": {
                          "type": "string",
                          "description": "City name",
                        },
                        "address.state": {
                          "type": "string",
                          "description": "State or region",
                        },
                        "address.zip": {
                          "type": "string",
                          "description": "Postal / ZIP code",
                        },
                        "address.country": {
                          "type": "string",
                          "description": "Country name",
                        },
                        "avatar": {
                          "type": "string",
                          "description":
                            "URL to a unique identicon avatar for the generated user (DiceBear)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "name": "Grace Lopez",
                    "email": "grace.lopez@example.org",
                    "phone": "555-123-4567",
                    "address": {
                      "street": "4821 Maple Avenue",
                      "city": "North Judyton",
                      "state": "California",
                      "zip": "94103",
                      "country": "United States of America",
                    },
                    "avatar":
                      "https://api.dicebear.com/9.x/identicon/svg?seed=Grace+Lopez",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "500": {
            "description": "Internal server error",
          },
        },
      },
    },
    "/v1/text/sentiment": {
      "post": {
        "summary": "Analyze Sentiment",
        "tags": [
          "sentiment",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Analyzes the sentiment of the provided text and returns a classification, confidence score, and a full breakdown across all three sentiment classes.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "text": {
                    "type": "string",
                    "description": "The text to analyze.",
                    "example":
                      "I absolutely love this product, it exceeded my expectations!",
                  },
                },
                "required": [
                  "text",
                ],
                "example": {
                  "text":
                    "I absolutely love this product, it exceeded my expectations!",
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "sentiment": {
                          "type": "string",
                          "description":
                            "The dominant sentiment class: positive, negative, or neutral",
                        },
                        "score": {
                          "type": "number",
                          "description":
                            "Confidence score for the dominant sentiment, between 0.0 and 1.0",
                        },
                        "breakdown.positive": {
                          "type": "number",
                          "description":
                            "Proportional score for positive sentiment (sums to 1.0 with other classes)",
                        },
                        "breakdown.negative": {
                          "type": "number",
                          "description":
                            "Proportional score for negative sentiment (sums to 1.0 with other classes)",
                        },
                        "breakdown.neutral": {
                          "type": "number",
                          "description":
                            "Proportional score for neutral sentiment (sums to 1.0 with other classes)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "sentiment": "positive",
                    "score": 0.97,
                    "breakdown": {
                      "positive": 0.97,
                      "negative": 0.01,
                      "neutral": 0.02,
                    },
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "422": {
            "description": "unprocessable_entity",
          },
        },
      },
    },
    "/v1/text/spellcheck": {
      "post": {
        "summary": "Check Spelling",
        "tags": [
          "spell-check",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Checks the input text for spelling mistakes and returns a corrected version along with per-word corrections.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "text": {
                    "type": "string",
                    "description": "The text to spell-check.",
                    "example": "Ths is a smiple tset",
                  },
                },
                "required": [
                  "text",
                ],
                "example": {
                  "text": "Ths is a smiple tset",
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "corrected": {
                          "type": "string",
                          "description":
                            "The full input text with all misspelled words replaced by their suggested corrections",
                        },
                        "corrections": {
                          "type": "string",
                          "description":
                            "List of individual corrections. Each item contains: original (the misspelled word), suggested (the correction), and position (0-based character offset in the original text)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "corrected": "This is a simple test",
                    "corrections": [
                      {
                        "original": "Ths",
                        "suggested": "This",
                        "position": 0,
                      },
                      {
                        "original": "smiple",
                        "suggested": "simple",
                        "position": 9,
                      },
                      {
                        "original": "tset",
                        "suggested": "test",
                        "position": 16,
                      },
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The request body is missing or malformed.",
          },
          "422": {
            "description": "The text field is missing or empty.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/entertainment/sudoku/batch": {
      "post": {
        "summary": "Batch Generate Sudoku Puzzles",
        "tags": [
          "sudoku",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Generate up to 20 Sudoku puzzles in a single request. Results are returned in the same order as the input array. Each puzzle in the batch counts as one unit of API usage.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "puzzles": {
                    "type": "array",
                    "items": {},
                    "description":
                      "Array of difficulty levels to generate (min: 1, max: 20). Each must be one of: easy, medium, hard.",
                    "example": '["easy", "medium", "hard"]',
                  },
                },
                "required": [
                  "puzzles",
                ],
                "example": {
                  "puzzles": [
                    "easy",
                    "hard",
                  ],
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "results": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Generated puzzles in the same order as the input array. Each item has the same fields as the single-puzzle endpoint.",
                        },
                        "results[].difficulty": {
                          "type": "string",
                          "description":
                            "The difficulty level of the puzzle (easy, medium, or hard)",
                        },
                        "results[].puzzle": {
                          "type": "string",
                          "description":
                            "9×9 grid representing the puzzle — 0 means an empty cell to be filled in",
                        },
                        "results[].solution": {
                          "type": "string",
                          "description":
                            "9×9 grid containing the complete, valid solution",
                        },
                        "total": {
                          "type": "integer",
                          "description":
                            "Number of puzzles returned. Matches the length of the input array.",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "results": [
                      {
                        "difficulty": "easy",
                        "puzzle": [
                          [
                            5,
                            3,
                            0,
                            0,
                            7,
                            0,
                            0,
                            0,
                            0,
                          ],
                          [
                            6,
                            0,
                            0,
                            1,
                            9,
                            5,
                            0,
                            0,
                            0,
                          ],
                          [
                            0,
                            9,
                            8,
                            0,
                            0,
                            0,
                            0,
                            6,
                            0,
                          ],
                          [
                            8,
                            0,
                            0,
                            0,
                            6,
                            0,
                            0,
                            0,
                            3,
                          ],
                          [
                            4,
                            0,
                            0,
                            8,
                            0,
                            3,
                            0,
                            0,
                            1,
                          ],
                          [
                            7,
                            0,
                            0,
                            0,
                            2,
                            0,
                            0,
                            0,
                            6,
                          ],
                          [
                            0,
                            6,
                            0,
                            0,
                            0,
                            0,
                            2,
                            8,
                            0,
                          ],
                          [
                            0,
                            0,
                            0,
                            4,
                            1,
                            9,
                            0,
                            0,
                            5,
                          ],
                          [
                            0,
                            0,
                            0,
                            0,
                            8,
                            0,
                            0,
                            7,
                            9,
                          ],
                        ],
                        "solution": [
                          [
                            5,
                            3,
                            4,
                            6,
                            7,
                            8,
                            9,
                            1,
                            2,
                          ],
                          [
                            6,
                            7,
                            2,
                            1,
                            9,
                            5,
                            3,
                            4,
                            8,
                          ],
                          [
                            1,
                            9,
                            8,
                            3,
                            4,
                            2,
                            5,
                            6,
                            7,
                          ],
                          [
                            8,
                            5,
                            9,
                            7,
                            6,
                            1,
                            4,
                            2,
                            3,
                          ],
                          [
                            4,
                            2,
                            6,
                            8,
                            5,
                            3,
                            7,
                            9,
                            1,
                          ],
                          [
                            7,
                            1,
                            3,
                            9,
                            2,
                            4,
                            8,
                            5,
                            6,
                          ],
                          [
                            9,
                            6,
                            1,
                            5,
                            3,
                            7,
                            2,
                            8,
                            4,
                          ],
                          [
                            2,
                            8,
                            7,
                            4,
                            1,
                            9,
                            6,
                            3,
                            5,
                          ],
                          [
                            3,
                            4,
                            5,
                            2,
                            8,
                            6,
                            1,
                            7,
                            9,
                          ],
                        ],
                      },
                    ],
                    "total": 1,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "The request body is missing or contains malformed JSON.",
          },
          "401": {
            "description": "Missing API key",
          },
          "403": {
            "description": "Invalid or revoked API key",
          },
          "422": {
            "description":
              "The puzzles array is missing, empty, exceeds 20 items, or contains a value other than easy, medium, or hard.",
          },
        },
      },
    },
    "/v1/entertainment/sudoku": {
      "get": {
        "summary": "Get Sudoku Puzzle",
        "tags": [
          "sudoku",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a randomly generated Sudoku puzzle and its solution. Difficulty defaults to medium when not specified.",
        "parameters": [
          {
            "name": "difficulty",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
            },
            "description":
              "Puzzle difficulty level. One of: easy, medium, hard. Defaults to medium.",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "difficulty": {
                          "type": "string",
                          "description":
                            "The difficulty level of the returned puzzle (easy, medium, or hard)",
                        },
                        "puzzle": {
                          "type": "string",
                          "description":
                            "9×9 grid representing the puzzle — 0 means an empty cell to be filled in",
                        },
                        "solution": {
                          "type": "string",
                          "description":
                            "9×9 grid containing the complete, valid solution",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "difficulty": "hard",
                    "puzzle": [
                      [
                        5,
                        3,
                        0,
                        0,
                        7,
                        0,
                        0,
                        0,
                        0,
                      ],
                      [
                        6,
                        0,
                        0,
                        1,
                        9,
                        5,
                        0,
                        0,
                        0,
                      ],
                      [
                        0,
                        9,
                        8,
                        0,
                        0,
                        0,
                        0,
                        6,
                        0,
                      ],
                      [
                        8,
                        0,
                        0,
                        0,
                        6,
                        0,
                        0,
                        0,
                        3,
                      ],
                      [
                        4,
                        0,
                        0,
                        8,
                        0,
                        3,
                        0,
                        0,
                        1,
                      ],
                      [
                        7,
                        0,
                        0,
                        0,
                        2,
                        0,
                        0,
                        0,
                        6,
                      ],
                      [
                        0,
                        6,
                        0,
                        0,
                        0,
                        0,
                        2,
                        8,
                        0,
                      ],
                      [
                        0,
                        0,
                        0,
                        4,
                        1,
                        9,
                        0,
                        0,
                        5,
                      ],
                      [
                        0,
                        0,
                        0,
                        0,
                        8,
                        0,
                        0,
                        7,
                        9,
                      ],
                    ],
                    "solution": [
                      [
                        5,
                        3,
                        4,
                        6,
                        7,
                        8,
                        9,
                        1,
                        2,
                      ],
                      [
                        6,
                        7,
                        2,
                        1,
                        9,
                        5,
                        3,
                        4,
                        8,
                      ],
                      [
                        1,
                        9,
                        8,
                        3,
                        4,
                        2,
                        5,
                        6,
                        7,
                      ],
                      [
                        8,
                        5,
                        9,
                        7,
                        6,
                        1,
                        4,
                        2,
                        3,
                      ],
                      [
                        4,
                        2,
                        6,
                        8,
                        5,
                        3,
                        7,
                        9,
                        1,
                      ],
                      [
                        7,
                        1,
                        3,
                        9,
                        2,
                        4,
                        8,
                        5,
                        6,
                      ],
                      [
                        9,
                        6,
                        1,
                        5,
                        3,
                        7,
                        2,
                        8,
                        4,
                      ],
                      [
                        2,
                        8,
                        7,
                        4,
                        1,
                        9,
                        6,
                        3,
                        5,
                      ],
                      [
                        3,
                        4,
                        5,
                        2,
                        8,
                        6,
                        1,
                        7,
                        9,
                      ],
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "The difficulty parameter is not one of easy, medium, or hard",
          },
          "401": {
            "description": "Missing API key",
          },
          "403": {
            "description": "Invalid or revoked API key",
          },
        },
      },
    },
    "/v1/finance/swift/{code}": {
      "get": {
        "summary": "Get SWIFT Code",
        "tags": [
          "swift-code",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description": "Look up bank metadata for a SWIFT/BIC code.",
        "parameters": [
          {
            "name": "code",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "DEUTDEDB",
            },
            "description": "SWIFT/BIC code (8 or 11 alphanumeric characters)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "swift_code": {
                          "type": "string",
                          "description": "Full 11-character SWIFT/BIC code",
                        },
                        "bank_code": {
                          "type": "string",
                          "description": "Institution code (characters 1-4)",
                        },
                        "country_code": {
                          "type": "string",
                          "description":
                            "ISO 3166-1 alpha-2 country code (characters 5-6)",
                        },
                        "location_code": {
                          "type": "string",
                          "description": "Location code (characters 7-8)",
                        },
                        "branch_code": {
                          "type": "string",
                          "description":
                            "Branch code (characters 9-11), XXX for primary office",
                        },
                        "bank_name": {
                          "type": "string",
                          "description": "Bank or institution name",
                        },
                        "city": {
                          "type": "string",
                          "description": "City of the branch or primary office",
                        },
                        "country_name": {
                          "type": "string",
                          "description": "Full country name",
                        },
                        "is_primary": {
                          "type": "boolean",
                          "description": "true when branch_code is XXX",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "swift_code": "DEUTDEDBXXX",
                    "bank_code": "DEUT",
                    "country_code": "DE",
                    "location_code": "DB",
                    "branch_code": "XXX",
                    "bank_name": "Deutsche Bank AG",
                    "city": "Frankfurt am Main",
                    "country_name": "Germany",
                    "is_primary": true,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "Invalid SWIFT/BIC format (must be 8 or 11 valid characters).",
          },
          "404": {
            "description": "SWIFT/BIC code not found in the dataset.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/finance/swift": {
      "get": {
        "summary": "List SWIFT Codes",
        "tags": [
          "swift-code",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "List SWIFT records with optional filters and pagination.",
        "parameters": [
          {
            "name": "country_code",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "DE",
            },
            "description":
              "Optional 2-letter country code filter (e.g. DE, US)",
          },
          {
            "name": "bank_code",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "DEUT",
            },
            "description": "Optional 4-letter bank code filter (e.g. DEUT)",
          },
          {
            "name": "q",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "deutsche",
            },
            "description":
              "Optional text search across swift_code, bank_name, and city",
          },
          {
            "name": "limit",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 50,
            },
            "description": "Max rows to return (default 50, max 200)",
          },
          {
            "name": "offset",
            "in": "query",
            "required": false,
            "schema": {
              "type": "integer",
              "example": 0,
            },
            "description": "Number of rows to skip (default 0)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {},
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "items": [
                      {
                        "swift_code": "DEUTDEDBXXX",
                        "bank_code": "DEUT",
                        "country_code": "DE",
                        "location_code": "DB",
                        "branch_code": "XXX",
                        "bank_name":
                          "Deutsche Bank Privat-Und Geschaeftskunden Ag - Head Office",
                        "city": "Frankfurt Am Main",
                        "country_name": "Germany",
                        "is_primary": true,
                      },
                    ],
                    "limit": 50,
                    "offset": 0,
                    "returned": 1,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "Invalid filter or pagination parameter.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/text/similarity": {
      "post": {
        "summary": "Compare Text Similarity",
        "tags": [
          "text-similarity",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Compares two texts and returns a cosine similarity score.",
        "requestBody": {
          "required": true,
          "content": {
            "application/json": {
              "schema": {
                "type": "object",
                "properties": {
                  "text1": {
                    "type": "string",
                    "description": "The first text to compare.",
                    "example": "The cat sat on the mat",
                  },
                  "text2": {
                    "type": "string",
                    "description": "The second text to compare.",
                    "example": "A cat was sitting on a mat",
                  },
                },
                "required": [
                  "text1",
                  "text2",
                ],
                "example": {
                  "text1": "The cat sat on the mat",
                  "text2": "A cat was sitting on a mat",
                },
              },
            },
          },
        },
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "similarity": {
                          "type": "number",
                          "description":
                            "Cosine similarity score between the two texts, in the range [0, 1].",
                        },
                        "method": {
                          "type": "string",
                          "description":
                            "The algorithm used. Currently always 'cosine'.",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "similarity": 0.4364,
                    "method": "cosine",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The request body is missing or malformed.",
          },
          "422": {
            "description": "One or both text fields are missing or empty.",
          },
          "500": {
            "description": "Unexpected server error.",
          },
        },
      },
    },
    "/v1/text/thesaurus/{word}": {
      "get": {
        "summary": "Thesaurus Lookup",
        "tags": [
          "thesaurus",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description": "Returns synonyms and antonyms for the given word.",
        "parameters": [
          {
            "name": "word",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "happy",
            },
            "description": "The word to look up in the thesaurus",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "word": {
                          "type": "string",
                          "description":
                            "The normalized (lowercased) word that was looked up",
                        },
                        "synonyms": {
                          "type": "string",
                          "description": "List of words with similar meaning",
                        },
                        "antonyms": {
                          "type": "string",
                          "description": "List of words with opposite meaning",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "word": "happy",
                    "synonyms": [
                      "joyful",
                      "cheerful",
                      "content",
                      "pleased",
                      "delighted",
                      "glad",
                      "elated",
                      "blissful",
                    ],
                    "antonyms": [
                      "sad",
                      "unhappy",
                      "miserable",
                      "sorrowful",
                      "dejected",
                      "gloomy",
                      "melancholy",
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The word path parameter is missing.",
          },
          "404": {
            "description": "The word was not found in the thesaurus dataset.",
          },
        },
      },
    },
    "/v1/places/timezone": {
      "get": {
        "summary": "Get Timezone",
        "tags": [
          "timezone",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns timezone information for the given coordinates or city name. Provide either `city` or both `lat` and `lon`.",
        "parameters": [
          {
            "name": "lat",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
            },
            "description":
              "Latitude of the location (-90 to 90). Required when using coordinate-based lookup.",
          },
          {
            "name": "lon",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
            },
            "description":
              "Longitude of the location (-180 to 180). Required when using coordinate-based lookup.",
          },
          {
            "name": "city",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
            },
            "description":
              "City name for city-based lookup (e.g. 'Tokyo', 'London'). Required when not using coordinates.",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "timezone": {
                          "type": "string",
                          "description":
                            'IANA timezone identifier (e.g. "Europe/London", "Asia/Tokyo")',
                        },
                        "offset": {
                          "type": "string",
                          "description":
                            "UTC offset in +HH:MM or -HH:MM format (e.g. '+05:30', '-05:00')",
                        },
                        "current_time": {
                          "type": "string",
                          "description":
                            'Current time in UTC, formatted as RFC 3339 (e.g. "2024-12-15T14:30:00Z")',
                        },
                        "is_dst": {
                          "type": "boolean",
                          "description":
                            "Whether the location is currently observing daylight saving time",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "timezone": "Europe/London",
                  "offset": "+00:00",
                  "current_time": "2024-12-15T14:30:00Z",
                  "is_dst": false,
                },
              },
            },
          },
        },
      },
    },
    "/v1/entertainment/trivia": {
      "get": {
        "summary": "Get Trivia Question",
        "tags": [
          "trivia",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a random trivia question with multiple-choice answers. Use the optional category and difficulty query parameters to filter the question pool.",
        "parameters": [
          {
            "name": "category",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
            },
            "description":
              "Filter by category. One of: science, history, geography, sports, music, movies, literature, math, technology, nature.",
          },
          {
            "name": "difficulty",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
            },
            "description": "Filter by difficulty. One of: easy, medium, hard.",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "question": {
                          "type": "string",
                          "description": "The trivia question text",
                        },
                        "options": {
                          "type": "string",
                          "description": "Four multiple-choice answer options",
                        },
                        "answer": {
                          "type": "string",
                          "description":
                            "The correct answer — always one of the values in options",
                        },
                        "category": {
                          "type": "string",
                          "description": "The category the question belongs to",
                        },
                        "difficulty": {
                          "type": "string",
                          "description":
                            "The difficulty level of the question (easy, medium, or hard)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "question":
                      "What is the largest planet in our solar system?",
                    "options": [
                      "Earth",
                      "Jupiter",
                      "Saturn",
                      "Mars",
                    ],
                    "answer": "Jupiter",
                    "category": "science",
                    "difficulty": "easy",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "An invalid category or difficulty value was provided",
          },
          "401": {
            "description": "Missing API key",
          },
          "403": {
            "description": "Invalid or revoked API key",
          },
          "404": {
            "description":
              "No questions match the given category and difficulty combination",
          },
        },
      },
    },
    "/v1/technology/convert": {
      "get": {
        "summary": "Convert Units",
        "tags": [
          "unit-conversion",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description": "Convert a value from one unit to another",
        "parameters": [
          {
            "name": "from",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "miles",
            },
            "description": "Source unit key (e.g. miles, kg, c)",
          },
          {
            "name": "to",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "km",
            },
            "description": "Target unit key (e.g. km, lb, f)",
          },
          {
            "name": "value",
            "in": "query",
            "required": true,
            "schema": {
              "type": "number",
              "example": 10,
            },
            "description": "Numeric value to convert",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "from": {
                          "type": "string",
                          "description": "Source unit key",
                        },
                        "to": {
                          "type": "string",
                          "description": "Target unit key",
                        },
                        "input": {
                          "type": "number",
                          "description": "The original input value",
                        },
                        "result": {
                          "type": "number",
                          "description":
                            "The converted value (rounded to 6 decimal places)",
                        },
                        "formula": {
                          "type": "string",
                          "description": "Human-readable conversion formula",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "from": "miles",
                    "to": "km",
                    "input": 10,
                    "result": 16.09344,
                    "formula": "miles × 1.609344",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
        },
      },
    },
    "/v1/technology/convert/units": {
      "get": {
        "summary": "List Available Units",
        "tags": [
          "unit-conversion",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns all available unit conversion types grouped by measurement category",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "length": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Available length units: millimeter (mm), centimeter (cm), meter (m), kilometer (km), inch (in), foot (ft), yard (yd), mile (miles), nautical mile (nmi)",
                        },
                        "weight": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Available weight units: milligram (mg), gram (g), kilogram (kg), metric ton (t), ounce (oz), pound (lb), stone (stone)",
                        },
                        "volume": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Available volume units: milliliter (ml), liter (l), teaspoon (tsp), tablespoon (tbsp), fluid ounce (fl_oz), cup (cup), pint (pt), quart (qt), gallon (gal)",
                        },
                        "temperature": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Available temperature units: celsius (c), fahrenheit (f), kelvin (k)",
                        },
                        "area": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Available area units: square millimeter (mm2), square centimeter (cm2), square meter (m2), square kilometer (km2), square inch (in2), square foot (ft2), square yard (yd2), acre (acre), hectare (ha)",
                        },
                        "speed": {
                          "type": "array",
                          "items": {},
                          "description":
                            "Available speed units: meters per second (m_s), kilometers per hour (km_h), miles per hour (mph), knots (knots)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "length": [
                      "cm",
                      "ft",
                      "in",
                      "km",
                      "m",
                      "miles",
                      "mm",
                      "nmi",
                      "yd",
                    ],
                    "weight": [
                      "g",
                      "kg",
                      "lb",
                      "mg",
                      "oz",
                      "stone",
                      "t",
                    ],
                    "volume": [
                      "cup",
                      "fl_oz",
                      "gal",
                      "l",
                      "ml",
                      "pt",
                      "qt",
                      "tbsp",
                      "tsp",
                    ],
                    "temperature": [
                      "c",
                      "f",
                      "k",
                    ],
                    "area": [
                      "acre",
                      "cm2",
                      "ft2",
                      "ha",
                      "in2",
                      "km2",
                      "m2",
                      "mm2",
                      "yd2",
                    ],
                    "speed": [
                      "km_h",
                      "knots",
                      "m_s",
                      "mph",
                    ],
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
        },
      },
    },
    "/v1/technology/useragent": {
      "get": {
        "summary": "Parse User Agent",
        "tags": [
          "useragent",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Parses a user agent string and returns structured information about the browser, OS, device, and bot status.",
        "parameters": [
          {
            "name": "ua",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example":
                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
            },
            "description": "The user agent string to parse.",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "browser": {
                          "type": "string",
                          "description":
                            "Detected browser name (e.g. Chrome, Firefox, Safari, Edge, Opera, Internet Explorer, Other)",
                        },
                        "browser_version": {
                          "type": "string",
                          "description":
                            "Detected browser version (major.minor)",
                        },
                        "os": {
                          "type": "string",
                          "description":
                            "Detected operating system (e.g. Windows, macOS, Linux, Android, iOS, ChromeOS, Other)",
                        },
                        "os_version": {
                          "type": "string",
                          "description":
                            "Detected OS version (format varies by platform)",
                        },
                        "device": {
                          "type": "string",
                          "description":
                            "Device type — one of desktop, mobile, tablet, bot, or unknown",
                        },
                        "is_bot": {
                          "type": "boolean",
                          "description":
                            "True when the user agent matches a known bot or crawler pattern",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "browser": "Chrome",
                    "browser_version": "120.0",
                    "os": "Windows",
                    "os_version": "10/11",
                    "device": "desktop",
                    "is_bot": false,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The ua query parameter is missing.",
          },
        },
      },
    },
    "/v1/networking/ip/vpn/{ip}": {
      "get": {
        "summary": "Check IP Address",
        "tags": [
          "vpn-detection",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Analyze an IP address to determine if it's a VPN, proxy, Tor exit node, or hosting provider. Returns detailed threat indicators and scores.",
        "parameters": [
          {
            "name": "ip",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "8.8.8.8",
            },
            "description": "The IP address to check (supports IPv4 and IPv6)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "ip": {
                          "type": "string",
                          "description": "The analyzed IP address",
                        },
                        "is_vpn": {
                          "type": "boolean",
                          "description":
                            "True when the IP belongs to a known VPN provider",
                        },
                        "is_proxy": {
                          "type": "boolean",
                          "description":
                            "True when the IP is a known public or web proxy",
                        },
                        "is_tor": {
                          "type": "boolean",
                          "description":
                            "True when the IP is a known Tor exit node",
                        },
                        "is_hosting": {
                          "type": "boolean",
                          "description":
                            "True when the IP belongs to a data-centre or hosting provider (DCH)",
                        },
                        "score": {
                          "type": "integer",
                          "description":
                            "Raw threat score (0-9+). Tor contributes 3, VPN or Proxy each contribute 2, Hosting contributes 1",
                        },
                        "threat": {
                          "type": "integer",
                          "description":
                            "Threat level derived from score: 0 = None, 1 = Low, 2–3 = Medium, 4–5 = High, 6+ = Critical",
                        },
                        "fraud_score": {
                          "type": "integer",
                          "description":
                            "Fraud risk score from 0 (no risk) to 100 (high risk). Available when using IP2Proxy PX5 or higher",
                        },
                        "asn_org": {
                          "type": "string",
                          "description":
                            'Organization name owning the Autonomous System containing the IP (e.g. "DIGITALOCEAN-ASN")',
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "ip": "8.8.8.8",
                    "is_vpn": false,
                    "is_proxy": false,
                    "is_tor": false,
                    "is_hosting": true,
                    "score": 1,
                    "threat": 1,
                    "fraud_score": 0,
                    "asn_org": "GOOGLE-ASN",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The IP address is missing or invalid",
          },
          "500": {
            "description": "Unexpected server error",
          },
        },
      },
    },
    "/v1/networking/whois/{domain}": {
      "get": {
        "summary": "WHOIS Lookup",
        "tags": [
          "whois",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns WHOIS registration information for a domain name.",
        "parameters": [
          {
            "name": "domain",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "example.com",
            },
            "description": "The domain name to look up (e.g. example.com)",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "domain": {
                          "type": "string",
                          "description": "The domain name that was looked up",
                        },
                        "registrar": {
                          "type": "string",
                          "description":
                            "The name of the registrar holding the domain registration",
                        },
                        "name_servers": {
                          "type": "string",
                          "description":
                            "List of authoritative name servers for the domain",
                        },
                        "status": {
                          "type": "string",
                          "description":
                            "EPP status codes for the domain (e.g. clientTransferProhibited)",
                        },
                        "created_date": {
                          "type": "string",
                          "description":
                            "Date the domain was first registered (ISO 8601)",
                        },
                        "updated_date": {
                          "type": "string",
                          "description":
                            "Date the domain record was last updated (ISO 8601)",
                        },
                        "expiry_date": {
                          "type": "string",
                          "description":
                            "Date the domain registration expires (ISO 8601)",
                        },
                        "dnssec": {
                          "type": "boolean",
                          "description":
                            "True when DNSSEC is enabled for the domain",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "domain": "example.com",
                    "registrar": "RESERVED-Internet Assigned Numbers Authority",
                    "name_servers": [
                      "A.IANA-SERVERS.NET",
                      "B.IANA-SERVERS.NET",
                    ],
                    "status": [
                      "clientDeleteProhibited",
                      "clientTransferProhibited",
                      "clientUpdateProhibited",
                    ],
                    "created_date": "1995-08-14T04:00:00Z",
                    "updated_date": "2023-08-14T07:01:38Z",
                    "expiry_date": "2024-08-13T04:00:00Z",
                    "dnssec": true,
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description": "The domain name format is invalid.",
          },
          "404": {
            "description": "No WHOIS record was found for the domain.",
          },
          "500": {
            "description":
              "Unexpected server error or upstream WHOIS query failure.",
          },
        },
      },
    },
    "/v1/text/words/random": {
      "get": {
        "summary": "Get Random Word",
        "tags": [
          "random-word",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns a random word with its definition and part of speech",
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "id": {
                          "type": "integer",
                          "description": "Unique identifier for the word",
                        },
                        "word": {
                          "type": "string",
                          "description": "The random word",
                        },
                        "definition": {
                          "type": "string",
                          "description": "Dictionary definition of the word",
                        },
                        "part_of_speech": {
                          "type": "string",
                          "description":
                            "Grammatical classification (e.g., noun, verb, adjective, adverb)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "id": 123,
                    "word": "ephemeral",
                    "definition": "lasting for a very short time",
                    "part_of_speech": "adjective",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "503": {
            "description": "No words available in the database",
          },
        },
      },
    },
    "/v1/places/working-days": {
      "get": {
        "summary": "Calculate Working Days",
        "tags": [
          "working-days",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Calculate the number of working days between two dates, optionally accounting for country-specific holidays",
        "parameters": [
          {
            "name": "from",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "2024-02-23",
            },
            "description": "Start date in YYYY-MM-DD format (ISO 8601)",
          },
          {
            "name": "to",
            "in": "query",
            "required": true,
            "schema": {
              "type": "string",
              "example": "2024-02-28",
            },
            "description":
              "End date in YYYY-MM-DD format (ISO 8601). Must be >= from date.",
          },
          {
            "name": "country",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "US",
            },
            "description":
              'ISO 3166-1 alpha-2 country code (e.g., "US", "GB", "FR"). When provided, country-specific holidays are excluded from working days count.',
          },
          {
            "name": "subdivision",
            "in": "query",
            "required": false,
            "schema": {
              "type": "string",
              "example": "NY",
            },
            "description":
              'ISO 3166-2 subdivision code for state/region within the country (e.g., "NY" for New York, "CA" for California). Only used when country is provided.',
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "working_days": {
                          "type": "integer",
                          "description":
                            "Number of working days between the two dates (excluding weekends and optionally holidays)",
                        },
                        "from": {
                          "type": "string",
                          "description": "Start date (echoed from request)",
                        },
                        "to": {
                          "type": "string",
                          "description": "End date (echoed from request)",
                        },
                        "country": {
                          "type": "string",
                          "description":
                            "Country code (echoed from request, empty string if not provided)",
                        },
                        "subdivision": {
                          "type": "string",
                          "description":
                            "Subdivision code (echoed from request, empty string if not provided)",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "working_days": 4,
                    "from": "2024-02-23",
                    "to": "2024-02-28",
                    "country": "US",
                    "subdivision": "NY",
                  },
                  "metadata": {
                    "timestamp": "2026-01-01T00:00:00Z",
                  },
                },
              },
            },
          },
          "400": {
            "description":
              "The from and to parameters are required, or to date is before from date, or invalid date format",
          },
        },
      },
    },
    "/v1/places/time/{timezone}": {
      "get": {
        "summary": "Get Current Time by Timezone",
        "tags": [
          "world-time",
        ],
        "security": [
          {
            "requiems-api-key": [],
          },
        ],
        "description":
          "Returns the current time for the given IANA timezone identifier. The timezone is supplied as a path parameter (e.g. `America/New_York`, `Europe/London`, `UTC`).",
        "parameters": [
          {
            "name": "timezone",
            "in": "path",
            "required": true,
            "schema": {
              "type": "string",
              "example": "America/New_York",
            },
            "description":
              "IANA timezone identifier (e.g. 'America/New_York', 'Europe/London', 'Asia/Kolkata')",
          },
        ],
        "responses": {
          "200": {
            "description": "Successful response",
            "content": {
              "application/json": {
                "schema": {
                  "type": "object",
                  "properties": {
                    "data": {
                      "type": "object",
                      "properties": {
                        "timezone": {
                          "type": "string",
                          "description":
                            'IANA timezone identifier (e.g. "America/New_York")',
                        },
                        "offset": {
                          "type": "string",
                          "description":
                            "UTC offset in +HH:MM or -HH:MM format (e.g. '-05:00', '+05:30')",
                        },
                        "current_time": {
                          "type": "string",
                          "description":
                            'Current time in UTC, formatted as RFC 3339 (e.g. "2024-12-15T14:30:00Z")',
                        },
                        "is_dst": {
                          "type": "boolean",
                          "description":
                            "Whether the timezone is currently observing daylight saving time",
                        },
                      },
                    },
                    "metadata": {
                      "type": "object",
                      "properties": {
                        "timestamp": {
                          "type": "string",
                          "format": "date-time",
                        },
                      },
                    },
                  },
                },
                "example": {
                  "data": {
                    "timezone": "America/New_York",
                    "offset": "-05:00",
                    "current_time": "2024-12-15T14:30:00Z",
                    "is_dst": false,
                  },
                  "metadata": {
                    "timestamp": "2024-12-15T14:30:00Z",
                  },
                },
              },
            },
          },
        },
      },
    },
  },
};

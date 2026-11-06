<p align="center">
  <p align="center">
    <a href="https://requiems.xyz/?utm_source=github&utm_medium=logo" target="_blank">
      <img src="https://raw.githubusercontent.com/bobadilla-tech/requiems-api/refs/heads/main/apps/dashboard/app/assets/images/logo.png" alt="Requiems API" width="280" />
    </a>
  </p>
  <p align="center">
    Real-world examples and tutorials for the Requiems API.
  </p>
  <p align="center">
    <i>A product by <a href="https://bobadilla.tech">Bobadilla Technologies</a></i>
  </p>
</p>

[![Get API Key](https://img.shields.io/badge/Get_API_Key-→-blue)](https://requiems.xyz)
[![Documentation](https://img.shields.io/badge/Documentation-📖-green)](https://requiems.xyz/en/apis)
[![Contributing](https://img.shields.io/badge/Contributing-🤝-orange)](./contributing.md)

---

This repo contains self-contained example projects that show how to use the
[Requiems API](https://requiems.xyz) in real applications — from form validation
to financial dashboards, built with a variety of languages and frameworks.

**One API key. 150+ endpoints. Start building in minutes.**

```bash
curl -X POST https://api.requiems.xyz/v1/email/validate \
  -H "requiems-api-key: YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{"email": "user@gmial.com"}'
# → { "valid": false, "suggestion": "gmail.com" }
```

Get your API key at [requiems.xyz](https://requiems.xyz).

---

## Projects

| Project                                              | Description                                                                                                                                    | Stack                                                           | API Used                  | Difficulty |
| ---------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------- | ------------------------- | ---------- |
| [nextjs-email-validation](./nextjs-email-validation) | Real-time email validation in a signup flow — syntax, MX, disposable checks, typo suggestions, and normalization. Built as a YouTube tutorial. | Next.js 16 · TypeScript · Tailwind v4 · Better Auth · shadcn/ui | `POST /v1/email/validate` | Beginner   |

---

## API Categories

All endpoints are documented at
[requiems.xyz/en/apis](https://requiems.xyz/en/apis).

| Category                                              | Description                                                                               |
| ----------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| [AI & Computer Vision](https://requiems.xyz/en/apis)  | Embeddings, face detection, image-to-text, object detection, sentiment, text similarity   |
| [Animals](https://requiems.xyz/en/apis)               | Facts about animals, cats, and dogs                                                       |
| [Email](https://requiems.xyz/en/apis)                 | Validate, normalize, and check disposable email addresses                                 |
| [Entertainment](https://requiems.xyz/en/apis)         | Advice, jokes, horoscopes, trivia, sudoku, quotes, and more                               |
| [Finance](https://requiems.xyz/en/apis)               | Stock prices, crypto, currency conversion, IBAN, VAT, mortgages, and 30+ more             |
| [Health & Wellness](https://requiems.xyz/en/apis)     | Nutrition, recipes, exercises, hospitals, cocktails                                       |
| [Internet & Technology](https://requiems.xyz/en/apis) | QR codes, barcodes, DNS lookup, IP lookup, WHOIS, password generation, user agent parsing |
| [Miscellaneous](https://requiems.xyz/en/apis)         | Unit conversion, random users, holidays, historical events, baby names                    |
| [Places](https://requiems.xyz/en/apis)                | Geocoding, timezones, weather, air quality, working days, country data                    |
| [Text](https://requiems.xyz/en/apis)                  | Spell check, thesaurus, profanity filter, lorem ipsum, language detection, rhymes         |
| [Transportation](https://requiems.xyz/en/apis)        | Cars, aircraft, airlines, airports, VIN lookup, electric vehicles                         |
| [Weather](https://requiems.xyz/en/apis)               | Current and forecast weather data                                                         |
| [Status](https://requiems.xyz/en/apis)                | API health and uptime                                                                     |

---

## Contributing

Want to add an example? See [contributing.md](./contributing.md) for the full
guide — naming conventions, tooling standards, and a step-by-step walkthrough.

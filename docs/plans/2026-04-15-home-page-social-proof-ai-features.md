# Home Page — Social Proof Carousels & AI Features Card

The home page lacked any social proof and understated the platform's AI-native
capabilities. Two things needed fixing:

1. No visual indication that real companies use the product.
2. The "Built for Developer Experience" section described generic DX features
   but said nothing about AI integrations — a key differentiator given the
   platform ships `llms.txt`, per-API Markdown docs, AI-paste code examples, and
   one-click "Open in Claude / ChatGPT" buttons.

---

## What Changed

### 1. Trusted-By Logo Carousel (between hero and features)

A scrolling logo strip placed immediately below the hero section. It signals
social proof to new visitors before they read any copy.

- Light background (`bg-white / dark:bg-gray-950`) with top and bottom borders
  to sit cleanly between the blue hero and the gray features section.
- Logos scroll left at a comfortable pace (34 s per cycle) and pause on hover.
- Side edges fade to transparent via `mask-image` gradient — the standard
  pattern used by Stripe, Linear, Vercel.
- Each logo sits in a fixed 220 px wide slot so logos with different native
  aspect ratios appear visually equal. Height is capped at 40 px; width is
  `auto` with a 172 px max.
- Logos render in full color at 60% opacity, brightening to 100% on hover with a
  soft indigo glow.

### 2. Social Proof Strip (below the CTA section)

A second, quieter carousel at the very bottom of the page. Reverse scroll
direction and slower speed (40 s) distinguish it from the top strip. Dark
background (`bg-gray-950 / dark:bg-black`) contrasts with the blue CTA gradient
above it and acts as a natural page closer.

### 3. Shared `_logo_marquee` Component

Both strips are powered by a single reusable partial at
`apps/dashboard/app/views/partials/shared/_logo_marquee.html.erb`.

Accepted locals:

| Local         | Default                 | Description                                          |
| ------------- | ----------------------- | ---------------------------------------------------- |
| `logos`       | `[]`                    | Array of `{ file:, alt: }` hashes                    |
| `label`       | `nil`                   | Uppercase label rendered above the strip             |
| `reverse`     | `false`                 | Scroll right-to-left (false) or left-to-right (true) |
| `speed`       | `"32s"`                 | CSS animation duration                               |
| `bg_class`    | light white/border      | Tailwind classes for the outer wrapper               |
| `label_class` | muted gray              | Tailwind classes for the label text                  |
| `logo_class`  | opacity + hover effects | Tailwind classes applied to each `<img>`             |

The animation keyframe name is derived from the `reverse` + `speed` params so
multiple instances on the same page never collide. The track element gets a
`SecureRandom.hex` id for the same reason.

### 4. Features Section — 4th Card: "Built for AI Agents"

The existing three feature cards (Live Playground, Code Examples, Precise Docs)
were joined by a fourth card dedicated to AI integrations.

**Content:**

- Every API ships with `llms.txt` and a full Markdown doc endpoint.
- Code examples include an AI-paste Markdown variant.
- One-click "Open in Claude" and "Open in ChatGPT" buttons on every API page.
- Designed for agents, copilots, and LLM-powered products.

**Layout:** Grid changed from `md:grid-cols-3` to
`sm:grid-cols-2 lg:grid-cols-4` (2 columns on tablet, 4 on desktop). Card uses a
yellow accent to visually distinguish it from the existing blue / green / purple
cards.

The "Precise Documentation" card description was simultaneously reverted to its
original clean copy — AI-specific messaging belongs only on the new dedicated
card.

### 5. Features Subheading Updated

From: `"Everything you need to integrate and ship fast"` To:
`"AI-native integrations, precise docs, and live playgrounds. Ship in minutes, not weeks."`

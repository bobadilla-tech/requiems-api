# Referral MVP

## Problem

The platform has no mechanism for existing users to refer new users or earn
credit for successful referrals. Competitors in the API space use referral
programs to lower acquisition costs and reward loyal users. Without a referral
system:

- Organic word-of-mouth growth cannot be measured or attributed
- No incentive exists for users to promote Requiems API to their networks
- First paid conversion of referred users is invisible to the system
- Stakeholders cannot evaluate referral-based growth as a channel

We need a minimal Referral MVP that captures attribution at signup, detects the
first paid conversion, and exposes a referral link in the dashboard — without
building the reward/payout layer yet.

---

## Goals

1. **Referral attribution** — when user B signs up via user A's referral link,
   record that relationship
2. **First paid conversion tracking** — detect when a referred user upgrades
   to a paid plan for the first time
3. **Self-referral prevention** — users cannot refer themselves or "game" the
   system by creating secondary accounts
4. **Signed referral context** — referral links use tamper-proof signed tokens,
   not raw user IDs
5. **Dashboard visibility** — referred users see a "Referred by" note; referrers
   see a referral link and basic stats
6. **Phase boundaries** — MVP captures data only; rewards/policy/commissions
   are explicitly out of scope

## Non-Goals (for MVP)

- No credit grants, cash payouts, or commissions
- No referral reward policy UI or admin management
- No "tier" system for referrers (bronze/silver/gold etc.)
- No cookie-based or multi-touch attribution
- No sales dashboard or money-sensitive reporting

---

## Current Architecture

### Signup flow (today)

```
User → GET /users/sign_up?plan=developer
     → Devise RegistrationsController#new
     → POST /users (registrations#create)
     → User record created, no referral context captured
```

The `Users::RegistrationsController` permits `name` and `company` only.
There is no parameter for referral attribution.

### Billing checkout flow (today)

```
User → Dashboard::BillingController#checkout (plan, billing_cycle)
     → LemonSqueezy checkout URL with custom_data: { user_id, plan }
     → User pays → LemonSqueezy POST /webhooks/lemonsqueezy
     → Webhooks::LemonsqueezyController#handle_subscription_created
     → Subscription record created, Cloudflare KV synced
```

Custom data is passed to LemonSqueezy and echoed back in webhooks. Today it
carries `user_id` and `plan` only.

### Schema relevant to referrals (today)

| Table           | Relevant columns                                                           |
| --------------- | -------------------------------------------------------------------------- |
| `users`         | `id`, `email`, `name`, `company`, `locale`, `admin`, `created_at`         |
| `subscriptions` | `id`, `user_id`, `plan_name`, `lemonsqueezy_subscription_id`, `status`, … |

There is **no** `referral_code`, `referred_by_id`, or any conversion tracking
column. The `custom_data` blob in LemonSqueezy webhooks is the extension point
for passing referral context into the paid conversion event.

---

## Design Decisions

x### Option A — Extend `users` table with referral columns (rejected)

Add three columns to `users`:

| Column           | Type                | Purpose                                        |
| ---------------- | ------------------- | ---------------------------------------------- |
| `referral_code`  | `string` (uniq idx) | Public token users share (e.g. `A3F9B2C8`)    |
| `referred_by_id` | `bigint FK->users`  | Who referred this user (null if organic)       |
| `referred_at`    | `datetime`          | When the referral relationship was established |

Add one column to `subscriptions`:

| Column                   | Type                | Purpose                                                                    |
| ------------------------ | ------------------- | -------------------------------------------------------------------------- |
| `first_paid_referrer_id` | `bigint FK → users` | FK to referrer if this subscription is the referred user's first paid plan |

**Why rejected:** spreads the referral relationship across two tables
(`users.referred_by_id` for attribution, `subscriptions.first_paid_referrer_id`
for conversion) — more fragmented than a dedicated table, not less. Phase 2
rewards will require per-referral audit rows anyway, forcing a painful migration
out of `users` at that point.

### Option B — Separate `referrals` table ✅ (chosen)

Attribution lives in a dedicated `referrals` table. One column is added to `users`:

| Column          | Type                | Purpose                                     |
| --------------- | ------------------- | ------------------------------------------- |
| `referral_code` | `string` (uniq idx) | Public token users share (e.g. `A3F9B2C8`) |

The `referrals` table:

| Column                       | Type                                  | Purpose                                              |
| ---------------------------- | ------------------------------------- | ---------------------------------------------------- |
| `id`                         | `bigint PK`                           |                                                      |
| `referrer_id`                | `bigint FK → users`                   | The user who shared the referral link                |
| `referred_user_id`           | `bigint FK → users`                   | The user who signed up via the link (unique index)   |
| `status`                     | `string` (enum)                       | `pending` or `converted`                             |
| `converted_at`               | `datetime nullable`                   | When the referred user made their first paid upgrade |
| `converting_subscription_id` | `bigint FK → subscriptions nullable`  | The subscription that triggered conversion           |
| `created_at`                 | `datetime`                            | When the referral relationship was established       |

### Referral link format

```
https://requiems.xyz/:locale/sign_up?ref=<signed_token>
```

The signed token is a Rails `signed_id`-style token encoding the referrer's
`referral_code` (not user ID, to avoid leaking internal IDs). The token is
generated using `Rails.application.message_verifier("referral")`.

**Why signed token over raw `referral_code` in query string:**

- Prevents users from guessing other referral codes and "claiming" a referral
  they weren't given
- Allows optional expiry (e.g., 30-day token TTL) to limit link longevity
- Standard Rails `MessageVerifier` — no external dependencies

### Self-referral prevention

- **Server-side check**: at signup, compare the referred user's email domain
  and name similarity against the referrer's. If the same user is detected,
  drop the referral silently (do not error — just treat as organic signup).
- **No hard-block**: we do not prevent the signup; we simply do not attribute
  it. This avoids a confusing error message and lets legitimate users with
  multiple email addresses still sign up.
- **One referral per user**: a unique index on `referrals.referred_user_id`
  enforces this at the database level. A user cannot be "re-referred."

---

## Implementation Summary

### Database

Migration: `add_referral_code_to_users`

- `referral_code` string, unique index, null: false after backfill

Migration: `create_referrals`

- `referrer_id` bigint, FK → users, not null
- `referred_user_id` bigint, FK → users, not null, unique index
- `status` string, not null, default: `"pending"`
- `converted_at` datetime nullable
- `converting_subscription_id` bigint nullable, FK → subscriptions
- `created_at` datetime, not null
- Index on `referrer_id`

#### Backfilling `referral_code`

Existing users will not have a `referral_code` after the migration adds the
column. A post-migration script or `before_create` callback (only when
`referral_code` is nil) will generate a code using:

```ruby
SecureRandom.alphanumeric(8).upcase  # e.g., "A3F9B2C8"
```

The column starts as nullable, then a second migration (or the same migration
with a data-fix step) sets it to `null: false` after backfill.

### Model changes

| File                        | Change                                                                    |
| --------------------------- | ------------------------------------------------------------------------- |
| `app/models/referral.rb` (new) | `belongs_to :referrer, class_name: "User"`                             |
|                             | `belongs_to :referred_user, class_name: "User"`                          |
|                             | `belongs_to :converting_subscription, class_name: "Subscription", optional: true` |
|                             | `enum status: { pending: "pending", converted: "converted" }`            |
| `app/models/user.rb`        | `has_many :referrals_given, class_name: "Referral", foreign_key: "referrer_id"` |
|                             | `has_one :referral_received, class_name: "Referral", foreign_key: "referred_user_id"` |
|                             | `before_create :ensure_referral_code`                                    |
|                             | `referral_url` helper method                                             |
| `app/models/subscription.rb` | `has_one :referral, foreign_key: "converting_subscription_id"`          |

### Controller changes

| File                                                | Change                                                                           |
| --------------------------------------------------- | --------------------------------------------------------------------------------- |
| `app/controllers/users/registrations_controller.rb`   | Accept `:ref` param in `configure_sign_up_params` (not permitted by Devise yet) |
|                                                     | Move referral logic to `before_action :capture_referral_token`                  |
|                                                     | Store `session[:referral_token]` before Devise processes the signup             |
|                                                     | Assign `referred_by` in `after_create` or custom `create` override             |
| `app/controllers/webhooks/lemonsqueezy_controller.rb` | In `handle_subscription_created`, check if user was referred & this is first paid |
|                                                     | Set `subscription.first_paid_referrer_id = user.referred_by_id`              |
| `app/controllers/dashboard/referrals_controller.rb` (new) | `show` action — displays referral link, referral stats                         |
| `app/controllers/dashboard/billing_controller.rb`    | Pass referral context into LemonSqueezy `custom_data` (if needed for tracking) |

**Note on Devise and `:ref` param:**

Devise's `RegistrationsController#create` does not automatically pass
arbitrary params to the model. The cleanest approach is:

1. In `RegistrationsController#new`, capture `params[:ref]` and store it in
   the session: `session[:referral_token] = params[:ref]`
2. Override `create` or use `after_create` path — assign `referred_by_id`
   based on the session token after the user is persisted
3. Clear `session[:referral_token]` after use

### Routes

| File               | Change                                                        |
| ------------------ | ------------------------------------------------------------- |
| `config/routes.rb` | Add `resources :referrals, only: [:show], controller: "dashboard/referrals"` inside `namespace :dashboard` |

### Dashboard UI

| File                                                    | Change                                              |
| ------------------------------------------------------- | --------------------------------------------------- |
| `app/views/dashboard/referrals/show.html.erb` (new)    | Referral link display, copy-to-clipboard button     |
|                                                         | Stats: total referred, total converted (first paid)  |
| `app/views/partials/referrals/_stats_card.html.erb` (new) | Reusable stats card partial                       |
| `app/views/devise/registrations/new.html.erb`            | Optionally show "Referred by [name]" if referral token |
| `app/views/dashboard/billing/show.html.erb`              | Add referral section in sidebar (link to referrals) |

### LemonSqueezy integration

When a referred user checks out via `Dashboard::BillingController#checkout`,
the `custom_data` hash already sent to LemonSqueezy should include referral
context so the webhook can detect conversion:

```ruby
# In billing_controller.rb checkout method, add to params:
"checkout[custom][referred_by_code]" => current_user.referrer&.referral_code,
```

Then in `Webhooks::LemonsqueezyController#handle_subscription_created`:

```ruby
prior_paid_exists = user.subscriptions.paid.where.not(id: subscription.id).exists?
if user.referred_by_id.present? && !prior_paid_exists
  subscription.first_paid_referrer_id = user.referred_by_id
end
```

### Localization (i18n)

| File                          | Change                     |
| ----------------------------- | -------------------------- |
| `config/locales/en/dashboard.en.yml` | Add `referrals` section keys  |
| `config/locales/es/dashboard.es.yml` | Add `referrals` section keys  |
| `config/locales/fr/dashboard.fr.yml` | Add `referrals` section keys  |

---

## Security Notes

- **Self-referral**: `referred_by_id` is assigned only on user creation and
  never updated. Server-side check compares email/name before assigning.
- **Signed tokens**: Referral links use `Rails.application.message_verifier`
  with a purpose of `"referral"` and optional `expires_in: 30.days`.
  Tampered or expired tokens are silently ignored (user signs up as organic).
- **No internal ID leakage**: The referral token encodes `referral_code`,
  not `id`. `referral_code` is a public-facing random string (8 chars,
  alphanumeric uppercase, ~2.8 trillion combinations).
- **Session storage**: The referral token is briefly stored in the session
  (server-side, signed cookie or Redis depending on config). It is cleared
  after use. No token is passed as a URL param after the initial click.
- **Rate limiting**: No new rate limits needed for MVP — the signup endpoint
  is already protected by Devise and Rails' built-in protections.
- **No reward calculation**: MVP does not grant credits or modify billing.
  This eliminates an entire class of abuse (fraudulent referrals for profit).

---

## Rollout Sequencing

### Phase 1 — Referral MVP (this design)

- [ ] Schema additions (`referral_code`, `referred_by_id`, `referred_at`)
- [ ] Referral link generation + signed token
- [ ] Signup flow captures referral attribution
- [ ] Webhook handler records first paid conversion
- [ ] Dashboard UI: referral link + basic stats
- [ ] i18n for new dashboard section

### Phase 2 — Rewards & Policy (out of scope for MVP)

- [ ] Define reward policy (e.g., "1 month free per 3 paid referrals")
- [ ] Credit grant mechanism (extend `credit_adjustments` or new table)
- [ ] Admin UI for reviewing/approving rewards
- [ ] Email notifications for reward earned
- [ ] Anti-abuse rules (minimum subscription duration before reward)

### Phase 3 — Sales Dashboard (out of scope for MVP)

- [ ] Aggregate referral conversion metrics
- [ ] Revenue attribution by referrer
- [ ] Top referrers leaderboard
- [ ] Export referral data for financial reporting
- [ ] Integration with existing `admin/analytics` namespace

---

## Open Stakeholder Decisions (block later phases)

1. **Reward structure**: What does a referrer get? (account credit, cash,
   service upgrade, swag?)
2. **Reward threshold**: How many paid conversions before a reward is
   granted? Or is it per-conversion?
3. **Minimum subscription duration**: Must a referred user stay paid for 30
   days before the referrer earns a reward? (Prevents "subscribe → cancel"
   abuse.)
4. **Reward budget**: Is there a cap on total referral payouts per month?
5. **Public vs. private**: Will referral links be shareable on social media
   or only via direct link? (Impacts anti-spam measures.)
6. **Terms & conditions**: Legal review needed before rewards go live. What
   disclosures are required?

---

## Verification Checklist

- [ ] Migration runs cleanly: `bin/rails db:migrate`
- [ ] Existing user backfill: all users have a `referral_code`
- [ ] New user signup with `?ref=<signed_token>` → `referred_by_id` set correctly
- [ ] New user signup without token → `referred_by_id` is nil (organic)
- [ ] Signed token expiry: expired token → referral silently ignored
- [ ] Tampered token: modified token → referral silently ignored
- [ ] Self-referral: same email/name → referral silently ignored
- [ ] Referred user upgrades to paid → `subscriptions.first_paid_referrer_id` set
- [ ] Non-referred user upgrades to paid → `first_paid_referrer_id` is nil
- [ ] Dashboard referrals page shows correct referral link
- [ ] Dashboard referrals page shows correct stats (total referred, converted)
- [ ] `bin/rails test` passes
- [ ] `bundle exec brakeman --no-pager` passes
- [ ] `bundle exec rubocop` passes (advisory)

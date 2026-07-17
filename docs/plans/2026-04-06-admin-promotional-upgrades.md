# Admin Promotional Plan Upgrades

## Problem

The current upgrade flow is 100% coupled to Lemon Squeezy. Even 100%-off
discount codes still require users to provide a credit card during checkout.
This creates friction for:

- Students / developers with no card
- Influencers, blog writers, video creators receiving access in exchange for
  content
- Early-access partners or community contributors
- Internal test accounts

We need a first-class way to manually upgrade a user's plan from the admin panel
— with no payment provider involved — while keeping a full audit trail and
automatic expiry.

---

## Requirements

1. **Admin UI** at `/admin/users/:id` — select a plan + duration + reason and
   apply it
2. **Auditability** — clear distinction between paid upgrades and admin-granted
   ones
3. **Temporary** — promotions must expire on a specific date and revert
   automatically
4. **Email notification** — user receives an email after being promoted

---

## Current Architecture

### Upgrade flow (Lemon Squeezy path)

```
User → LemonSqueezy checkout → webhook POST /webhooks/lemonsqueezy
  → Subscription update (plan_name, status, current_period_end)
  → Cloudflare::ApiManagementService.sync_user_plan (PATCH KV + D1)
```

### Key models

| Model / Table   | Relevant columns                                                                                                            |
| --------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `subscriptions` | `plan_name`, `status`, `lemonsqueezy_subscription_id`, `current_period_start`, `current_period_end`, `cancel_at_period_end` |
| `audit_logs`    | `action`, `admin_user_id`, `details` (JSON text), `user_id`                                                                 |

Today the only way to distinguish "paid" from "free" is
`lemonsqueezy_subscription_id IS NOT NULL`. There is no "promoted" concept.

### Cloudflare sync

`Subscription` has an `after_update :sync_to_cloudflare` callback that fires
whenever `plan_name` changes. It calls
`Cloudflare::ApiManagementService#sync_user_plan`, which PATCHes every active
API key in Cloudflare KV and D1. This is already correct and reusable — no
changes needed to the sync path.

---

## Design Decisions

### Option A — Extend `subscriptions` with promotion columns ✅ (chosen)

Add three nullable columns to `subscriptions`:

| Column                 | Type                | Purpose                         |
| ---------------------- | ------------------- | ------------------------------- |
| `promoted_by_id`       | `bigint FK → users` | Admin who granted the promotion |
| `promotion_reason`     | `text`              | Why the promotion was given     |
| `promotion_expires_at` | `datetime`          | When it auto-reverts to free    |

**Differentiating source of a plan:**

| State                  | How to detect                                                     |
| ---------------------- | ----------------------------------------------------------------- |
| Free (organic)         | `lemonsqueezy_subscription_id IS NULL AND promoted_by_id IS NULL` |
| Paid (Lemon Squeezy)   | `lemonsqueezy_subscription_id IS NOT NULL`                        |
| Promoted (admin grant) | `promoted_by_id IS NOT NULL`                                      |

**Why this approach over a separate table:**\
The subscription model is already the single source of truth for a user's
current plan. Splitting "promotional grants" into a separate model would mean
maintaining two sources of truth and adding joins everywhere the current plan is
read. Extending the existing model keeps all plan state in one place and lets
the existing `after_update :sync_to_cloudflare` callback handle Cloudflare sync
automatically.

### Option B — Separate `promotional_grants` table (rejected)

Requires a separate model, controller, and a join or union query whenever the
effective plan is resolved. Adds complexity for no benefit at this scale.

---

## Expiry Handling

A background job (`ExpirePromotionalSubscriptionsJob`) runs hourly and finds
subscriptions where:

- `promoted_by_id IS NOT NULL` (is a promotion)
- `promotion_expires_at <= Time.current` (has expired)
- `plan_name != 'free'` (hasn't already been reverted)

For each match it downgrades to free, clears the promotion fields, syncs
Cloudflare, and writes an audit log entry.

**Why hourly and not daily?**\
Promotions might expire mid-day. A daily job could give up to 23 extra hours of
access beyond the agreed term. Hourly is accurate enough without being
expensive.

**What if the user pays during an active promotion?**\
When a `subscription_created` webhook arrives from Lemon Squeezy, the handler
now clears `promoted_by_id`, `promotion_reason`, and `promotion_expires_at`. The
paid subscription supersedes the promotion. The `lemonsqueezy_subscription_id`
becomes non-null, so the audit trail still shows which path the user is on.

---

## Audit Trail

All promotion actions write to the existing `audit_logs` table:

| Action              | When                                                           |
| ------------------- | -------------------------------------------------------------- |
| `promotion_granted` | Admin grants a promotion via `POST /admin/users/:id/promotion` |
| `promotion_revoked` | Admin revokes early via `DELETE /admin/users/:id/promotion`    |
| `promotion_expired` | Job auto-reverts after `promotion_expires_at`                  |

The `details` column stores JSON: `{ plan_name, expires_at, reason }` for
grants, `{ previous_plan }` for expirations.

---

## Implementation Summary

### Database

Migration: `add_promotion_fields_to_subscriptions`

- `promoted_by_id` bigint nullable, FK → users
- `promotion_reason` text nullable
- `promotion_expires_at` datetime nullable, indexed

### Backend

| File                                                  | Change                                                      |
| ----------------------------------------------------- | ----------------------------------------------------------- |
| `app/models/subscription.rb`                          | `belongs_to :promoted_by`, `promoted?`, `promotional` scope |
| `app/controllers/admin/promotions_controller.rb`      | New: `create`, `destroy`                                    |
| `app/controllers/webhooks/lemonsqueezy_controller.rb` | Clear promotion fields on `subscription_created`            |
| `app/jobs/expire_promotional_subscriptions_job.rb`    | New: hourly expiry job                                      |
| `config/recurring.yml`                                | Add hourly `expire_promotional_subscriptions`               |
| `config/routes.rb`                                    | Nested `resource :promotion` under admin users              |

### Mailer

`PromotionMailer#upgrade_notification` — sends HTML + text email to user with
plan name, expiry date, and a brief personal note.

### Admin UI

New partial `app/views/partials/admin_users/show/_promotion_section.html.erb`
rendered in the admin show page above the "Admin Actions" card.

- **Active promotion**: shows current plan badge, expiry date, who granted it,
  reason, and a "Revoke Promotion" button
- **No active promotion**: shows a form with plan selector, duration presets (1
  mo / 3 mo / 6 mo / 1 yr + custom date), and reason textarea

---

## Security Notes

- The promotion form is inside the `authenticate :user, ->(u) { u.admin? }`
  block in routes — only admins can access it
- `Admin::PromotionsController` also has `before_action :require_admin!`
  inherited from `ApplicationController` (same as other admin controllers)
- `promotion_expires_at` must be a future date — validated server-side
- `plan_name` must be in the allowed list (`developer`, `business`,
  `professional`)
- `reason` is required — prevents silent promotions with no paper trail

---

## Verification Checklist

- [ ] Migration runs cleanly: `bin/rails db:migrate`
- [ ] Free user → admin applies developer plan for 1 month → subscription
      updated, audit log created, email enqueued, Cloudflare synced
- [ ] Admin show page reflects promotion badge with revoke button
- [ ] Admin revokes early → plan reverts to free, audit log created, Cloudflare
      synced
- [ ] Set `promotion_expires_at` to past → run job → plan reverts, audit log
      created
- [ ] Lemon Squeezy `subscription_created` webhook → clears promotion fields
- [ ] `bin/rails test` passes
- [ ] `bundle exec brakeman --no-pager` passes

# Local Payment Testing with ngrok

This guide covers testing the LemonSqueezy payment integration locally using
ngrok. LemonSqueezy needs a public URL to deliver webhooks — ngrok tunnels that
to your local Rails instance.

## How it works

`LEMONSQUEEZY_TEST_MODE=true` is set by default in `.env.example`, so the dev
stack always runs in test mode. `AppConfig` loads all `_TEST` suffixed env vars
(variant IDs, checkout UUIDs, signing secret) when the flag is `true`. To
temporarily disable test mode locally, override with
`LEMONSQUEEZY_TEST_MODE=false` in `infra/docker/.env.local`.

## One-time setup

### Step 1: Get your ngrok hostname

Start your ngrok tunnel and note the public hostname (e.g.
`unprized-unseditiously-marilynn.ngrok-free.dev`). Add it to
`infra/docker/.env` (no protocol, no trailing slash):

```bash
NGROK_HOST=your-hostname.ngrok-free.dev
```

This tells Rails' `HostAuthorization` middleware to allow requests from that
host. Restart the Rails server after changing this value.

### Step 2: Configure the webhook in LemonSqueezy dashboard

1. Enable test mode in LemonSqueezy (toggle in the top bar)
2. Go to Settings > Webhooks > Add webhook
3. Set the URL to:
   `https://<your NGROK_HOST>/webhooks/lemonsqueezy`
4. Select all subscription events:
   - `subscription_created`
   - `subscription_updated`
   - `subscription_cancelled`
   - `subscription_expired`
   - `subscription_resumed`
   - `subscription_payment_success`
5. Save and copy the **Signing Secret**

### Step 3: Add the test signing secret to `.env`

```bash
LEMONSQUEEZY_SIGNING_SECRET_TEST=<signing secret from step 2>
```

This is a separate secret from the production webhook — each webhook in
LemonSqueezy has its own signing secret.

> **Note:** If your ngrok hostname changes, update `NGROK_HOST` in `.env`,
> update the webhook URL in LemonSqueezy, and restart Rails.

---

## Each testing session

### 1. Start local services

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up
```

### 2. Start ngrok

```bash
ngrok http 3000
```

If you have a reserved domain, pass it via `--domain`:

```bash
ngrok http --domain=your-hostname.ngrok-free.dev 3000
```

Confirm the tunnel is live: `https://<your NGROK_HOST>` → `http://localhost:3000`

Test mode is already `true` by default — nothing to toggle.

### 3. Run a test checkout

1. Go to `http://localhost:3000`
2. Register or log in as a user
3. Navigate to Billing → choose a plan → click Upgrade
4. Complete the LemonSqueezy checkout with the test card:
   - Card number: `4242 4242 4242 4242`
   - Expiry: any future date
   - CVV: any 3 digits

### 5. Verify the webhook was received

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml logs -f dashboard | grep LemonSqueezy
```

Expected output:

```text
[LemonSqueezy Webhook] Received: subscription_created
[LemonSqueezy] Subscription created for user 1: developer
```

Confirm the subscription in Rails console:

```bash
docker exec -it requiem-dev-dashboard-1 rails console
User.last.subscription.plan_name   # => "developer"
User.last.subscription.status      # => "active"
```

### 6. Optionally verify Cloudflare KV sync

```bash
cd apps/workers/auth-gateway
npx wrangler kv:key get "key:YOUR_API_KEY" --namespace-id=<CLOUDFLARE_KV_NAMESPACE_ID>
```

Should return the updated plan.

---

## Troubleshooting

### Webhook returns 401 Unauthorized

The signing secret doesn't match.

1. Verify `LEMONSQUEEZY_SIGNING_SECRET_TEST` matches the one shown in
   LemonSqueezy > Settings > Webhooks
2. Check logs:
   `docker compose -f docker-compose.dev.yml logs dashboard | grep "Invalid signature"`
3. If needed, regenerate the signing secret in LemonSqueezy and update `.env`

### Webhook not being delivered

1. Confirm ngrok is running:
   `https://<your NGROK_HOST>/healthz` should return 200
2. Check delivery attempts in LemonSqueezy > Settings > Webhooks > (your
   webhook) > Recent deliveries
3. Make sure LemonSqueezy test mode is ON when triggering test checkouts

### Blocked host errors in Rails logs

If you see `[ActionDispatch::HostAuthorization] Blocked hosts: ...`, your
`NGROK_HOST` env var doesn't match the active tunnel hostname. Update it in
`infra/docker/.env` and restart Rails.

### User not found in webhook

The webhook includes `custom_data.user_id` to identify the buyer. This is
automatically set by `billing_controller.rb` when building the checkout URL.
Make sure you're logged in as a real user in the Rails dashboard before clicking
Upgrade.

### ngrok tunnel not working

Re-run the ngrok command and verify the tunnel URL matches `NGROK_HOST` in
`.env`. If the hostname changed, also update the webhook URL in LemonSqueezy.

### Missing `_TEST` env vars on startup

If Rails fails to start with `MissingConfigError`, check that all `_TEST` vars
are filled in `infra/docker/.env`. With `LEMONSQUEEZY_TEST_MODE=true`, the app
requires all `_TEST` variants to be present.

# Lemon Squeezy webhooks

Rails receives and verifies Lemon Squeezy webhooks, then updates users,
subscriptions, and plans in PostgreSQL. Go reads that same state for quota
and authorization. There is no KV synchronization or API-management
callback.

Configure the signing secret and plan variant IDs through the Rails deployment
environment. Test locally with `RAILS_ENV=test` and a disposable database.
After a webhook, verify the Rails subscription row and then make a safe
read-only Go request with the user's API key.


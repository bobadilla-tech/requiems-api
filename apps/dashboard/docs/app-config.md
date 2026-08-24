# Dashboard configuration

Rails reads configuration through `AppConfig`.

## Runtime API settings

- `INTERNAL_API_URL`: private Go destination; production is
  `http://requiems-api:8080`.
- `API_BASE_URL`: public Cloudflare API URL, normally
  `https://requiems.xyz`.
- `PLAYGROUND_API_KEY`: raw key used only by server-side Playground demos.
- `LOCAL_DEV_API_KEY`: local seed credential; it is not a production secret.

The dashboard no longer reads `API_MANAGEMENT_URL`,
`API_MANAGEMENT_API_KEY`, or the normal `BACKEND_SECRET`.

## Other configuration

Lemon Squeezy, SMTP, Rails secret, and private-deployment checkout variables
are loaded as documented by the deployment configuration. The
private-deployment `tenant_secret` is a separate encrypted customer
credential and remains supported.


# Integration tests

These black-box tests exercise the direct public API at
`https://requiems.xyz` (or `http://localhost:8080` locally). They
verify health, missing/invalid API-key behavior, valid requests, usage headers,
and response shapes.

```bash
cd tests/integration
pnpm install
cp .env.example .env
# set REQUIEMS_API_KEY to a disposable valid key
pnpm test
```

Set `API_BASE_URL` only to the Cloudflare-fronted public hostname or the
isolated local Go API. Do not target a Worker port.


# Load testing

The k6 scenarios target the direct Go API. They require an explicitly supplied
key; the repository no longer seeds `rq_*` Worker fixtures.

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up -d api db redis dashboard
cd ../..
LOCAL_DEV_API_KEY='requiem_<24 alphanumeric characters>' \
  BASE_URL=http://localhost:8080 ./tests/load/run.sh baseline
```

Use disposable isolated test credentials and databases. Production smoke
traffic must use a separate pre-launch key and a safe read-only endpoint; never
use a load-test credential for the required 429 check.

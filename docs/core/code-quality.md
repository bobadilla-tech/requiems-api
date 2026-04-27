# Code Quality

Commands for maintaining code quality across the monorepo.

## Go API (apps/api)

```bash
# Format and lint (auto-fix)
docker exec requiem-dev-api-1 sh -lc 'go fmt ./... && golangci-lint run --fix'

# Run tests
docker exec requiem-dev-api-1 go test ./...

# Run tests with coverage
docker exec requiem-dev-api-1 go test -race -coverprofile=coverage.out ./...

# Run specific test
docker exec requiem-dev-api-1 go test ./services/text/advice -v -run TestGetAdvice
```

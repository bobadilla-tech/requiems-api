FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-build,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/api . && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/seed-bins ./cmd/seed-bins && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/seed-inflation ./cmd/seed-inflation && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/seed-iban ./cmd/seed-iban && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/seed-commodities ./cmd/seed-commodities && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/seed-swift ./cmd/seed-swift && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/seed-exercises ./cmd/seed-exercises

FROM alpine:3.20

RUN adduser -D appuser

WORKDIR /app
ENV PORT=8080

COPY --from=build /out/api /app/api
COPY --from=build /out/seed-* /app/
COPY migrations /app/migrations
COPY dbs /app/dbs

USER appuser

EXPOSE 8080

CMD ["/app/api"]

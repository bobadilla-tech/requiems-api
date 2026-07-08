FROM oven/bun:1 AS build

WORKDIR /app

COPY package.json bun.lock ./
RUN bun install --production --frozen-lockfile

COPY . .

FROM oven/bun:1-slim

WORKDIR /app
ENV MCP_TRANSPORT=http
ENV MCP_HTTP_PORT=3000

COPY --from=build /app /app

USER bun

EXPOSE 3000

CMD ["bun", "run", "src/server.ts"]

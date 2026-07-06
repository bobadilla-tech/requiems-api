FROM oven/bun:1

WORKDIR /app

CMD ["bun", "run", "--watch", "src/server.ts"]

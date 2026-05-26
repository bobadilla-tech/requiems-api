FROM node:24-slim

RUN corepack enable pnpm
RUN corepack prepare pnpm@10.26.0 --activate

WORKDIR /workers/api-management

import { describe, expect, it } from "vitest";
import worker from "../index";
import { authedRequest, makeBindings, makeCtx } from "./helpers";

const PROTECTED_PATH = "/api-keys";

describe("apiKeyAuthMiddleware", () => {
  it("returns 401 when X-API-Management-Key header is absent", async () => {
    const req = new Request(`http://localhost${PROTECTED_PATH}`);
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    expect(res.status).toBe(401);
    const body = (await res.json()) as { error: string };
    expect(body.error).toMatch(/unauthorized/i);
  });

  it("returns 401 when the wrong key is sent", async () => {
    const req = authedRequest(PROTECTED_PATH, {}, "wrong-key-that-is-not-valid-at-all");
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    expect(res.status).toBe(401);
  });

  it("passes through when the correct key is sent", async () => {
    const req = authedRequest(PROTECTED_PATH);
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    // 200 (empty list) means auth succeeded and route handled the request
    expect(res.status).toBe(200);
  });
});

describe("docsMiddleware", () => {
  it("allows access to /docs without credentials in development", async () => {
    const bindings = makeBindings({ ENVIRONMENT: "development" });
    const req = new Request("http://localhost/docs");
    const res = await worker.fetch(req, bindings, makeCtx());

    // /docs renders the SwaggerUI page — anything except 401/403 means it passed auth
    expect(res.status).not.toBe(401);
    expect(res.status).not.toBe(403);
  });

  it("challenges unauthenticated access to /docs in production", async () => {
    const bindings = makeBindings({
      ENVIRONMENT: "production",
      SWAGGER_USERNAME: "admin",
      SWAGGER_PASSWORD: "secret",
    });
    const req = new Request("http://localhost/docs");
    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).toBe(401);
    expect(res.headers.get("WWW-Authenticate")).toMatch(/Basic/);
  });

  it("allows access to /docs with correct credentials in production", async () => {
    const bindings = makeBindings({
      ENVIRONMENT: "production",
      SWAGGER_USERNAME: "admin",
      SWAGGER_PASSWORD: "secret",
    });
    const credentials = btoa("admin:secret");
    const req = new Request("http://localhost/docs", {
      headers: { Authorization: `Basic ${credentials}` },
    });
    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).not.toBe(401);
  });
});

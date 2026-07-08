import { afterEach, beforeEach, describe, expect, test } from "bun:test";
import { apiKeyContext, RequiemsApiError, requiemsRequest } from "./runtime";

const originalFetch = globalThis.fetch;
const originalBaseUrl = process.env.REQUIEMS_BASE_URL;
const originalApiKey = process.env.REQUIEMS_API_KEY;

function mockFetchOnce(response: Response) {
    let capturedUrl: string | undefined;
    let capturedInit: RequestInit | undefined;
    globalThis.fetch = (async (url: string, init?: RequestInit) => {
        capturedUrl = url;
        capturedInit = init;
        return response;
    }) as typeof fetch;
    return {
        get url() {
            return capturedUrl;
        },
        get init() {
            return capturedInit;
        },
    };
}

beforeEach(() => {
    process.env.REQUIEMS_BASE_URL = "https://api.example.test";
    delete process.env.REQUIEMS_API_KEY;
});

afterEach(() => {
    globalThis.fetch = originalFetch;
    if (originalBaseUrl === undefined) delete process.env.REQUIEMS_BASE_URL;
    else process.env.REQUIEMS_BASE_URL = originalBaseUrl;
    if (originalApiKey === undefined) delete process.env.REQUIEMS_API_KEY;
    else process.env.REQUIEMS_API_KEY = originalApiKey;
});

describe("requiemsRequest — URL building", () => {
    test("substitutes path params and serializes query params", async () => {
        const captured = mockFetchOnce(
            new Response(JSON.stringify({ ok: true }), {
                headers: { "content-type": "application/json" },
            }),
        );

        await requiemsRequest({
            method: "GET",
            path: "/v1/finance/bin/{bin}",
            pathParams: { bin: "45717360" },
            query: { extended: "true", skip: undefined },
        });

        expect(captured.url).toBe(
            "https://api.example.test/v1/finance/bin/45717360?extended=true",
        );
    });

    test("strips a trailing slash from REQUIEMS_BASE_URL", async () => {
        process.env.REQUIEMS_BASE_URL = "https://api.example.test/";
        const captured = mockFetchOnce(new Response("{}", { headers: { "content-type": "application/json" } }));

        await requiemsRequest({ method: "GET", path: "/v1/entertainment/chuck-norris" });

        expect(captured.url).toBe("https://api.example.test/v1/entertainment/chuck-norris");
    });

    test("throws a clear error when REQUIEMS_BASE_URL is unset", async () => {
        delete process.env.REQUIEMS_BASE_URL;
        await expect(
            requiemsRequest({ method: "GET", path: "/v1/entertainment/chuck-norris" }),
        ).rejects.toThrow(/REQUIEMS_BASE_URL is not set/);
    });
});

describe("requiemsRequest — auth header", () => {
    test("uses the per-request key from apiKeyContext when present, over the env var", async () => {
        process.env.REQUIEMS_API_KEY = "server-owned-key";
        const captured = mockFetchOnce(
            new Response("{}", { headers: { "content-type": "application/json" } }),
        );

        await apiKeyContext.run("caller-key", () =>
            requiemsRequest({ method: "GET", path: "/v1/entertainment/chuck-norris" }),
        );

        const headers = captured.init?.headers as Record<string, string>;
        expect(headers["requiems-api-key"]).toBe("caller-key");
    });

    test("falls back to REQUIEMS_API_KEY when no per-request key is in context (stdio path)", async () => {
        process.env.REQUIEMS_API_KEY = "server-owned-key";
        const captured = mockFetchOnce(
            new Response("{}", { headers: { "content-type": "application/json" } }),
        );

        await requiemsRequest({ method: "GET", path: "/v1/entertainment/chuck-norris" });

        const headers = captured.init?.headers as Record<string, string>;
        expect(headers["requiems-api-key"]).toBe("server-owned-key");
    });

    test("sends an empty key when neither context nor env var is set", async () => {
        const captured = mockFetchOnce(
            new Response("{}", { headers: { "content-type": "application/json" } }),
        );

        await requiemsRequest({ method: "GET", path: "/v1/entertainment/chuck-norris" });

        const headers = captured.init?.headers as Record<string, string>;
        expect(headers["requiems-api-key"]).toBe("");
    });
});

describe("requiemsRequest — response handling", () => {
    test("parses JSON responses", async () => {
        mockFetchOnce(
            new Response(JSON.stringify({ fact: "hi" }), {
                headers: { "content-type": "application/json" },
            }),
        );
        const result = await requiemsRequest({ method: "GET", path: "/v1/x" });
        expect(result).toEqual({ fact: "hi" });
    });

    test("base64-encodes non-JSON responses (e.g. binary QR/barcode output)", async () => {
        mockFetchOnce(
            new Response(new Uint8Array([1, 2, 3]), {
                headers: { "content-type": "image/png" },
            }),
        );
        const result = await requiemsRequest<string>({ method: "GET", path: "/v1/x" });
        expect(result).toBe(Buffer.from([1, 2, 3]).toString("base64"));
    });

    test("throws RequiemsApiError with the parsed body on non-ok responses", async () => {
        mockFetchOnce(
            new Response(JSON.stringify({ message: "nope" }), {
                status: 401,
                statusText: "Unauthorized",
                headers: { "content-type": "application/json" },
            }),
        );

        await expect(requiemsRequest({ method: "GET", path: "/v1/x" })).rejects.toMatchObject({
            status: 401,
            body: { message: "nope" },
        });
    });

    test("throws RequiemsApiError instances specifically", async () => {
        mockFetchOnce(new Response("boom", { status: 500 }));
        await expect(requiemsRequest({ method: "GET", path: "/v1/x" })).rejects.toBeInstanceOf(
            RequiemsApiError,
        );
    });
});

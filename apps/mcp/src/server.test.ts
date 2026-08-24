import { afterAll, beforeAll, describe, expect, test } from "bun:test";

// Integration tests: spawn the real server process (no mocking of server.ts
// itself) and drive it over real HTTP/stdio, the same way a real MCP client
// would. This is what actually caught the content-wrapping, header-rejection,
// and stdout-pollution bugs earlier — unit tests on the pieces wouldn't have.

const ROOT = new URL("..", import.meta.url).pathname;

let receivedKeys: (string | null)[] = [];
let mcpPort: number;
let mockBackend: ReturnType<typeof Bun.serve> | undefined;
let serverProc: ReturnType<typeof Bun.spawn> | undefined;

async function waitForHealthy(url: string, timeoutMs = 10_000) {
    const deadline = Date.now() + timeoutMs;
    while (Date.now() < deadline) {
        try {
            const res = await fetch(url);
            if (res.ok) return;
        } catch {
            // not up yet
        }
        await new Promise((r) => setTimeout(r, 100));
    }
    throw new Error(`Server at ${url} never became healthy within ${timeoutMs}ms`);
}

/**
 * Consumes a stream exactly once via async iteration (safe for Bun's
 * ReadableStream — unlike manually racing individual reader.read() calls,
 * this never releases a reader lock while a read is still in-flight), until
 * every needle has appeared or the timeout elapses. Returns whatever was
 * accumulated either way, so callers get a useful diff on failure.
 */
async function readStreamUntil(
    stream: ReadableStream<Uint8Array>,
    needles: string[],
    timeoutMs = 10_000,
): Promise<string> {
    const decoder = new TextDecoder();
    let text = "";
    const collect = (async () => {
        for await (const chunk of stream) {
            text += decoder.decode(chunk, { stream: true });
            if (needles.every((n) => text.includes(n))) break;
        }
        return text;
    })();
    return Promise.race([
        collect,
        new Promise<string>((resolve) => setTimeout(() => resolve(text), timeoutMs)),
    ]);
}

async function callTool(name: string, args: Record<string, unknown>, headers: Record<string, string> = {}) {
    const res = await fetch(`http://localhost:${mcpPort}/mcp`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json, text/event-stream",
            ...headers,
        },
        body: JSON.stringify({
            jsonrpc: "2.0",
            id: 1,
            method: "tools/call",
            params: { name, arguments: args },
        }),
    });
    return res.json() as Promise<any>;
}

beforeAll(async () => {
    receivedKeys = [];
    mockBackend = Bun.serve({
        port: 0,
        fetch(req) {
            receivedKeys.push(req.headers.get("requiems-api-key"));
            return new Response(JSON.stringify({ fact: "mock fact" }), {
                headers: { "content-type": "application/json" },
            });
        },
    });

    const proc = Bun.spawn({
        cmd: ["bun", "run", "src/server.ts"],
        cwd: ROOT,
        env: {
            ...process.env,
            MCP_TRANSPORT: "http",
            MCP_HTTP_PORT: "0",
            REQUIEMS_BASE_URL: `http://localhost:${mockBackend.port}`,
        },
        stdout: "pipe",
        stderr: "pipe",
    });
    serverProc = proc;

    const startupLog = await readStreamUntil(
        proc.stderr,
        ["MCP server running (http transport, port"],
    );
    const portMatch = startupLog.match(/MCP server running \(http transport, port (\d+)\)/);
    if (!portMatch) {
        throw new Error(`MCP server failed to report its HTTP port. Startup log:\n${startupLog}`);
    }
    mcpPort = Number(portMatch[1]);

    await waitForHealthy(`http://localhost:${mcpPort}/healthz`);
});

afterAll(async () => {
    if (serverProc) {
        serverProc.kill();
        await serverProc.exited;
    }
    if (mockBackend) {
        await mockBackend.stop(true);
    }
});

describe("HTTP transport", () => {
    test("/healthz returns 200", async () => {
        const res = await fetch(`http://localhost:${mcpPort}/healthz`);
        expect(res.status).toBe(200);
    });

    test("rejects a tool call with no requiems-api-key header, without hitting the backend", async () => {
        const before = receivedKeys.length;
        const body = await callTool("entertainment_chuck_norris", {});
        expect(body.result.isError).toBe(true);
        expect(body.result.content[0].text).toBe("Missing requiems-api-key header");
        expect(receivedKeys.length).toBe(before); // never reached the backend
    });

    test("rejects a tool call with an empty requiems-api-key header", async () => {
        const body = await callTool("entertainment_chuck_norris", {}, { "requiems-api-key": "" });
        expect(body.result.isError).toBe(true);
        expect(body.result.content[0].text).toBe("Missing requiems-api-key header");
    });

    test("forwards the caller's own key and wraps the real result in content", async () => {
        const body = await callTool("entertainment_chuck_norris", {}, {
            "requiems-api-key": "test-caller-key",
        });
        expect(body.result.isError).toBeUndefined();
        expect(JSON.parse(body.result.content[0].text)).toEqual({ fact: "mock fact" });
        expect(receivedKeys.at(-1)).toBe("test-caller-key");
    });

    test("two concurrent callers each get their own key forwarded (no cross-request leak)", async () => {
        await Promise.all([
            callTool("entertainment_chuck_norris", {}, { "requiems-api-key": "caller-A" }),
            callTool("entertainment_chuck_norris", {}, { "requiems-api-key": "caller-B" }),
        ]);
        const last2 = receivedKeys.slice(-2).sort();
        expect(last2).toEqual(["caller-A", "caller-B"]);
    });

    test("initialize handshake works without any api key header", async () => {
        const res = await fetch(`http://localhost:${mcpPort}/mcp`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Accept: "application/json, text/event-stream",
            },
            body: JSON.stringify({
                jsonrpc: "2.0",
                id: 1,
                method: "initialize",
                params: {
                    protocolVersion: "2025-06-18",
                    capabilities: {},
                    clientInfo: { name: "test", version: "0" },
                },
            }),
        });
        const body = (await res.json()) as any;
        expect(body.result.serverInfo.name).toBe("requiems-api");
    });
});

describe("stdio transport", () => {
    test("boots cleanly with stdout reserved for JSON-RPC only (no [server] logs on stdout)", async () => {
        const proc = Bun.spawn({
            cmd: ["bun", "run", "src/server.ts"],
            cwd: ROOT,
            env: {
                ...process.env,
                MCP_TRANSPORT: "stdio",
                REQUIEMS_API_KEY: "requiem_test_key",
                REQUIEMS_BASE_URL: "https://api.example.test",
            },
            stdout: "pipe",
            stderr: "pipe",
        });

        const stderr = await readStreamUntil(proc.stderr, [
            "Registered",
            "MCP server running (stdio transport)",
        ]);

        proc.kill();
        await proc.exited;

        const stdout = await new Response(proc.stdout).text();

        expect(stdout).toBe("");
        expect(stderr).toContain("Registered");
        expect(stderr).toContain("MCP server running (stdio transport)");
    });

    test("fails fast with a clear stderr message when REQUIEMS_BASE_URL is missing", async () => {
        const proc = Bun.spawn({
            cmd: ["bun", "run", "src/server.ts"],
            cwd: ROOT,
            env: {
                ...process.env,
                MCP_TRANSPORT: "stdio",
                REQUIEMS_API_KEY: "requiem_test_key",
                REQUIEMS_BASE_URL: "",
            },
            stdout: "pipe",
            stderr: "pipe",
        });

        const exitCode = await proc.exited;
        const stderr = await new Response(proc.stderr).text();

        expect(exitCode).toBe(1);
        expect(stderr).toContain("Missing REQUIEMS_BASE_URL");
    });
});

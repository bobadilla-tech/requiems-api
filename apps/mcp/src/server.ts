import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { WebStandardStreamableHTTPServerTransport } from "@modelcontextprotocol/sdk/server/webStandardStreamableHttp.js";
import { tools } from "../generated/index.ts";
import { apiKeyContext } from "../generated/runtime.ts";


// Validate required runtime env (used by generated handlers)
const MCP_TRANSPORT = process.env.MCP_TRANSPORT ?? "stdio";
const REQUIEMS_API_KEY = process.env.REQUIEMS_API_KEY;
const REQUIEMS_BASE_URL = process.env.REQUIEMS_BASE_URL;

if (!REQUIEMS_BASE_URL) {
    console.error("[server] Missing REQUIEMS_BASE_URL");
    process.exit(1);
}

// In HTTP mode the caller's own key arrives per-request via the
// `requiems-api-key` header, so there's no static key to require here.
if (MCP_TRANSPORT !== "http" && !REQUIEMS_API_KEY) {
    console.error("[server] Missing REQUIEMS_API_KEY");
    process.exit(1);
}


// -----------------------------------------------------------------------------
// Tool registration with defensive guards
// -----------------------------------------------------------------------------

function registerTools(server: McpServer) {
    if (!Array.isArray(tools)) {
        throw new Error("Generated tools export is invalid (expected array)");
    }

    for (const tool of tools) {
        if (!tool?.name || typeof tool.handler !== "function") {
            throw new Error(`Invalid tool definition detected: ${JSON.stringify(tool)}`);
        }

        try {
            server.tool(
                tool.name,
                tool.description ?? tool.name,
                tool.inputSchema?.shape ?? tool.inputSchema, // defensive compatibility
                async (args: any, extra: any) => {
                    // Under HTTP transport, extra.requestInfo.headers carries the caller's
                    // own key; under stdio there's no requestInfo, so runtime.ts falls back
                    // to the process-wide REQUIEMS_API_KEY.
                    const requestInfo = extra?.requestInfo;
                    const callerKey = requestInfo?.headers?.["requiems-api-key"];

                    if (requestInfo && !callerKey) {
                        return {
                            content: [{ type: "text", text: "Missing requiems-api-key header" }],
                            isError: true,
                        };
                    }

                    try {
                        const result = await apiKeyContext.run(callerKey ?? "", () => tool.handler(args));
                        // Generated handlers return the raw Requiems API result, not an MCP
                        // CallToolResult — wrap it here so clients actually see the data
                        // (the SDK validates the return against CallToolResultSchema, which
                        // defaults `content` to [] when it's missing).
                        return {
                            content: [
                                {
                                    type: "text",
                                    text: typeof result === "string" ? result : JSON.stringify(result),
                                },
                            ],
                        };
                    } catch (err) {
                        console.error(`[tool:${tool.name}] execution failed`, err);
                        throw err;
                    }
                },
            );
        } catch (err) {
            console.error(`[server] Failed to register tool: ${tool.name}`, err);
            throw err;
        }
    }

    console.error(`[server] Registered ${tools.length} tools`);
}

function createServer(): McpServer {
    const server = new McpServer({
        name: "requiems-api",
        version: "0.1.0",
    });
    registerTools(server);
    return server;
}

// -----------------------------------------------------------------------------
// Transport + lifecycle
// -----------------------------------------------------------------------------

async function startStdio() {
    const server = createServer();
    const transport = new StdioServerTransport();
    await server.connect(transport);
    console.error("[server] MCP server running (stdio transport)");
}

async function startHttp() {
    const port = Number(process.env.MCP_HTTP_PORT ?? 3000);

    Bun.serve({
        port,
        async fetch(req: Request) {
            const url = new URL(req.url);

            if (url.pathname === "/healthz") {
                return new Response("ok", { status: 200 });
            }

            // Stateless mode: a fresh server + transport per request, per SDK contract
            // ("Stateless transport cannot be reused across requests").
            const server = createServer();
            const transport = new WebStandardStreamableHTTPServerTransport({
                sessionIdGenerator: undefined,
                enableJsonResponse: true,
            });

            try {
                await server.connect(transport);
                return await transport.handleRequest(req);
            } finally {
                await transport.close();
                await server.close();
            }
        },
        error(err) {
            console.error("[server] Unhandled HTTP error:", err);
            return new Response(
                JSON.stringify({
                    jsonrpc: "2.0",
                    id: null,
                    error: { code: -32603, message: "Internal server error" },
                }),
                { status: 500, headers: { "content-type": "application/json" } },
            );
        },
    });

    console.error(`[server] MCP server running (http transport, port ${port})`);
}

async function start() {
    try {
        console.error("[server] Starting MCP server...");

        if (MCP_TRANSPORT === "http") {
            await startHttp();
        } else {
            await startStdio();
        }
    } catch (err) {
        console.error("[server] Fatal startup error:", err);
        process.exit(1);
    }
}

start();

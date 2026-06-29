import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { tools } from "../generated/index";


// Validate required runtime env (used by generated handlers)
const REQUIEMS_API_KEY = process.env.REQUIEMS_API_KEY;
const REQUIEMS_BASE_URL = process.env.REQUIEMS_BASE_URL;

if (!REQUIEMS_API_KEY) {
    console.error("[server] Missing REQUIEMS_API_KEY");
    process.exit(1);
}


// -----------------------------------------------------------------------------
// MCP Server setup
// -----------------------------------------------------------------------------

const server = new McpServer({
    name: "requiems-api",
    version: "0.1.0",
});

// -----------------------------------------------------------------------------
// Tool registration with defensive guards
// -----------------------------------------------------------------------------

function registerTools() {
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
                async (args: any) => {
                    try {
                        return await tool.handler(args);
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

    console.log(`[server] Registered ${tools.length} tools`);
}

// -----------------------------------------------------------------------------
// Transport + lifecycle
// -----------------------------------------------------------------------------

const transport = new StdioServerTransport();

async function start() {
    try {
        console.log("[server] Starting MCP server...");

        registerTools();

        await server.connect(transport);

        console.log("[server] MCP server running (stdio transport)");
    } catch (err) {
        console.error("[server] Fatal startup error:", err);
        process.exit(1);
    }
}

start();
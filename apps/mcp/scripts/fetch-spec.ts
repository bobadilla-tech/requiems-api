const SPEC_URL = "https://api.requiems.xyz/openapi.json";

console.log(`Downloading OpenAPI spec from ${SPEC_URL}...`);

const res = await fetch(SPEC_URL);

if (!res.ok) {
    throw new Error(`Failed to download spec: ${res.status} ${res.statusText}`);
}

await Bun.write("openapi.json", await res.text());

console.log("✓ Saved to openapi.json");
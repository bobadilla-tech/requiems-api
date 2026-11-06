# Developer Note

Full spec and architecture: [DESIGN.md](DESIGN.md). The generalized pipeline
pattern, if you're adapting this for another API:
[HOW_TO_BUILD_AGENT_SKILLS.md](HOW_TO_BUILD_AGENT_SKILLS.md).

To regenerate skills locally against a checkout of `requiems-api`:

```bash
node scripts/build/index.ts --source ../requiems-api/apps/dashboard/config/api_docs --output ./skills
```

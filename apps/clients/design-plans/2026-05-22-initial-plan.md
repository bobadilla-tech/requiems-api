# Spike: OpenAPI Client Generator Evaluation

The goal of this proof of concept was to evaluate and choose the best tool for
automatically generating HTTP API clients from the Requiems OpenAPI
specification.

The generated clients should reduce manual boilerplate, avoid type drift between
the API and client packages, improve developer experience, and give us a stable
foundation for publishing official clients in multiple ecosystems.

The generators were tested against the full Gateway OpenAPI specification.

## Tools Evaluated

We tested the main options that could fit this repository:

1. **OpenAPI Generator CLI** via
   [`OpenAPITools/openapi-generator-cli`](https://github.com/OpenAPITools/openapi-generator-cli)
2. **Kiota** by Microsoft
3. **Orval** for TypeScript and JavaScript clients

Testing the alternatives gave us a practical comparison instead of choosing only
from documentation or assumptions.

## Comparison Table

| Tool                      | Main Ecosystem          | Code Weight        | Adoption Curve | License    | Best Fit                                              |
| :------------------------ | :---------------------- | :----------------- | :------------- | :--------- | :---------------------------------------------------- |
| **OpenAPI Generator CLI** | Multi-language          | Heavy but complete | Medium         | Apache 2.0 | Public SDKs and multi-language client generation      |
| **Kiota**                 | C# / .NET / TypeScript  | Lightweight        | Medium         | MIT        | Modern .NET backends and builder-style API clients    |
| **Orval**                 | TypeScript / JavaScript | Lightweight        | Low            | MIT        | Frontend or Node.js clients with Axios/fetch patterns |

## Detailed Analysis

### OpenAPI Generator CLI

OpenAPI Generator CLI is a mature and widely adopted generator that supports a
large number of languages and client styles.

**Pros**

- Strong multi-language support, including TypeScript, Python, and C#.
- Mature project with broad community usage.
- Good documentation and a large amount of existing examples.
- Generates complete clients that are ready to package and publish.
- Stable enough for long-term automation in CI/CD.
- Works well for an official SDK repository where consistency matters more than
  producing the smallest possible generated output.

**Cons**

- The generated code can be verbose.
- Some generated clients include more boilerplate than a hand-written client.
- Pull requests can be large because generated SDKs naturally touch many files.

### Kiota

Kiota is a modern generator from Microsoft focused on strongly typed clients and
a fluent builder-style API.

**Pros**

- Clean builder pattern with strong editor autocomplete.
- Good fit for modern .NET projects.
- Lightweight generated models.
- Integrates well with `HttpClient` and modern dependency injection patterns.

**Cons**

- Requires more manual setup for authentication providers and request adapters.
- Generated folder structures can become deep because they mirror URL segments.
- Less suitable as the single standard for all target package ecosystems in this
  repository.

### Orval

Orval is focused on TypeScript and JavaScript client generation.

**Pros**

- Generates readable, idiomatic TypeScript.
- Supports functional clients using `axios` or `fetch`.
- Good fit for frontend projects and Node.js clients.
- Can generate useful frontend integrations, such as React Query helpers.

**Cons**

- TypeScript/JavaScript only.
- Does not help with Python or C# package generation.
- Would require a separate generator strategy for the other official clients.

## Final Decision

We decided to use
[`OpenAPITools/openapi-generator-cli`](https://github.com/OpenAPITools/openapi-generator-cli)
as the standard generator for this repository.

The main reasons are stability, industry adoption, strong documentation, and
multi-language support. Even though the generated code is heavier than Kiota or
Orval, OpenAPI Generator CLI gives us the most reliable foundation for producing
official clients across npm, PyPI, and C# packages from the same OpenAPI
contract.

The other tools were still valuable to test. Kiota showed a strong modern .NET
experience, and Orval produced excellent TypeScript output, but using separate
generators per language would add operational complexity. For this repository,
standardization and release automation are more important than optimizing each
language client independently.

## Current Generation Approach

The clients should be generated from the canonical OpenAPI document:

```text
https://api.requiems.xyz/openapi.json
```

OpenAPI Generator CLI should be used to generate the supported package clients,
including:

- TypeScript client for npm
- Python client for PyPI
- C# client for NuGet

## Installation and Local Generation

All commands should be run from the repository root.

### Install OpenAPI Generator CLI

```sh
npm install @openapitools/openapi-generator-cli
```

The generator can also be executed with `npx` if it is not installed globally.

### Fetch the OpenAPI Specification

```sh
curl -o openapi.json https://api.requiems.xyz/openapi.json
```

### Generate the TypeScript Client

```sh
npx openapi-generator-cli generate \
  -i openapi.json \
  -g typescript-axios \
  -o typescript
```

### Generate the Python Client

```sh
npx openapi-generator-cli generate \
  -i openapi.json \
  -g python \
  -o python
```

### Generate the C# Client

```sh
npx openapi-generator-cli generate \
  -i openapi.json \
  -g csharp \
  -o csharp
```

## Next Steps

### 1. Weekly Client Regeneration Workflow

Create a GitHub Actions workflow that runs periodically once per week.

The workflow should:

- Fetch the latest OpenAPI JSON file from
  `https://api.requiems.xyz/openapi.json`.
- Regenerate the TypeScript, Python, and C# clients with OpenAPI Generator CLI.
- Detect whether the generated clients changed.
- Bump the client package version when there are changes.
- Open a pull request containing the updated OpenAPI file, regenerated clients,
  and version bump.

This keeps the client packages aligned with the live API contract while still
requiring review before changes are merged.

### 2. Publish on Release Workflow

Create a second GitHub Actions workflow called **Publish on Release**.

This workflow should run when a GitHub release is created manually. It should
publish the generated client packages to:

- npm for the TypeScript package
- PyPI for the Python package
- NuGet for the C# package

This separates automatic regeneration from package publishing. Regeneration can
happen weekly through pull requests, while publishing remains an intentional
manual release action.

## Conclusion

After testing the available generators, we chose OpenAPI Generator CLI because
it is the most stable and industry-standard option for this repository. Its
documentation, ecosystem support, and multi-language generation make it the best
fit for maintaining official Requiems API clients across TypeScript, Python, and
C#.

Kiota and Orval were useful comparison points, and both are strong tools in
their specific ecosystems. However, OpenAPI Generator CLI gives us the simplest
long-term operating model: one generator, one OpenAPI contract, repeatable
weekly client updates, and manual publishing through GitHub releases.

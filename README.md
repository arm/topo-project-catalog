# Topo Project Catalog

Catalog generation and release tooling for Topo Projects.

## Contents

- `data/catalog.schema.json` — JSON Schema for the catalog format.
- `data/github_sources.json` — source repositories and pinned commits used to generate the catalog.

## Updating the catalog

The release workflow generates `data/catalog.json` and publishes it as a release asset. It is not stored in the repository.

Edit `data/github_sources.json` to add, remove, or change pinned project repositories. To generate the catalog locally, run:

```sh
CATALOG_VERSION=v1.2.3 go run ./cmd/generate-catalog
```

The updater uses `GITHUB_TOKEN` if present, which is recommended to avoid GitHub API rate limits:

```sh
GITHUB_TOKEN=... CATALOG_VERSION=v1.2.3 go run ./cmd/generate-catalog
```

## Validation

Run the test suite with:

```sh
go test ./...
```

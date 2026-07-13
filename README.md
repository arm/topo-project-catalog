# Topo Project Catalog

Catalog data and update tooling for Topo Projects.

## Contents

- `data/catalog.json` — generated catalog of Topo Project repositories.
- `data/catalog.schema.json` — JSON Schema for the catalog format.
- `data/github_sources.json` — source repositories and pinned commits used to generate the catalog.

## Updating the catalog

Edit `data/github_sources.json` to add, remove, or change pinned project repositories, then run:

```sh
go run ./cmd/update-catalog
```

The updater uses `GITHUB_TOKEN` if present, which is recommended to avoid GitHub API rate limits:

```sh
GITHUB_TOKEN=... go run ./cmd/update-catalog
```

## Validation

Run the test suite with:

```sh
go test ./...
```

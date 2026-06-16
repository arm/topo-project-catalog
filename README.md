# Topo Project Catalog

Catalog data and update tooling for Topo Template projects.

## Contents

- `data/catalog.json` — generated catalog of Topo Template repositories.
- `data/catalog.schema.json` — JSON Schema for the catalog format.
- `data/github_sources.json` — source repositories and pinned commits used to generate the catalog.

## Updating the catalog

Edit `data/github_sources.json` to add, remove, or change pinned template repositories, then run:

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

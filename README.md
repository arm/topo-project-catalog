# Topo Project Catalog

Catalog generation and release tooling for Topo Projects.

## Contents

- `data/catalog.schema.json` — JSON Schema for the catalog format.
- `data/github_sources.json` — source repositories and pinned commits used to generate the catalog.
- `docs/project-acceptance-criteria.md` — criteria for including a Project in the catalog.

## Updating the catalog

The release workflow generates `data/catalog.json` and publishes it as a release asset. It is not stored in the repository.

Before proposing a Project, review the [Project acceptance criteria](docs/project-acceptance-criteria.md).

Edit `data/github_sources.json` to add, remove, or change pinned project repositories. To generate the catalog locally, run:

```sh
CATALOG_VERSION=v1.2.3 go run ./cmd/generate-catalog
```

The updater uses `GITHUB_TOKEN` if present, which is recommended to avoid GitHub API rate limits:

```sh
GITHUB_TOKEN=... CATALOG_VERSION=v1.2.3 go run ./cmd/generate-catalog
```

## Refreshing source SHAs locally

Use `scripts/refresh_github_sources.py` to update all project pins to their latest default-branch commits before generating a catalog:

```sh
gh auth login --hostname github.com
gh auth status
python3 scripts/refresh_github_sources.py
git diff -- data/github_sources.json
```
The script updates only the local `data/github_sources.json` file.
The daily/manual workflow runs the same script and opens or updates a PR when changes exist.

## Validation

Run the test suite with:

```sh
go test ./...
```

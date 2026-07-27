package updater

import (
	"log"
	"os"
)

func Run() {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Println("⚠️ GITHUB_TOKEN is not set: you might get rate limited")
	}

	githubClient := NewGitHubClient(githubToken)

	sourcesFilePath, err := GithubSourcesFilePath()
	if err != nil {
		log.Fatalf("failed to find sources file: %v\n", err)
	}

	sources, err := ListGithubSources(sourcesFilePath)
	if err != nil {
		log.Fatalf("failed to list sources: %v\n", err)
	}

	catalogSchemaFilePath, err := CatalogSchemaFilePath()
	if err != nil {
		log.Fatalf("failed to find catalog schema file: %v\n", err)
	}

	validator, err := NewCatalogSchema(catalogSchemaFilePath)
	if err != nil {
		log.Fatalf("failed to create schema validator: %v\n", err)
	}

	projects := make([]Project, 0, len(sources))
	for _, source := range sources {
		project, err := FetchProject(githubClient, source)
		if err != nil {
			log.Fatalf("failed to fetch %s: %v\n", source, err)
		}
		if err := validator.ValidateProject(project); err != nil {
			log.Fatalf("invalid project %s: %v\n", source, err)
		}
		log.Printf("fetched %s\n", source)
		projects = append(projects, project)
	}

	document := Catalog{
		Schema:   validator.SchemaURL(),
		Projects: projects,
	}
	if err := validator.ValidateCatalog(document); err != nil {
		log.Fatalf("invalid catalog file: %v\n", err)
	}
	catalogFilePath, err := CatalogFilePath()
	if err != nil {
		log.Fatalf("failed to find catalog file: %v\n", err)
	}
	if err := WriteCatalog(catalogFilePath, document); err != nil {
		log.Fatalf("failed to write catalog file: %v\n", err)
	}
	log.Printf("written catalog to %s\n", catalogFilePath)
}

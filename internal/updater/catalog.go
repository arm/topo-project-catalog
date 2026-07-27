package updater

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	relativeCatalogOutputPath = "data/catalog.json"
)

type Catalog struct {
	Schema   string    `json:"$schema"`
	Projects []Project `json:"projects"`
}

func WriteCatalog(path string, document Catalog) error {
	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create catalog output: %w", err)
	}
	enc := json.NewEncoder(outputFile)
	enc.SetIndent("", "  ")
	writeErr := enc.Encode(document)
	closeErr := outputFile.Close()
	if writeErr != nil {
		return fmt.Errorf("failed to write projects: %w", writeErr)
	}
	if closeErr != nil {
		return fmt.Errorf("failed to close catalog output: %w", closeErr)
	}
	return nil
}

func CatalogFilePath() (string, error) {
	repoRoot, err := findModuleRoot()
	if err != nil {
		return "", err
	}

	return filepath.Join(repoRoot, filepath.FromSlash(relativeCatalogOutputPath)), nil
}

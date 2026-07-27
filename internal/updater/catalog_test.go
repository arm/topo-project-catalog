package updater

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteCatalog(t *testing.T) {
	t.Run("writes catalog document to file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "catalog.json")
		want := Catalog{
			Schema: "https://example.com/schema.json",
			Projects: []Project{
				{
					XTopo: XTopo{
						Name:        "death-star-trench-run",
						Description: "Use the Force to benchmark impossible shots",
						Features:    []string{"X-wing", "Astromech", "Proton torpedoes"},
					},
					URL: "ssh://death-star.example",
					Ref: "rebellion",
				},
			},
		}

		err := WriteCatalog(path, want)

		require.NoError(t, err)
		gotBytes, err := os.ReadFile(path)
		require.NoError(t, err)
		var got Catalog
		err = json.Unmarshal(gotBytes, &got)
		require.NoError(t, err)
		assert.Equal(t, want, got)
	})
}

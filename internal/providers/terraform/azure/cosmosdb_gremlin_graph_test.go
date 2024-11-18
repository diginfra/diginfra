package azure_test

import (
	"testing"

	"github.com/diginfra/diginfra/internal/providers/terraform/tftest"
)

func TestAzureRMCosmosDBGremlinGraphGoldenFile(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "cosmosdb_gremlin_graph_test")
}

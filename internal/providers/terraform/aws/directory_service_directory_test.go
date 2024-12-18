package aws_test

import (
	"testing"

	"github.com/diginfra/diginfra/internal/providers/terraform/tftest"
)

func TestDirectoryServiceDirectoryGoldenFile(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "directory_service_directory_test")
}

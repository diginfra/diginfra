package aws_test

import (
	"testing"

	"github.com/diginfra/diginfra/internal/providers/terraform/tftest"
)

func TestMWAAEnvironmentGoldenFile(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "mwaa_environment_test")
}

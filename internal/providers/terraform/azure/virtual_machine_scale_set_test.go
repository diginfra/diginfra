package azure_test

import (
	"testing"

	"github.com/diginfra/diginfra/internal/providers/terraform/tftest"
)

func TestAzureVirtualMachineScaleSetGoldenFile(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "virtual_machine_scale_set_test")
}

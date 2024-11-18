package azure_test

import (
	"testing"

	"github.com/diginfra/diginfra/internal/providers/terraform/tftest"
)

func TestAzureRMNotificationHubNamespace(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "notification_hub_namespace_test")
}

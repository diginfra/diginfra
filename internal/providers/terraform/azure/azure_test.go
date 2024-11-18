package azure_test

import (
	"os"
	"testing"

	"github.com/diginfra/diginfra/internal/providers/terraform/tftest"
)

func TestMain(m *testing.M) {
	tftest.EnsurePluginsInstalled()
	code := m.Run()
	os.Exit(code)
}

package main_test

import (
	"github.com/diginfra/diginfra/internal/testutil"
	"testing"
)

func TestAuthLoginHelpFlag(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"auth", "login", "--help"}, nil)
}

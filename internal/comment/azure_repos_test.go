package comment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildAzureAPIURL(t *testing.T) {
	tests := []struct {
		repoURL string
		want    string
	}{
		{
			repoURL: "https://diginfra-user@dev.azure.com/diginfra/my%20project/_git/my%20repo",
			want:    "https://dev.azure.com/diginfra/my%20project/_apis/git/repositories/my%20repo/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.repoURL, func(t *testing.T) {
			got, err := buildAzureAPIURL(tt.repoURL)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

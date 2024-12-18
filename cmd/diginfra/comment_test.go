package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/diginfra/diginfra/internal/testutil"
)

func TestComment(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"comment"}, nil)
}

func TestCommentHelp(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"comment", "--help"}, nil)
}

func TestCommentBackoffRetry(t *testing.T) {
	var attempts int

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Helper()

		attempts += 1
		assert.Equal(t, "/api/v3/repos/diginfra/diginfra/issues/8/comments", r.RequestURI)
		if (attempts % 3) < 2 {
			w.WriteHeader(400)
			return
		}

		fmt.Fprintf(w, `{
  "id": 1,
  "node_id": "MDEyOklzc3VlQ29tbWVudDE=",
  "url": "https://api.github.com/repos/diginfra/diginfra/issues/comments/1",
  "html_url": "https://github.com/diginfra/diginfra/issues/1347#issuecomment-1",
  "body": "Me too",
  "user": {
    "login": "diginfra",
    "id": 1,
    "node_id": "MDQ6VXNlcjE=",
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/diginfra",
    "html_url": "https://github.com/diginfra",
    "followers_url": "https://api.github.com/users/diginfra/followers",
    "following_url": "https://api.github.com/users/diginfra/following{/other_user}",
    "gists_url": "https://api.github.com/users/diginfra/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/diginfra/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/diginfra/subscriptions",
    "organizations_url": "https://api.github.com/users/diginfra/orgs",
    "repos_url": "https://api.github.com/users/diginfra/repos",
    "events_url": "https://api.github.com/users/diginfra/events{/privacy}",
    "received_events_url": "https://api.github.com/users/diginfra/received_events",
    "type": "User",
    "site_admin": false
  },
  "created_at": "2011-04-14T16:00:49Z",
  "updated_at": "2011-04-14T16:00:49Z",
  "issue_url": "https://api.github.com/repos/diginfra/diginfra/issues/1347",
  "author_association": "COLLABORATOR"
}`)
	}))
	defer ts.Close()

	dir := testutil.CalcGoldenFileTestdataDirName()
	GoldenFileCommandTest(t, dir, []string{
		"comment",
		"github",
		"--github-api-url", ts.URL,
		"--github-token", "test-token",
		"--pull-request", "8",
		"--behavior", "new",
		"--path", path.Join("./testdata", dir, "diginfra.json"),
		"--repo", "diginfra/diginfra",
	}, nil)

	assert.Equal(t, 2, attempts%3)
}

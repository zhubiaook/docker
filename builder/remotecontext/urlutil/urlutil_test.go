package urlutil // import "github.com/zhubiaook/docker/builder/remotecontext/urlutil"

import "testing"

var (
	gitUrls = []string{
		"git://github.com/zhubiaook/docker",
		"git@github.com:docker/docker.git",
		"git@bitbucket.org:atlassianlabs/atlassian-docker.git",
		"https://github.com/zhubiaook/docker.git",
		"http://github.com/zhubiaook/docker.git",
		"http://github.com/zhubiaook/docker.git#branch",
		"http://github.com/zhubiaook/docker.git#:dir",
	}
	incompleteGitUrls = []string{
		"github.com/zhubiaook/docker",
	}
	invalidGitUrls = []string{
		"http://github.com/zhubiaook/docker.git:#branch",
	}
)

func TestIsGIT(t *testing.T) {
	for _, url := range gitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range incompleteGitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range invalidGitUrls {
		if IsGitURL(url) {
			t.Fatalf("%q should not be detected as valid Git prefix", url)
		}
	}
}

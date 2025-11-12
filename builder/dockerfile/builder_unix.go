//go:build !windows
// +build !windows

package dockerfile // import "github.com/zhubiaook/docker/builder/dockerfile"

func defaultShellForOS(os string) []string {
	return []string{"/bin/sh", "-c"}
}

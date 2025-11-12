//go:build !windows
// +build !windows

package archive // import "github.com/zhubiaook/docker/pkg/archive"

import (
	"path/filepath"
)

func normalizePath(path string) string {
	return filepath.ToSlash(path)
}

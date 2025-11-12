//go:build !linux
// +build !linux

package vfs // import "github.com/zhubiaook/docker/daemon/graphdriver/vfs"

import (
	"github.com/zhubiaook/docker/pkg/chrootarchive"
	"github.com/zhubiaook/docker/pkg/idtools"
)

func dirCopy(srcDir, dstDir string) error {
	return chrootarchive.NewArchiver(idtools.IdentityMapping{}).CopyWithTar(srcDir, dstDir)
}

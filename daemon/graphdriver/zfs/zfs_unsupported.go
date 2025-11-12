//go:build !linux && !freebsd
// +build !linux,!freebsd

package zfs // import "github.com/zhubiaook/docker/daemon/graphdriver/zfs"

func checkRootdirFs(rootdir string) error {
	return nil
}

func getMountpoint(id string) string {
	return id
}

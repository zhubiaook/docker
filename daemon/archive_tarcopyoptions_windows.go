package daemon // import "github.com/zhubiaook/docker/daemon"

import (
	"github.com/zhubiaook/docker/container"
	"github.com/zhubiaook/docker/pkg/archive"
)

func (daemon *Daemon) tarCopyOptions(container *container.Container, noOverwriteDirNonDir bool) (*archive.TarOptions, error) {
	return daemon.defaultTarCopyOptions(noOverwriteDirNonDir), nil
}

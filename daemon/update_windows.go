package daemon // import "github.com/zhubiaook/docker/daemon"

import (
	"github.com/zhubiaook/docker/api/types/container"
	libcontainerdtypes "github.com/zhubiaook/docker/libcontainerd/types"
)

func toContainerdResources(resources container.Resources) *libcontainerdtypes.Resources {
	// We don't support update, so do nothing
	return nil
}

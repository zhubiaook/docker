package libcontainerd // import "github.com/zhubiaook/docker/libcontainerd"

import (
	"context"

	"github.com/containerd/containerd"
	"github.com/zhubiaook/docker/libcontainerd/local"
	"github.com/zhubiaook/docker/libcontainerd/remote"
	libcontainerdtypes "github.com/zhubiaook/docker/libcontainerd/types"
	"github.com/zhubiaook/docker/pkg/system"
)

// NewClient creates a new libcontainerd client from a containerd client
func NewClient(ctx context.Context, cli *containerd.Client, stateDir, ns string, b libcontainerdtypes.Backend) (libcontainerdtypes.Client, error) {
	if !system.ContainerdRuntimeSupported() {
		return local.NewClient(ctx, cli, stateDir, ns, b)
	}
	return remote.NewClient(ctx, cli, stateDir, ns, b)
}

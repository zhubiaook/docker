//go:build !exclude_graphdriver_fuseoverlayfs && linux
// +build !exclude_graphdriver_fuseoverlayfs,linux

package register // import "github.com/zhubiaook/docker/daemon/graphdriver/register"

import (
	// register the fuse-overlayfs graphdriver
	_ "github.com/zhubiaook/docker/daemon/graphdriver/fuse-overlayfs"
)

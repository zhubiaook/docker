//go:build !exclude_graphdriver_btrfs && linux
// +build !exclude_graphdriver_btrfs,linux

package register // import "github.com/zhubiaook/docker/daemon/graphdriver/register"

import (
	// register the btrfs graphdriver
	_ "github.com/zhubiaook/docker/daemon/graphdriver/btrfs"
)

//go:build (!exclude_graphdriver_zfs && linux) || (!exclude_graphdriver_zfs && freebsd)
// +build !exclude_graphdriver_zfs,linux !exclude_graphdriver_zfs,freebsd

package register // import "github.com/zhubiaook/docker/daemon/graphdriver/register"

import (
	// register the zfs driver
	_ "github.com/zhubiaook/docker/daemon/graphdriver/zfs"
)

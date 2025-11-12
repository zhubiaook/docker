//go:build !exclude_graphdriver_devicemapper && !static_build && linux
// +build !exclude_graphdriver_devicemapper,!static_build,linux

package register // import "github.com/zhubiaook/docker/daemon/graphdriver/register"

import (
	// register the devmapper graphdriver
	_ "github.com/zhubiaook/docker/daemon/graphdriver/devmapper"
)

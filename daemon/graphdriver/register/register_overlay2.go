//go:build !exclude_graphdriver_overlay2 && linux
// +build !exclude_graphdriver_overlay2,linux

package register // import "github.com/zhubiaook/docker/daemon/graphdriver/register"

import (
	// register the overlay2 graphdriver
	_ "github.com/zhubiaook/docker/daemon/graphdriver/overlay2"
)

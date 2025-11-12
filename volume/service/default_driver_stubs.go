//go:build !linux && !windows
// +build !linux,!windows

package service // import "github.com/zhubiaook/docker/volume/service"

import (
	"github.com/zhubiaook/docker/pkg/idtools"
	"github.com/zhubiaook/docker/volume/drivers"
)

func setupDefaultDriver(_ *drivers.Store, _ string, _ idtools.Identity) error { return nil }

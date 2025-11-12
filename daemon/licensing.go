package daemon // import "github.com/zhubiaook/docker/daemon"

import (
	"github.com/zhubiaook/docker/api/types"
	"github.com/zhubiaook/docker/dockerversion"
)

func (daemon *Daemon) fillLicense(v *types.Info) {
	v.ProductLicense = dockerversion.DefaultProductLicense
}

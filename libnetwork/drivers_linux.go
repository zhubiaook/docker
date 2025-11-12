package libnetwork

import (
	"github.com/zhubiaook/docker/libnetwork/drivers/bridge"
	"github.com/zhubiaook/docker/libnetwork/drivers/host"
	"github.com/zhubiaook/docker/libnetwork/drivers/ipvlan"
	"github.com/zhubiaook/docker/libnetwork/drivers/macvlan"
	"github.com/zhubiaook/docker/libnetwork/drivers/null"
	"github.com/zhubiaook/docker/libnetwork/drivers/overlay"
)

func getInitializers() []initializer {
	in := []initializer{
		{bridge.Register, "bridge"},
		{host.Register, "host"},
		{ipvlan.Register, "ipvlan"},
		{macvlan.Register, "macvlan"},
		{null.Register, "null"},
		{overlay.Register, "overlay"},
	}
	return in
}

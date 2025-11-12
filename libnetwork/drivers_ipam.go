package libnetwork

import (
	"github.com/zhubiaook/docker/libnetwork/ipamapi"
	builtinIpam "github.com/zhubiaook/docker/libnetwork/ipams/builtin"
	nullIpam "github.com/zhubiaook/docker/libnetwork/ipams/null"
	remoteIpam "github.com/zhubiaook/docker/libnetwork/ipams/remote"
	"github.com/zhubiaook/docker/libnetwork/ipamutils"
	"github.com/zhubiaook/docker/pkg/plugingetter"
)

func initIPAMDrivers(r ipamapi.Registerer, pg plugingetter.PluginGetter, addressPool []*ipamutils.NetworkToSplit) error {
	// TODO: pass address pools as arguments to builtinIpam.Init instead of
	// indirectly through global mutable state. Swarmkit references that
	// function so changing its signature breaks the build.
	if err := builtinIpam.SetDefaultIPAddressPool(addressPool); err != nil {
		return err
	}

	for _, fn := range [](func(ipamapi.Registerer) error){
		builtinIpam.Register,
		nullIpam.Register,
	} {
		if err := fn(r); err != nil {
			return err
		}
	}

	return remoteIpam.Register(r, pg)
}

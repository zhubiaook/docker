package cnmallocator

import (
	"github.com/zhubiaook/docker/libnetwork/drivers/overlay/ovmanager"
	"github.com/zhubiaook/docker/libnetwork/drivers/remote"
	"github.com/moby/swarmkit/v2/manager/allocator/networkallocator"
)

var initializers = []initializer{
	{remote.Init, "remote"},
	{ovmanager.Init, "overlay"},
	{StubManagerInit("internal"), "internal"},
	{StubManagerInit("l2bridge"), "l2bridge"},
	{StubManagerInit("nat"), "nat"},
}

// PredefinedNetworks returns the list of predefined network structures
func PredefinedNetworks() []networkallocator.PredefinedNetworkData {
	return []networkallocator.PredefinedNetworkData{
		{Name: "nat", Driver: "nat"},
	}
}

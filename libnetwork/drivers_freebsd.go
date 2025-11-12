package libnetwork

import (
	"github.com/zhubiaook/docker/libnetwork/drivers/null"
)

func getInitializers() []initializer {
	return []initializer{
		{null.Register, "null"},
	}
}

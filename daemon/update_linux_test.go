package daemon // import "github.com/zhubiaook/docker/daemon"

import (
	"testing"

	"github.com/zhubiaook/docker/api/types/container"
)

func TestToContainerdResources_Defaults(t *testing.T) {
	checkResourcesAreUnset(t, toContainerdResources(container.Resources{}))
}

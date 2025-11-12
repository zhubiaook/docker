package daemon // import "github.com/zhubiaook/docker/daemon"

import (
	"testing"

	"github.com/zhubiaook/docker/api/types"
	"github.com/zhubiaook/docker/dockerversion"
	"gotest.tools/v3/assert"
)

func TestFillLicense(t *testing.T) {
	v := &types.Info{}
	d := &Daemon{
		root: "/var/lib/docker/",
	}
	d.fillLicense(v)
	assert.Assert(t, v.ProductLicense == dockerversion.DefaultProductLicense)
}

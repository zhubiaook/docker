package images

import (
	"context"
	"errors"

	"github.com/zhubiaook/docker/container"
	"github.com/zhubiaook/docker/pkg/archive"
)

func (i *ImageService) Changes(ctx context.Context, container *container.Container) ([]archive.Change, error) {
	container.Lock()
	defer container.Unlock()

	if container.RWLayer == nil {
		return nil, errors.New("RWLayer of container " + container.Name + " is unexpectedly nil")
	}
	return container.RWLayer.Changes()
}

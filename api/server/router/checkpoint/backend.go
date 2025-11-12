package checkpoint // import "github.com/zhubiaook/docker/api/server/router/checkpoint"

import "github.com/zhubiaook/docker/api/types"

// Backend for Checkpoint
type Backend interface {
	CheckpointCreate(container string, config types.CheckpointCreateOptions) error
	CheckpointDelete(container string, config types.CheckpointDeleteOptions) error
	CheckpointList(container string, config types.CheckpointListOptions) ([]types.Checkpoint, error)
}

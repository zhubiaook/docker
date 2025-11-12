package volume // import "github.com/zhubiaook/docker/api/types/volume"

import "github.com/zhubiaook/docker/api/types/filters"

// ListOptions holds parameters to list volumes.
type ListOptions struct {
	Filters filters.Args
}

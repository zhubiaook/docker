package daemon // import "github.com/zhubiaook/docker/daemon"

import (
	"context"

	"github.com/zhubiaook/docker/api/types/registry"
	"github.com/zhubiaook/docker/dockerversion"
)

// AuthenticateToRegistry checks the validity of credentials in authConfig
func (daemon *Daemon) AuthenticateToRegistry(ctx context.Context, authConfig *registry.AuthConfig) (string, string, error) {
	return daemon.registryService.Auth(ctx, authConfig, dockerversion.DockerUserAgent(ctx))
}

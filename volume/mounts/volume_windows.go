package mounts // import "github.com/zhubiaook/docker/volume/mounts"

func (p *linuxParser) HasResource(m *MountPoint, absolutePath string) bool {
	return false
}

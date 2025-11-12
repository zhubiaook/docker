//go:build !linux
// +build !linux

package daemon // import "github.com/zhubiaook/docker/daemon"

func ensureDefaultAppArmorProfile() error {
	return nil
}

// DefaultApparmorProfile returns an empty string.
func DefaultApparmorProfile() string {
	return ""
}

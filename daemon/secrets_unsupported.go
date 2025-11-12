//go:build !linux && !windows
// +build !linux,!windows

package daemon // import "github.com/zhubiaook/docker/daemon"

func secretsSupported() bool {
	return false
}

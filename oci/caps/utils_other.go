//go:build !linux
// +build !linux

package caps // import "github.com/zhubiaook/docker/oci/caps"

func initCaps() {
	// no capabilities on Windows
}

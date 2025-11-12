//go:build !linux
// +build !linux

package archive // import "github.com/zhubiaook/docker/pkg/archive"

func getWhiteoutConverter(format WhiteoutFormat, inUserNS bool) (tarWhiteoutConverter, error) {
	return nil, nil
}

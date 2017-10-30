package cdn

import (
	"fmt"
)

func RefreshCdnUrl(urls ...string) ([]byte, error) {
	var options []string
	for k, v := range urls {
		options = append(options, fmt.Sprintf("urls.%v=%v", k, v))
	}
	return DoAction("RefreshCdnUrl", options...)
}

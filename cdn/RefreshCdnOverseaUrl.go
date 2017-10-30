package cdn

import (
	"fmt"
)

func RefreshCdnOverSeaUrl(urls ...string) ([]byte, error) {
	var options []string
	for k, v := range urls {
		options = append(options, fmt.Sprintf("urls.%v=%v", k, v))
	}
	return DoAction("RefreshCdnOverSeaUrl", options...)
}

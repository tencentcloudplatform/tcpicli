package cdn

import (
	"fmt"
)

func RefreshCdnOverSeaDir(dirs ...string) ([]byte, error) {
	var options []string
	for k, v := range dirs {
		options = append(options, fmt.Sprintf("dirs.%v=%v", k, v))
	}
	return DoAction("RefreshCdnOverSeaDir", options...)
}

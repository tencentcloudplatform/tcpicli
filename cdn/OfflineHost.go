package cdn

import (
	"fmt"
	"strconv"
)

func OfflineHost(hostOrId string) ([]byte, error) {
	if _, err := strconv.Atoi(hostOrId); err == nil {
		return DoAction("OfflineHost", fmt.Sprintf("hostId=%v", hostOrId))
	}
	return DoAction("OfflineHost", fmt.Sprintf("host=%v", hostOrId))
}

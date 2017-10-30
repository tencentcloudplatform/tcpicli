package cdn

import (
	"fmt"
	"strconv"
)

func OnlineHost(hostOrId string) ([]byte, error) {
	if _, err := strconv.Atoi(hostOrId); err == nil {
		return DoAction("OnlineHost", fmt.Sprintf("hostId=%v", hostOrId))
	}
	return DoAction("OnlineHost", fmt.Sprintf("host=%v", hostOrId))
}

package cdn

import (
	"fmt"
	"strconv"
)

func DeleteCdnHost(hostOrId string) ([]byte, error) {
	if _, err := strconv.Atoi(hostOrId); err == nil {
		return DoAction("DeleteCdnHost", fmt.Sprintf("hostId=%v", hostOrId))
	}
	return DoAction("DeleteCdnHost", fmt.Sprintf("host=%v", hostOrId))
}

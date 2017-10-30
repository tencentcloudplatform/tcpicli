package cdn

import ()

func GetCdnRefreshLog(options ...string) ([]byte, error) {
	return DoAction("GetCdnRefreshLog", options...)
}

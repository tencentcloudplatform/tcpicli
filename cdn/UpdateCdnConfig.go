package cdn

import ()

func UpdateCdnConfig(options ...string) ([]byte, error) {
	return DoAction("UpdateCdnConfig", options...)
}

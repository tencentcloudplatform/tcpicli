package cdn

import ()

func UpdateCache(options ...string) ([]byte, error) {
	return DoAction("UpdateCache", options...)
}

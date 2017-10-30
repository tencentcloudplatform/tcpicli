package cdn

import ()

func UpdateCdnProject(options ...string) ([]byte, error) {
	return DoAction("UpdateCdnProject", options...)
}

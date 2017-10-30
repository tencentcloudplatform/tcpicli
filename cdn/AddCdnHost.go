package cdn

import ()

func AddCdnHost(options ...string) ([]byte, error) {
	return DoAction("AddCdnHost", options...)
}

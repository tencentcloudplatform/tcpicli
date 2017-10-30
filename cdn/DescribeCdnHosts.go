package cdn

import ()

func DescribeCdnHosts(options ...string) ([]byte, error) {
	return DoAction("DescribeCdnHosts", options...)
}

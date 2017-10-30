package cdn

import (
	"bytes"
)

func GetCdnLogList(options ...string) ([]byte, error) {
	b, err := DoAction("GetCdnLogList", options...)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	return b, err
}

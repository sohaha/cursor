package cursor

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/sohaha/zlsgo/zstring"
)

func Parse(data []byte) (string, error) {
	var begin bool
	buffer := &bytes.Buffer{}
	arr := bytes.Split(data, []byte("data: "))
	for i := range arr {
		if bytes.Contains(arr[i], []byte("<|BEGIN_message|>")) {
			begin = true
			continue
		}
		if bytes.Contains(arr[i], []byte("<|END_message|>")) {
			break
		}
		if begin {
			p, err := zstring.RegexExtractAll(`"(.+)"`, zstring.Bytes2String(arr[i]), -1)
			if err != nil {
				return "", err
			}
			if len(p) == 0 {
				return "", errors.New("parse error")
			}
			code, _ := strconv.Unquote(p[0][0])
			buffer.WriteString(code)
		}
	}
	return zstring.Bytes2String(buffer.Bytes()), nil
}

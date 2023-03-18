package cursor

import (
	"github.com/sohaha/zlsgo/zhttp"
)

var request = zhttp.New()

func NewRequest(method string, url string, payload []byte) (res *zhttp.Res, err error) {
	headers := zhttp.Header{
		"Content-Type":    "application/json",
		"Accept":          "*/*",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Cursor/0.1.1 Chrome/108.0.5359.62 Electron/22.0.0 Safari/537.36",
		"Sec-Fetch-Site":  "cross-site",
		"Sec-Fetch-Mode":  "cors",
		"Sec-Fetch-Dest":  "empty",
		"Accept-Encoding": "gzip, deflate, br",
	}

	return request.Do(method, url, headers, zhttp.BodyJSON(payload))
}

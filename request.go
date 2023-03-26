package cursor

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/zhttp"
	"github.com/sohaha/zlsgo/zlog"
)

var Timeout = time.Second * 10
var request = zhttp.New()

func NewRequest(method string, url string, payload []byte) (res *zhttp.Res, err error) {

	values := []interface{}{
		zhttp.Header{
			"Content-Type":    "application/json",
			"Accept":          "*/*",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Cursor/0.1.1 Chrome/108.0.5359.62 Electron/22.0.0 Safari/537.36",
			"Sec-Fetch-Site":  "cross-site",
			"Sec-Fetch-Mode":  "cors",
			"Sec-Fetch-Dest":  "empty",
			"Accept-Encoding": "gzip, deflate, br",
		},
		zhttp.BodyJSON(payload),
	}

	if Timeout > 0 {
		ctx, cancel := context.WithTimeout(context.Background(), Timeout)
		defer cancel()
		values = append(values, ctx)
	}

	res, err = request.Do(method, url, values...)

	if err != nil {
		zlog.Error(err)
		if strings.Contains(err.Error(), "context deadline exceeded") {
			err = errors.New("timeout")
		}
		return
	}

	if res.StatusCode() == 504 {
		err = errors.New("maximum capacity reached")
		return
	}

	return
}

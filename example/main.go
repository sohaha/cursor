package main

import (
	"github.com/sohaha/cursor"
	"github.com/sohaha/zlsgo/zlog"
)

func main() {
	zlog.ResetFlags(zlog.BitLevel)

	res, err := cursor.Conv("get the current time zone")
	if err != nil {
		zlog.Error(err)
		return
	}
	zlog.Success("\n" + res)
}

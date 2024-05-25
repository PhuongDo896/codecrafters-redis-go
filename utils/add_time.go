package utils

import (
	"sync/atomic"
	"time"
)

func AddTime(expireTime int64) int64 {
	now := time.Now().UnixMilli()
	return atomic.AddInt64(&now, expireTime)
}

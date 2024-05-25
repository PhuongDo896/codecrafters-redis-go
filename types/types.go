package types

import (
	"sync"
	"time"

	"github.com/codecrafters-io/redis-starter-go/utils"
)

type GlobalMap struct {
	Mu    sync.Mutex
	Store map[string]TValue
}

type TValue struct {
	Value        string
	ExpiredTime  int64
	ExpireOption bool
}

func (g *GlobalMap) NSet(key, value string) {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	g.Store[key] = TValue{Value: value, ExpireOption: false}
}

func (g *GlobalMap) ESet(key, value string, expireTime int64) {
	defer g.Mu.Unlock()
	g.Mu.Lock()
	g.Store[key] = TValue{
		Value:        value,
		ExpiredTime:  utils.AddTime(expireTime),
		ExpireOption: true,
	}
}

func (g *GlobalMap) Get(key string) string {
	val, ok := g.Store[key]
	if !ok {
		return ""
	}

	if !val.ExpireOption {
		return val.Value
	} else {
		if time.Now().UnixMilli() < val.ExpiredTime {
			return val.Value
		} else {
			delete(g.Store, key)
			return "$-1\r\n"
		}
	}
}

// only delete if expire option is true
func (g *GlobalMap) Del(key string) {
	val, ok := g.Store[key]

	if !ok {
		return
	}

	if !val.ExpireOption {
		return
	}

	delete(g.Store, key)
}

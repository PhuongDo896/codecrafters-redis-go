package types

import "sync"

type GlobalMap struct {
	Mu    sync.Mutex
	Store map[string]string
}

func (g *GlobalMap) Set(key, value string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Store[key] = value
}

func (g *GlobalMap) Get(key string) string {
	return g.Store[key]
}

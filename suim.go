package suim

import "sync"

var (
	formConfigs = map[string]*FormConfig{}
	gridConfigs = map[string]*GridConfig{}

	mtx = new(sync.RWMutex)
)

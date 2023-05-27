package wslogic

import "sync"

var (
	userCount uint64
	rwLock    *sync.RWMutex
)

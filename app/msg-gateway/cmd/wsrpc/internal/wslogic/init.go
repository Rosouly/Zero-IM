package wslogic

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

var (
	userCount       uint64
	rwLock          *sync.RWMutex
	validate        *validator.Validate
	sendMsgAllCount uint64
)

func init() {
	rwLock = new(sync.RWMutex)
	validate = validator.New()
}

package packer

import (
	"os"
	"sync"
)

var empty Manager
var defaultMgr Manager
var defaultErr error
var once sync.Once
var (
	Stdout = os.Stdout
	Stderr = os.Stderr
)

func Default() (mgr Manager, err error) {
	once.Do(func() {
		defaultMgr, defaultErr = DetectManager()
	})
	mgr = defaultMgr
	err = defaultErr
	return
}

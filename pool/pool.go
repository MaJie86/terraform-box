package pool

import (
	"github.com/goinggo/workpool"
	"runtime"
)

var pool *workpool.WorkPool

func Init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pool = workpool.New(runtime.NumCPU(), 800)
}

func Run(goRoutine string, work workpool.PoolWorker) {
	pool.PostWork(goRoutine, work)
}

func Cancel(goRoutine string, work workpool.PoolWorker) {
	pool.Shutdown(goRoutine)
}

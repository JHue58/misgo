package runner

import (
	"github.com/jhue/misgo/internal/mislog"
	"github.com/panjf2000/ants/v2"
)

var pool, _ = ants.NewPool(-1, ants.WithPanicHandler(func(i interface{}) {
	mislog.DefaultLogger.Errorf("panic in Go: %v", i)
}))

func Go(f func()) {
	_ = pool.Submit(f)
}

package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jhue/misgo/internal/mislog"
	"github.com/jhue58/latency/recorder"
	"time"

	"github.com/jhue58/latency"
	"github.com/jhue58/latency/buckets"
	"github.com/jhue58/latency/duration"
	"strings"
	"sync"
)

type PprofKey struct {
}

type Pprof struct {
	mp     sync.Map
	format latency.FormatFunc
}

func NewPProf() *Pprof {
	return &Pprof{format: latency.DefaultFormat}
}

func (m *Pprof) Report() (str string) {

	dict := make(map[string]recorder.RecordedSnapshot)
	m.mp.Range(func(key, value any) bool {
		method, ok := key.(string)
		if !ok {
			mislog.DefaultLogger.Errorf("pprof中key string断言失败")
			return ok
		}

		r, ok := value.(recorder.Recorder)
		if !ok {
			mislog.DefaultLogger.Errorf("pprof中reporter断言失败")
			return ok
		}
		sp := recorder.RecordedSnapshot{}
		r.Snapshot(sp)
		dict[method] = sp
		return true
	})

	var builder strings.Builder
	for s, snapshot := range dict {
		builder.WriteString(s)
		builder.WriteString(" :\n")
		builder.WriteString(m.format(snapshot))
		builder.WriteByte('\n')
	}
	return builder.String()

}

func (m *Pprof) Record(c *app.RequestContext, d duration.Duration) {
	var builder strings.Builder
	builder.Write(c.Request.Path())
	builder.WriteByte(' ')
	builder.Write(c.Method())
	method := builder.String()

	r, ok := m.mp.Load(method)
	if !ok {
		r, _ = m.mp.LoadOrStore(method, newRecorder())
	}
	b, ok := r.(recorder.Recorder)
	if !ok {
		mislog.DefaultLogger.Errorf("pprof中reporter断言失败")
		return
	}
	b.Record(d)

}

func newRecorder() recorder.Recorder {
	bucket := buckets.NewBucketsRecorder()
	return bucket
}

func PprofMiddleware() func(ctx context.Context, c *app.RequestContext) {
	p := NewPProf()

	return func(ctx context.Context, c *app.RequestContext) {
		ctx = context.WithValue(ctx, PprofKey{}, p)
		start := time.Now()
		c.Next(ctx)
		d := duration.NewDuration(time.Since(start))
		p.Record(c, d)
	}
}

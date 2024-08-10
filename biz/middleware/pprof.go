package middleware

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jhue/misgo/internal/mislog"
	"github.com/jhue/misgo/pkg/formater"
	"github.com/jhue58/latency/recorder"
	"time"

	"github.com/jhue58/latency"
	"github.com/jhue58/latency/buckets"
	"github.com/jhue58/latency/duration"
	"strings"
	"sync"
)

var apiSlice = []byte("/api/")
var otherSlice = []byte("other")

type PprofKey struct {
}

type Pprof struct {
	mp     sync.Map
	format latency.FormatFunc
}

func NewPProf() *Pprof {
	return &Pprof{format: latency.DefaultFormat}
}

func (m *Pprof) Report() string {
	var builder formater.Builder
	m.ReportWithBuilder(&builder)
	return builder.String()
}

func (m *Pprof) ReportWithBuilder(builder *formater.Builder) {

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

	for s, snapshot := range dict {
		builder.WriteStringItem(s, m.format(snapshot))
	}
	return
}

func (m *Pprof) Record(c *app.RequestContext, d duration.Duration) {
	var builder strings.Builder
	path := c.Request.Path()
	if !bytes.Contains(path, apiSlice) {
		path = otherSlice
	}

	builder.Write(path)
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

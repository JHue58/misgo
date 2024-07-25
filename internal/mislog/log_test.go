package mislog

import (
	"github.com/jhue/misgo/internal/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLog(t *testing.T) {
	i, err := LogFileIO(".")
	assert.NoError(t, err)
	InitLogger(i, conf.LogConfig{Level: "debug"})

	for i := 0; i < 100; i++ {
		DefaultLogger.Error("hello")
	}
}

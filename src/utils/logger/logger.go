package logger

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"os"
)

func BuildLogger(service string) (logger log.Logger) {
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "time", log.DefaultTimestampUTC)
	logger = log.With(logger, "service", service)
	return
}
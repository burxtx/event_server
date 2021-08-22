package log

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Format      string        `mapstructure:"format"`
	Path        string        `mapstructure:"path"`
	Output      string        `mapstructure:"output"`
	Level       string        `mapstructure:"level"`
	MaxAgeDay   time.Duration `mapstructure:"max_age_day"`
	RotationDay time.Duration `mapstructure:"rotation_day"`
}

type MyLogger struct {
	*logrus.Logger
	Cfg          Config
	RequestIDKey string
}

func NewLogger(cfg Config) (MyLogger, error) {
	log := logrus.New()
	// set output
	switch cfg.Output {
	case "stdout":
		log.Out = os.Stdout
	case "stderr":
		log.Out = os.Stderr
	default:
		f, err := os.OpenFile(cfg.Output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)
		if err != nil {
			return MyLogger{}, err
		}
		log.Out = f
	}

	// set level
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		panic(fmt.Sprintf("设置日志级别失败, err: %s", err.Error()))
	}
	log.SetLevel(level)

	// set format
	switch cfg.Format {
	case "txt":
		log.SetFormatter(&logrus.TextFormatter{})
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	default:
		log.SetFormatter(&logrus.JSONFormatter{})
	}

	my := MyLogger{
		Logger: log,
		Cfg:    cfg,
	}
	return my, nil
}

// addContext is used when context need to be added to logger
func (l *MyLogger) addContext(ctx context.Context, keys ...interface{}) *logrus.Entry {
	fields := logrus.Fields{}
	for _, key := range keys {
		v := ctx.Value(key)
		if v != nil {
			fields[fmt.Sprint(key)] = v
		}
	}
	// add default field: request_id
	// fields[l.RequestIDKey] = middleware.GetReqID(ctx)

	return l.WithFields(fields)

}

func (l *MyLogger) GetLogger(ctx context.Context, keys ...interface{}) *logrus.Entry {
	return l.addContext(ctx, keys...)
}

func (l *MyLogger) GetLoggerWithFields(ctx context.Context, fields map[string]interface{}, keys ...interface{}) *logrus.Entry {
	return l.addContext(ctx, keys...).WithFields(logrus.Fields(fields))
}

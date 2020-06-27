package log

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(s *builder) error

type builder struct {
	options []zap.Option
	writers []zapcore.WriteSyncer
	config  zapcore.EncoderConfig
	encoder zapcore.Encoder
	level   zapcore.LevelEnabler
}

func newBuilder() *builder {
	return &builder{
		writers: []zapcore.WriteSyncer{os.Stdout},
		options: make([]zap.Option, 0),
		config: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "name",
			CallerKey:      "caller",
			StacktraceKey:  "stack",
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
	}
}

func (b *builder) init() {
	if b.encoder == nil {
		b.encoder = zapcore.NewConsoleEncoder(b.config)
	}

	if b.level == nil {
		b.level = zap.InfoLevel
	}
}

func AddCaller() Option {
	return func(s *builder) (err error) {
		s.options = append(s.options, zap.AddCaller())
		return
	}
}

func AddStacktrace(lvl string) Option {
	return func(s *builder) (err error) {
		s.options = append(s.options, zap.AddStacktrace(parseLevel(lvl)))
		return
	}
}

func SetLevel(lvl string) Option {
	return func(s *builder) (err error) {
		s.level = parseLevel(lvl)
		s.options = append(s.options, zap.IncreaseLevel(s.level))

		return
	}
}

func WithFields(fields map[string]interface{}) Option {
	return func(s *builder) (err error) {
		zapFields := make([]zap.Field, 0, len(fields))

		for key, val := range fields {
			zapFields = append(zapFields, zap.Any(key, val))
		}

		s.options = append(s.options, zap.Fields(zapFields...))

		return
	}
}

func EnableJSON() Option {
	return func(s *builder) (err error) {
		s.encoder = zapcore.NewJSONEncoder(s.config)
		s.config.EncodeLevel = zapcore.CapitalLevelEncoder

		return
	}
}

func NewLogger(options ...Option) (*zap.SugaredLogger, error) {
	build := newBuilder()

	for _, option := range options {
		if err := option(build); err != nil {
			return nil, err
		}
	}

	build.init()

	core := zapcore.NewCore(
		build.encoder,
		zapcore.NewMultiWriteSyncer(build.writers...),
		build.level,
	)

	return zap.New(core, build.options...).Sugar(), nil
}

func parseLevel(lvl string) zapcore.Level {
	switch strings.ToLower(lvl) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn", "warning":
		return zap.WarnLevel
	case "err", "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

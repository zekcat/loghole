package interfaces

import (
	"context"

	"github.com/lissteron/loghole/collector/internal/app/domain"
)

type Logger interface {
	Debug(ctx context.Context, args ...interface{})
	Debugf(ctx context.Context, template string, args ...interface{})
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, template string, args ...interface{})
	Warn(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, template string, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, template string, args ...interface{})
}

type Storage interface {
	StoreEntryList(ctx context.Context, list []*domain.Entry) (err error)
}

package usecases

import (
	"context"

	"github.com/gadavy/tracing"

	"github.com/lissteron/loghole/collector/internal/app/domain"
	"github.com/lissteron/loghole/collector/internal/app/usecases/interfaces"
)

type StoreEntryList struct {
	storage interfaces.Storage
	logger  interfaces.Logger
}

func (s *StoreEntryList) Do(ctx context.Context, data []byte) (err error) {
	defer tracing.ChildSpan(&ctx).Finish()

	list, err := s.parseEntryList(ctx, data[:])
	if err != nil {
		return err
	}

	err = s.storage.StoreEntryList(ctx, list)
	if err != nil {
		return err
	}

	return nil
}

func (s *StoreEntryList) parseEntryList(ctx context.Context, data []byte) ([]*domain.Entry, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	return nil, nil
}

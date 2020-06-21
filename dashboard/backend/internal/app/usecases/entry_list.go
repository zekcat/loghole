package usecases

import (
	"context"
	"strings"

	"github.com/gadavy/tracing"
	"github.com/lissteron/simplerr"

	"github.com/lissteron/loghole/dashboard/internal/app/codes"
	"github.com/lissteron/loghole/dashboard/internal/app/domain"
	"github.com/lissteron/loghole/dashboard/internal/app/usecases/interfaces"
)

type ListEntryIn struct {
	*domain.Query
}

type ListEntry struct {
	storage interfaces.Storage
	logger  interfaces.Logger
}

func NewListEntry(
	storage interfaces.Storage,
	logger interfaces.Logger,
) *ListEntry {
	return &ListEntry{
		storage: storage,
		logger:  logger,
	}
}

func (l *ListEntry) Do(ctx context.Context, input *ListEntryIn) ([]*domain.Entry, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	params := l.prepareParams(ctx, input.Params)

	result, err := l.storage.ListEntry(ctx, params, input.Limit, input.Offset)
	if err != nil {
		l.logger.Errorf(ctx, "get entry list failed: %v", err)
		return nil, simplerr.WrapWithCode(err, codes.DatabaseError, "get entry list failed")
	}

	return result, nil
}

func (l *ListEntry) prepareParams(ctx context.Context, params []*domain.QueryParam) (prepared [][]*domain.QueryParam) {
	defer tracing.ChildSpan(&ctx).Finish()

	var counter = 1

	for idx, val := range params {
		if val.Join == domain.JoinOr && idx != 0 {
			counter++
		}
	}

	prepared = make([][]*domain.QueryParam, counter)

	counter = 0

	for _, val := range params {
		switch strings.ToUpper(val.Join) {
		case domain.JoinOr:
			counter++
		case "":
			val.Join = domain.JoinAnd
		}

		prepared[counter] = append(prepared[counter], val)
	}

	return prepared
}

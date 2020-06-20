package request

import (
	"encoding/json"
	"net/http"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
	"github.com/lissteron/loghole/dashboard/internal/app/usecases"
)

type ListEntryRequest struct {
	*domain.Query
}

func ReadListEntryRequest(r *http.Request) (*ListEntryRequest, error) {
	req := &ListEntryRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}

	if err := req.Validate(); err != nil {
		return nil, err
	}

	return req, nil
}

func (r *ListEntryRequest) Validate() error {
	return nil
}

func (r *ListEntryRequest) ToInput() *usecases.ListEntryIn {
	return &usecases.ListEntryIn{
		Query: r.Query,
	}
}

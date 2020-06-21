package request

import (
	"encoding/json"
	"net/http"
	"strings"

	validation "github.com/gadavy/ozzo-validation/v4"

	"github.com/lissteron/loghole/dashboard/internal/app/codes"
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
	if err := r.validateQueryParams(); err != nil {
		return err
	}

	return validation.ValidateStruct(r,
		validation.Field(&r.Limit, r.limitRules()...),
		validation.Field(&r.Offset, r.offsetRules()...),
	)
}

func (r *ListEntryRequest) ToInput() *usecases.ListEntryIn {
	return &usecases.ListEntryIn{
		Query: r.Query,
	}
}

func (r *ListEntryRequest) validateQueryParams() error {
	for _, param := range r.Params {
		param.Type = strings.ToLower(param.Type)
		param.Join = strings.ToUpper(param.Join)

		return validation.ValidateStruct(param,
			validation.Field(&param.Type, r.queryParamsTypeRules()...),
			validation.Field(&param.Join, r.queryParamsTypeRules()...),
			validation.Field(&param.Key, r.queryParamsKeyRules()...),
			validation.Field(&param.Value, r.queryParamsValueRules()...),
			validation.Field(&param.Operator, r.queryParamsOperatorRules()...),
		)
	}

	return nil
}

func (r *ListEntryRequest) limitRules() []validation.Rule {
	return []validation.Rule{
		validation.Min(int64(0)).ErrorCode(codes.ValidMinLimit.String()),
		validation.Max(int64(1000)).ErrorCode(codes.ValidMaxLimit.String()),
	}
}

func (r *ListEntryRequest) offsetRules() []validation.Rule {
	return []validation.Rule{
		validation.Min(int64(0)).ErrorCode(codes.ValidMixOffset.String()),
	}
}

func (r *ListEntryRequest) queryParamsTypeRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidQueryParamsTypeRequired.String()),
		validation.In(domain.TypeKey, domain.TypeColumn).ErrorCode(codes.ValidQueryParamsTypeIn.String()),
	}
}

func (r *ListEntryRequest) queryParamsJoinRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidQueryParamsJoinRequired.String()),
		validation.In(domain.JoinOr, domain.JoinAnd).ErrorCode(codes.ValidQueryParamsJoinIn.String()),
	}
}

func (r *ListEntryRequest) queryParamsKeyRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidQueryParamsKeyRequired.String()),
	}
}

func (r *ListEntryRequest) queryParamsValueRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidQueryParamsValueRequired.String()),
	}
}

func (r *ListEntryRequest) queryParamsOperatorRules() []validation.Rule {
	return []validation.Rule{
		validation.Required.ErrorCode(codes.ValidQueryParamsOperatorRequired.String()),
		validation.In("<", "<=", ">=", "<>", ">", "!=", "=", "LIKE").
			ErrorCode(codes.ValidQueryParamsOperatorIn.String()),
	}
}

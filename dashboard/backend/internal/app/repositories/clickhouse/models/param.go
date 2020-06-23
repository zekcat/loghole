package models

import (
	"errors"
	"strings"

	"github.com/Masterminds/squirrel"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
	"github.com/lissteron/loghole/dashboard/internal/app/repositories/clickhouse/tools"
)

const (
	jsonFieldString = "params_string"
	jsonFieldFloat  = "params_float"
)

const (
	stringParams = "(has(params_string.keys, '%s') AND params_string.values[indexOf(params_string.keys, '%s')] %s ?)"
	floatParams  = "(has(params_float.keys, '%s') AND params_float.values[indexOf(params_float.keys, '%s')] %s ?)"
)

type Param struct {
	domain.QueryParam
}

func ParamFromDomain(param *domain.QueryParam) *Param {
	return &Param{*param}
}

// nolint:golint,stylecheck,gocritic
func (p *Param) ToSql() (query string, args []interface{}, err error) {
	if p.IsTypeJSON() {
		return p.prepareJSON()
	}

	if p.Operator == domain.OperatorIn {
		return squirrel.Eq{p.Key: p.Value.List}.ToSql()
	}

	if p.Operator == domain.OperatorNotIn {
		return squirrel.NotEq{p.Key: p.Value.List}.ToSql()
	}

	if p.Operator == domain.OperatorLike {
		return p.prepareLikeParam()
	}

	if p.Operator == domain.OperatorNotLike {
		return p.prepareNotLikeParam()
	}

	return strings.Join([]string{p.Key, p.Operator, "?"}, ""), append(args, p.Value.Item), nil
}

func (p *Param) prepareLikeParam() (query string, args []interface{}, err error) {
	if len(p.Value.List) > 0 {
		builder := make(squirrel.Or, 0, len(p.Value.List))

		for _, value := range p.Value.List {
			builder = append(builder, squirrel.Like{p.Key: tools.CreateLike(value)})
		}

		return builder.ToSql()
	}

	return squirrel.Like{p.Key: tools.CreateLike(p.Value.Item)}.ToSql()
}

func (p *Param) prepareNotLikeParam() (query string, args []interface{}, err error) {
	if len(p.Value.List) > 0 {
		builder := make(squirrel.Or, 0, len(p.Value.List))

		for _, value := range p.Value.List {
			builder = append(builder, squirrel.NotLike{p.Key: tools.CreateLike(value)})
		}

		return builder.ToSql()
	}

	return squirrel.NotLike{p.Key: tools.CreateLike(p.Value.Item)}.ToSql()
}

func (p *Param) prepareJSON() (query string, args []interface{}, err error) {
	return "", nil, errors.New("unimplemented")
}

package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Masterminds/squirrel"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
	"github.com/lissteron/loghole/dashboard/internal/app/repositories/clickhouse/tools"
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
		return p.prepareParamLike()
	}

	if p.Operator == domain.OperatorNotLike {
		return p.prepareParamNotLike()
	}

	return strings.Join([]string{p.Key, p.Operator, "?"}, ""), append(args, p.Value.Item), nil
}

func (p *Param) prepareParamLike() (query string, args []interface{}, err error) {
	if len(p.Value.List) > 0 {
		builder := make(squirrel.And, 0, len(p.Value.List))

		for _, value := range p.Value.List {
			builder = append(builder, squirrel.Like{p.Key: tools.CreateLike(value)})
		}

		return builder.ToSql()
	}

	return squirrel.Like{p.Key: tools.CreateLike(p.Value.Item)}.ToSql()
}

func (p *Param) prepareParamNotLike() (query string, args []interface{}, err error) {
	if len(p.Value.List) > 0 {
		builder := make(squirrel.And, 0, len(p.Value.List))

		for _, value := range p.Value.List {
			builder = append(builder, squirrel.NotLike{p.Key: tools.CreateLike(value)})
		}

		return builder.ToSql()
	}

	return squirrel.NotLike{p.Key: tools.CreateLike(p.Value.Item)}.ToSql()
}

func (p *Param) prepareJSON() (query string, args []interface{}, err error) {
	if p.Operator == ">" || p.Operator == "<" {
		if len(p.Value.List) > 0 {
			return "", nil, errors.New("json params unimplemented")
		}

		if value, ok := valueToFloat(p.Value.Item); ok {
			return fmt.Sprintf(floatParams, p.Key, p.Key, p.Operator), append(args, value), nil
		}
	}

	if len(p.Value.List) > 0 {
		return "", nil, errors.New("json params unimplemented")
	}

	return fmt.Sprintf(stringParams, p.Key, p.Key, p.Operator), append(args, p.Value.Item), nil
}

func valueToFloat(val string) (float64, bool) {
	f, err := strconv.ParseFloat(val, 64)

	return f, err == nil
}

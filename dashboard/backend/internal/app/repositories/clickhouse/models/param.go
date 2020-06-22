package models

import (
	"errors"
	"strings"

	"github.com/Masterminds/squirrel"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
)

const (
	stringParams = "(has(params_string.keys, '%s') AND params_string.values[indexOf(params_string.keys, '%s')] %s ?)"
	floatParams  = "(has(params_float.keys, '%s') AND params_float.values[indexOf(params_float.keys, '%s')] %s ?)"
)

type Param struct {
	*domain.QueryParam
}

func ParamFromDomain(param *domain.QueryParam) *Param {
	return &Param{param}
}

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
		if len(p.Value.List) > 0 {
			return "", nil, errors.New("unimplemented")
		}

		return squirrel.Like{p.Key: p.Value.Item}.ToSql()
	}

	if p.Operator == domain.OperatorNotLike {
		if len(p.Value.List) > 0 {
			return "", nil, errors.New("unimplemented")
		}

		return squirrel.NotLike{p.Key: p.Value.Item}.ToSql()
	}

	return strings.Join([]string{p.Key, p.Operator, "?"}, ""), append(args, p.Value.Item), nil
}

func (p *Param) prepareJSON() (query string, args []interface{}, err error) {
	return "", nil, errors.New("unimplemented")
}
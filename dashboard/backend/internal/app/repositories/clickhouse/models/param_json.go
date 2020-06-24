package models

import (
	"fmt"
	"strconv"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
)

const (
	stringParams = "(has(params_string.keys, ?) AND params_string.values[indexOf(params_string.keys, ?)] %s ?)"
	floatParams  = "(has(params_float.keys, ?) AND params_float.values[indexOf(params_float.keys, ?)] %s ?)"
)

type JSONParam struct {
	domain.QueryParam
}

func JSONParamFromDomain(param *domain.QueryParam) *JSONParam {
	return &JSONParam{*param}
}

func (p *JSONParam) getIn() (query string, args []interface{}, err error) {
	return "", nil, ErrNotImplemented
}

func (p *JSONParam) getNotIn() (query string, args []interface{}, err error) {
	return "", nil, ErrNotImplemented
}

func (p *JSONParam) getLike() (query string, args []interface{}, err error) {
	return "", nil, ErrNotImplemented
}

func (p *JSONParam) getNotLike() (query string, args []interface{}, err error) {
	return "", nil, ErrNotImplemented
}

func (p *JSONParam) getLikeWithValue(val string) (query string, args []interface{}, err error) {
	return "", nil, ErrNotImplemented
}

func (p *JSONParam) getNotLikeWithValue(val string) (query string, args []interface{}, err error) {
	return "", nil, ErrNotImplemented
}

func (p *JSONParam) getLtGtString() (query string, args []interface{}, err error) {
	return fmt.Sprintf(stringParams, p.Operator), []interface{}{p.Key, p.Key, p.Value.Item}, nil
}

func (p *JSONParam) getLtGtFloat(val float64) (query string, args []interface{}, err error) {
	return fmt.Sprintf(floatParams, p.Operator), []interface{}{p.Key, p.Key, val}, nil
}

func (p *JSONParam) getDefault() (query string, args []interface{}, err error) {
	return fmt.Sprintf(stringParams, p.Operator), []interface{}{p.Key, p.Key, p.Value.Item}, nil
}

func valueToFloat(val string) (float64, bool) {
	f, err := strconv.ParseFloat(val, 64)

	return f, err == nil
}

/*
func (p *Param) prepareJSON() (query string, args []interface{}, err error) {
	if len(p.Value.List) > 0 {
		return p.prepareArrayJSON()
	}

	switch p.Operator {
	case ">", "<", ">=", "<=":
		if value, ok := valueToFloat(p.Value.Item); ok {
			return fmt.Sprintf(floatParams, p.Operator), []interface{}{p.Key, p.Key, value}, nil
		}

		return
	}

	return fmt.Sprintf(stringParams, p.Operator), []interface{}{p.Key, p.Key, p.Value.Item}, nil
}

func (p *Param) prepareArrayJSON() (query string, args []interface{}, err error) {
	switch p.Operator {
	case ">", "<":
		builder := make([]string, 0, len(p.Value.List))

		for _, listValue := range p.Value.List {
			value, ok := valueToFloat(listValue)
			if !ok {
				return "", nil, errors.New()
			}

			builder = append(builder, fmt.Sprintf(floatParams, p.Operator)),
				args = append(args, p.Key, p.Key, value)
		}
	}
}
*/

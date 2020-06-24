package models

import (
	"strings"

	"github.com/Masterminds/squirrel"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
	"github.com/lissteron/loghole/dashboard/internal/app/repositories/clickhouse/tools"
)

type Param struct {
	domain.QueryParam
}

func ParamFromDomain(param *domain.QueryParam) *Param {
	return &Param{*param}
}

// nolint:golint,stylecheck,gocritic
func (p *Param) ToSql() (query string, args []interface{}, err error) {
	return buildSQL(p)
}

func (p *Param) getIn() (query string, args []interface{}, err error) {
	return squirrel.Eq{p.Key: p.Value.List}.ToSql()
}

func (p *Param) getNotIn() (query string, args []interface{}, err error) {
	return squirrel.NotEq{p.Key: p.Value.List}.ToSql()
}

func (p *Param) getLike() (query string, args []interface{}, err error) {
	return squirrel.Like{p.Key: tools.CreateLike(p.Value.Item)}.ToSql()
}

func (p *Param) getNotLike() (query string, args []interface{}, err error) {
	return squirrel.NotLike{p.Key: tools.CreateLike(p.Value.Item)}.ToSql()
}

func (p *Param) getLikeWithValue(val string) (query string, args []interface{}, err error) {
	return squirrel.Like{p.Key: tools.CreateLike(val)}.ToSql()
}

func (p *Param) getNotLikeWithValue(val string) (query string, args []interface{}, err error) {
	return squirrel.NotLike{p.Key: tools.CreateLike(val)}.ToSql()
}

func (p *Param) getLtGtString() (query string, args []interface{}, err error) {
	return p.getDefault()
}

func (p *Param) getLtGtFloat(val float64) (query string, args []interface{}, err error) {
	return p.getDefault()
}

func (p *Param) getDefault() (query string, args []interface{}, err error) {
	return strings.Join([]string{p.Key, p.Operator, "?"}, ""), []interface{}{p.Value.Item}, nil
}

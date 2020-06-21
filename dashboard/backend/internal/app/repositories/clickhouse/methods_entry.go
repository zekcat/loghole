package clickhouse

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/gadavy/tracing"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
	"github.com/lissteron/loghole/dashboard/internal/app/repositories/clickhouse/models"
)

const (
	stringParams = "(has(params_string.keys, '%s') AND params_string.values[indexOf(params_string.keys, '%s')] %s '%s')"
	floatParams  = "(has(params_float.keys, '%s') AND params_float.values[indexOf(params_float.keys, '%s')] %s %s)"
)

func (r *Repository) ListEntry(
	ctx context.Context,
	params [][]*domain.QueryParam,
	limit,
	offset int64,
) ([]*domain.Entry, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	query, _, err := r.buildListEntryQuery(ctx, params, limit, offset)
	if err != nil {
		return nil, err
	}

	var dest []*models.Entry

	if err := r.db.SelectContext(ctx, &dest, query); err != nil {
		return nil, err
	}

	result := make([]*domain.Entry, 0, len(dest))

	for _, val := range dest {
		result = append(result, val.ToDomain())
	}

	return result, nil
}

func (r *Repository) buildListEntryQuery(
	ctx context.Context,
	params [][]*domain.QueryParam,
	limit,
	offset int64,
) (query string, args []interface{}, err error) {
	defer tracing.ChildSpan(&ctx).Finish()

	return squirrel.Select("time", "nsec", "namespace", "source", "host",
		"trace_id", "message", "params", "build_commit", "config_hash").
		From("internal_logs_buffer").
		Where(r.buildListEntryWhere(params)).
		OrderBy("time DESC").
		Suffix(fmt.Sprintf("LIMIT %d, %d", offset, limit)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
}

func (r *Repository) buildListEntryWhere(params [][]*domain.QueryParam) (where string) {
	builder := make([]string, 0, len(params))

	for _, list := range params {
		builder = append(builder, r.prepareParamList(list))
	}

	return strings.Join(builder, " OR ")
}

func (r *Repository) prepareParamList(list []*domain.QueryParam) (where string) {
	builder := make([]string, 0, len(list))

	for _, param := range list {
		builder = append(builder, r.prepareParam(param))
	}

	return strings.Join([]string{"(", strings.Join(builder, " AND "), ")"}, "")
}

func (r *Repository) prepareParam(param *domain.QueryParam) string {
	if param.Type == domain.TypeKey {
		return r.prepareJSONParam(param)
	}

	return strings.Join([]string{param.Key, param.Operator, "'", param.Value, "'"}, "")
}

func (r *Repository) prepareJSONParam(param *domain.QueryParam) string {
	switch param.Operator {
	case "<", ">":
		return fmt.Sprintf(floatParams, param.Key, param.Key, param.Operator, param.Value)
	default:
		return fmt.Sprintf(stringParams, param.Key, param.Key, param.Operator, param.Value)
	}
}

package clickhouse

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/gadavy/tracing"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
	"github.com/lissteron/loghole/dashboard/internal/app/repositories/clickhouse/models"
)

const (
	stringParams = "(has(params_string.keys, '%s') AND params_string.values[indexOf(params_string.keys, '%s')] %s ?)"
	floatParams  = "(has(params_float.keys, '%s') AND params_float.values[indexOf(params_float.keys, '%s')] %s ?)"
)

func (r *Repository) ListEntry(
	ctx context.Context,
	params [][]*domain.QueryParam,
	limit,
	offset int64,
) ([]*domain.Entry, error) {
	defer tracing.ChildSpan(&ctx).Finish()

	query, args, err := r.buildListEntryQuery(ctx, params, limit, offset)
	if err != nil {
		return nil, err
	}

	var dest []*models.Entry

	if err := r.db.SelectContext(ctx, &dest, query, args...); err != nil {
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

	where, args := buildListEntryWhere(params)

	return squirrel.Select("time", "nsec", "namespace", "source", "host",
		"trace_id", "message", "params", "build_commit", "config_hash").
		From("internal_logs_buffer").
		Where(where, args...).
		OrderBy("time DESC").
		Suffix(fmt.Sprintf("LIMIT %d, %d", offset, limit)).
		PlaceholderFormat(squirrel.Question).
		ToSql()
}

func buildListEntryWhere(params [][]*domain.QueryParam) (where string, args []interface{}) {
	builder := make([]string, 0, len(params))
	args = make([]interface{}, 0)

	for _, list := range params {
		buf := make([]string, 0, len(list))

		for _, param := range list {
			q, arg := prepareListEntryParam(param)

			args, buf = append(args, arg), append(buf, q)
		}

		builder = append(builder, strings.Join([]string{"(", strings.Join(buf, " AND "), ")"}, ""))
	}

	return strings.Join(builder, " OR "), args
}

func prepareListEntryParam(param *domain.QueryParam) (q string, arg interface{}) {
	if param.Type == domain.TypeKey {
		return prepareListEntryJSONParam(param)
	}

	return strings.Join([]string{param.Key, param.Operator, "?"}, ""), param.Value
}

func prepareListEntryJSONParam(param *domain.QueryParam) (q string, arg interface{}) {
	if param.Operator == "<" || param.Operator == ">" {
		value, err := strconv.ParseFloat(param.Value, 64)
		if err == nil {
			return fmt.Sprintf(floatParams, param.Key, param.Key, param.Operator), value
		}
	}

	return fmt.Sprintf(stringParams, param.Key, param.Key, param.Operator), param.Value
}

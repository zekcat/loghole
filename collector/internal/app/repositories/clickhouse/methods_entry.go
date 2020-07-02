package clickhouse

import (
	"context"

	"github.com/gadavy/tracing"

	"github.com/lissteron/loghole/collector/internal/app/domain"
)

func (r *Repository) StoreEntryList(ctx context.Context, list []*domain.Entry) (err error) {
	defer tracing.ChildSpan(&ctx).Finish()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func() { _ = tx.Rollback() }()

	query := `INSERT INTO internal_logs_buffer (time,date,nsec,namespace,source,host,leve,trace_id,message,params,
		params_string.keys,params_string.values,params_float.keys,params_float.values,build_commit,config_hash)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	for _, val := range list {
		_, err = tx.Exec(query, val.Time, val.Time, val.Time.UnixNano(), val.Namespace, val.Source,
			val.Host, val.Level, val.TraceID, val.Message, val.Params, val.StringKey,
			val.StringVal, val.FloatKey, val.FloatVal, val.BuildCommit, val.ConfigHash)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}


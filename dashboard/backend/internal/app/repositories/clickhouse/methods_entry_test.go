package clickhouse

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
)

func TestRepository_buildListEntryQuery(t *testing.T) {
	tests := []struct {
		name          string
		input         [][]*domain.QueryParam
		expectedQuery string
		expectedArgs   []interface{}
	}{
		{
			name: "#1",
			input: [][]*domain.QueryParam{
				{
					{
						Type:     "column",
						Join:     "AND",
						Key:      "namespace",
						Value:    "prod",
						Operator: "=",
					},
				},
			},
			expectedQuery: "SELECT time, nsec, namespace, source, host, trace_id, message, params, build_commit, config_hash FROM internal_logs_buffer WHERE (namespace=?) ORDER BY time DESC LIMIT 0, 100",
			expectedArgs:  []interface{}{"prod"},
		},
		{
			name: "#2",
			input: [][]*domain.QueryParam{
				{
					{
						Type:     "column",
						Join:     "AND",
						Key:      "namespace",
						Value:    "prod",
						Operator: "=",
					},
					{
						Type:     "column",
						Join:     "AND",
						Key:      "time",
						Value:    "2020-06-23 10:22:59",
						Operator: ">",
					},
				},
				{
					{
						Type:     "column",
						Join:     "AND",
						Key:      "time",
						Value:    "2020-06-23 10:22:59",
						Operator: "<",
					},
					{
						Type:     "key",
						Join:     "AND",
						Key:      "kf2",
						Value:    "55",
						Operator: "<",
					},
				},
			},
			expectedQuery: "SELECT time, nsec, namespace, source, host, trace_id, message, params, build_commit, config_hash FROM internal_logs_buffer WHERE (namespace=? AND time>?) OR (time<? AND (has(params_float.keys, 'kf2') AND params_float.values[indexOf(params_float.keys, 'kf2')] < ?)) ORDER BY time DESC LIMIT 0, 100",
			expectedArgs:  []interface{}{"prod", "2020-06-23 10:22:59", "2020-06-23 10:22:59", float64(55)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{}

			query, args, err := repo.buildListEntryQuery(context.TODO(), tt.input, 100, 0)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expectedQuery, query)
			assert.Equal(t, tt.expectedArgs, args)
		})
	}
}

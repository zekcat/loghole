package clickhouse

import (
	"context"
	"log"
	"testing"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
)

func TestRepository_ListEntry(t *testing.T) {
	tests := []struct {
		name          string
		input         *domain.Query
		expectedQuery string
		expectedArgs  []interface{}
	}{
		{
			name: "#1",
			input: &domain.Query{
				Params: []*domain.QueryParam{
					{
						Type: "column",
						Key:  "time",
						Value: domain.ParamValue{
							Item: "2020-06-21 10:22:59",
						},
						Operator: "=",
					},
				},
				Limit:  100,
				Offset: 0,
			},
			expectedQuery: "SELECT time, nsec, namespace, source, host, trace_id, message, params, build_commit, config_hash FROM internal_logs_buffer WHERE time=? ORDER BY time DESC LIMIT 0, 100",
			expectedArgs:  []interface{}{"2020-06-23 10:22:59"},
		},
		{
			name: "#2",
			input: &domain.Query{
				Params: []*domain.QueryParam{
					{
						Type: "column",
						Key:  "time",
						Value: domain.ParamValue{
							Item: "2020-06-21 10:22:59",
						},
						Operator: "=",
					},
					{
						Type: "column",
						Key:  "namespace",
						Value: domain.ParamValue{
							Item: "prod",
						},
						Operator: "LIKE",
					},
				},
				Limit:  100,
				Offset: 0,
			},
			expectedQuery: "SELECT time, nsec, namespace, source, host, trace_id, message, params, build_commit, config_hash FROM internal_logs_buffer WHERE namespace LIKE ? ORDER BY time DESC LIMIT 0, 100",
			expectedArgs:  []interface{}{"2020-06-21 10:22:59", "%prod%"},
		},
		{
			name: "#3",
			input: &domain.Query{
				Params: []*domain.QueryParam{
					{
						Type: "column",
						Key:  "namespace",
						Value: domain.ParamValue{
							List: []string{"prod", "prod1"},
						},
						Operator: "IN",
					},
				},
				Limit:  100,
				Offset: 0,
			},
			expectedQuery: "SELECT time, nsec, namespace, source, host, trace_id, message, params, build_commit, config_hash FROM internal_logs_buffer WHERE namespace IN (?,?) ORDER BY time DESC LIMIT 0, 100",
			expectedArgs:  []interface{}{"prod", "prod1"},
		},
		{
			name: "#4",
			input: &domain.Query{
				Params: []*domain.QueryParam{
					{
						Type: "column",
						Key:  "namespace",
						Value: domain.ParamValue{
							List: []string{"prod", "prod1"},
						},
						Operator: "NOT IN",
					},
				},
				Limit:  100,
				Offset: 0,
			},
			expectedQuery: "SELECT time, nsec, namespace, source, host, trace_id, message, params, build_commit, config_hash FROM internal_logs_buffer WHERE namespace NOT IN (?,?) ORDER BY time DESC LIMIT 0, 100",
			expectedArgs:  []interface{}{"prod", "prod1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, args, err := buildListEntryQuery(context.TODO(), tt.input)
			if err != nil {
				t.Fatal(err)
			}

			log.Println(query)
			log.Println(args...)
		})
	}
}

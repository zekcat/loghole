package clickhouse

import (
	"context"
	"log"
	"testing"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
)

func TestRepository_buildListEntryQuery(t *testing.T) {
	tests := []struct {
		name          string
		input         [][]*domain.QueryParam
		expectedQuery string
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
			expectedQuery: "SELECT * FROM internal_logs_buffer WHERE (namespace='prod') ORDER BY time DESC LIMIT 0, 100",
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
			expectedQuery: "SELECT * FROM internal_logs_buffer WHERE (namespace='prod') ORDER BY time DESC LIMIT 0, 100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{}

			query, args, err := repo.buildListEntryQuery(context.TODO(), tt.input, 100, 0)
			if err != nil {
				t.Fatal(err)
			}

			log.Println(query)
			log.Println(args)
		})
	}
}

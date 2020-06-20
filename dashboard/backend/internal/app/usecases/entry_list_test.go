package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lissteron/loghole/dashboard/internal/app/domain"
)

func TestListLogs_prepareParams(t *testing.T) {
	tests := []struct {
		name         string
		input        []*domain.QueryParam
		expectedRes  [][]*domain.QueryParam
	}{
		{
			name: "#1",
			input: []*domain.QueryParam{
				{
					Type:     "column",
					Join:     "AND",
					Key:      "key1",
					Value:    "val1",
					Operator: "=",
				},
				{
					Type:     "column",
					Join:     "AND",
					Key:      "key2",
					Value:    "val2",
					Operator: "=",
				},
				{
					Type:     "column",
					Join:     "AND",
					Key:      "key3",
					Value:    "val3",
					Operator: "=",
				},
			},
			expectedRes: [][]*domain.QueryParam{
				{
					{
						Type:     "column",
						Join:     "AND",
						Key:      "key1",
						Value:    "val1",
						Operator: "=",
					},
					{
						Type:     "column",
						Join:     "AND",
						Key:      "key2",
						Value:    "val2",
						Operator: "=",
					},
					{
						Type:     "column",
						Join:     "AND",
						Key:      "key3",
						Value:    "val3",
						Operator: "=",
					},
				},
			},
		},
		{
			name: "#2",
			input: []*domain.QueryParam{
				{
					Type:     "column",
					Join:     "AND",
					Key:      "key1",
					Value:    "val1",
					Operator: "=",
				},
				{
					Type:     "column",
					Join:     "OR",
					Key:      "key2",
					Value:    "val2",
					Operator: "=",
				},
				{
					Type:     "column",
					Join:     "AND",
					Key:      "key3",
					Value:    "val3",
					Operator: "=",
				},
			},
			expectedRes: [][]*domain.QueryParam{
				{
					{
						Type:     "column",
						Join:     "AND",
						Key:      "key1",
						Value:    "val1",
						Operator: "=",
					},
				},
				{
					{
						Type:     "column",
						Join:     "OR",
						Key:      "key2",
						Value:    "val2",
						Operator: "=",
					},
					{
						Type:     "column",
						Join:     "AND",
						Key:      "key3",
						Value:    "val3",
						Operator: "=",
					},
				},
			},
		},
		{
			name: "#3",
			input: []*domain.QueryParam{
				{
					Type:     "key",
					Join:     "AND",
					Key:      "key1",
					Value:    "val1",
					Operator: "=",
				},
				{
					Type:     "column",
					Join:     "OR",
					Key:      "key2",
					Value:    "val2",
					Operator: "=",
				},
				{
					Type:     "column",
					Join:     "AND",
					Key:      "key3",
					Value:    "val3",
					Operator: "=",
				},
			},
			expectedRes: [][]*domain.QueryParam{
				{
					{
						Type:     "key",
						Join:     "AND",
						Key:      "key1",
						Value:    "val1",
						Operator: "=",
					},
				},
				{
					{
						Type:     "column",
						Join:     "OR",
						Key:      "key2",
						Value:    "val2",
						Operator: "=",
					},
					{
						Type:     "column",
						Join:     "AND",
						Key:      "key3",
						Value:    "val3",
						Operator: "=",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := ListEntry{}

			result := usecase.prepareParams(tt.input)

			assert.Equal(t, tt.expectedRes, result)
		})
	}
}

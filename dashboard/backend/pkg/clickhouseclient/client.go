package clickhouseclient

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/ClickHouse/clickhouse-go" // driver
)

const connStr = "tcp://%s?username=%s&database=%s&read_timeout=10&write_timeout=20"

type Client struct {
	db *sqlx.DB
}

func NewClient(addr, user, database string) (*Client, error) {
	db, err := sqlx.Connect("clickhouse", fmt.Sprintf(connStr, addr, user, database))
	if err != nil {
		return nil, err
	}

	return &Client{db: db}, nil
}

func (c *Client) Client() *sqlx.DB {
	return c.db
}

func (c *Client) Close() error {
	return c.db.Close()
}

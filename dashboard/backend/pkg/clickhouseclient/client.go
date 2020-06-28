package clickhouseclient

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/ClickHouse/clickhouse-go" // driver
)

const connStr = "tcp://%s?username=%s&database=%s&read_timeout=%d&write_timeout=%d"

type Options struct {
	Addr         string
	User         string
	Database     string
	ReadTimeout  int
	WriteTimeout int
}

type Client struct {
	db *sqlx.DB
}

func NewClient(options *Options) (*Client, error) {
	db, err := sqlx.Connect("clickhouse", connString(options))
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

func connString(options *Options) string {
	return fmt.Sprintf(connStr, options.Addr, options.User, options.Database, options.ReadTimeout, options.WriteTimeout)
}

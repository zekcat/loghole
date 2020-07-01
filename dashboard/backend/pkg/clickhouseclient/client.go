package clickhouseclient

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/ClickHouse/clickhouse-go"
)

const (
	codeNoDB = 81
)

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
		if isDBNotFound(err) {
			if err := initDB(options); err != nil {
				return nil, err
			}

			return NewClient(options)
		}

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

func isDBNotFound(err error) bool {
	if exception, ok := err.(*clickhouse.Exception); ok {
		return exception.Code == codeNoDB
	}

	return false
}

func initDB(options *Options) error {
	f, err := os.Open("/db/init.sql")
	if err != nil {
		return err
	}

	defer f.Close()

	query, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	db, err := sqlx.Connect("clickhouse", connStringWithoutDB(options))
	if err != nil {
		return err
	}

	defer db.Close()

	for _, val := range strings.Split(string(query), ";") {
		if val = strings.TrimSpace(val); val == "" {
			continue
		}

		if _, err = db.Exec(val); err != nil {
			return err
		}
	}

	return nil
}

func connStringWithoutDB(options *Options) string {
	return fmt.Sprintf("tcp://%s?username=%s&read_timeout=%d&write_timeout=%d",
		options.Addr, options.User, options.ReadTimeout, options.WriteTimeout)
}

func connString(options *Options) string {
	return fmt.Sprintf("tcp://%s?username=%s&database=%s&read_timeout=%d&write_timeout=%d",
		options.Addr, options.User, options.Database, options.ReadTimeout, options.WriteTimeout)
}

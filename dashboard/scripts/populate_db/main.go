package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	_ "github.com/ClickHouse/clickhouse-go" // driver
)

const connStr = "tcp://%s?username=%s&database=%s&read_timeout=10&write_timeout=20"

func main() {
	viper.AutomaticEnv()

	time.Sleep(viper.GetDuration("SLEEP"))

	db, err := sqlx.Connect("clickhouse",
		fmt.Sprintf(connStr,
			viper.GetString("CLICKHOUSE_URI"),
			viper.GetString("CLICKHOUSE_USER"),
			viper.GetString("CLICKHOUSE_DATABASE"),
		),
	)
	if err != nil {
		log.Fatalln("db connect", err)
	}

	defer db.Close()

	generator := NewErrorGenerator()

	query := `INSERT INTO internal_logs_buffer (
		namespace, source, host, level, trace_id, message, params, build_commit, config_hash)
		VALUES (?,?,?,?,?,?,?,?,?,?)
`

	log.Println("start")

	for i := 0; i < viper.GetInt("COUNT"); i++ {
		entry, err := generator.GenerateEntry()
		if err != nil {
			log.Println("GenerateEntry", err)
			continue
		}

		tx, err := db.Begin()
		if err != nil {
			log.Println("Begin", err)
			continue
		}

		_, err = tx.Exec(query,
			entry.Namespace,
			entry.Source,
			entry.Host,
			entry.Level,
			entry.TraceID,
			entry.Message,
			entry.Params,
			entry.BuildCommit,
			entry.ConfigHash)
		if err != nil {
			log.Println("Exec", err)
			continue
		}

		if err = tx.Commit(); err != nil {
			log.Println("Commit", err)
			continue
		}

		time.Sleep(time.Millisecond * 3)
	}

	log.Println("success")
}

type ErrorGenerator struct {
	template     []string
	errorTxt     []string
	loggerParams []params
	rnd          *rand.Rand
}

func NewErrorGenerator() *ErrorGenerator {
	return &ErrorGenerator{
		template:     errTemplates,
		errorTxt:     errText,
		loggerParams: loggerParams,
		rnd:          rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (e *ErrorGenerator) randomMessage() string {
	return fmt.Sprintf(e.template[e.rnd.Intn(len(e.template)-1)], e.errorTxt[e.rnd.Intn(len(e.errorTxt)-1)])
}

type Entry struct {
	Time        time.Time `json:"time"`
	NSec        int64     `json:"n_sec"`
	Namespace   string    `json:"namespace"`
	Source      string    `json:"source"`
	Host        string    `json:"host"`
	Level       string    `json:"level"`
	TraceID     string    `json:"trace_id"`
	Message     string    `json:"message"`
	Params      string    `json:"params,omitempty"`
	BuildCommit string    `json:"build_commit"`
	ConfigHash  string    `json:"config_hash"`
}

func (e *ErrorGenerator) GenerateEntry() (*Entry, error) {
	now := time.Now()
	param := e.randomParams()

	entry := &Entry{
		Time:        now,
		NSec:        now.UnixNano(),
		Namespace:   param.namespace,
		Source:      param.source,
		Host:        param.host,
		Level:       e.randomLevel(),
		TraceID:     strconv.FormatInt(time.Now().UnixNano(), 32) + strconv.FormatInt(rand.Int63n(1000), 32),
		Message:     e.randomMessage(),
		Params:      "",
		BuildCommit: param.buildCommit,
		ConfigHash:  param.configHash,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return nil, err
	}

	entry.Params = string(data)

	return entry, nil
}

func (e *ErrorGenerator) randomParams() params {
	return e.loggerParams[e.rnd.Intn(len(e.loggerParams)-1)]
}

func (e *ErrorGenerator) randomLevel() string {
	switch e.rnd.Intn(4) {
	case 0:
		return "debug"
	case 1:
		return "info"
	case 2:
		return "warn"
	case 3:
		return "error"
	default:
		return "info"
	}
}

type params struct {
	buildCommit string
	configHash  string
	host        string
	source      string
	namespace   string
}

var loggerParams = []params{
	{
		buildCommit: "db957a22b3c1d6e508c0828917a5e14c572fb007",
		configHash:  "130362a6fd10cf2f939dd0cfc0ab222cee6a99ec",
		host:        "127.100.0.1:50000",
		source:      "app_1",
		namespace:   "prod",
	},
	{
		buildCommit: "7fca4555eb6f5eb20ce7f81900595b9c4ee5f47e",
		configHash:  "f365c5c9e9217df5cd3271d3cd25a5ebbf94cfeb",
		host:        "127.11.10.1:45673",
		source:      "app_1",
		namespace:   "dev",
	},
	{
		buildCommit: "0448b96036e3683623a9d6712fd6f6f891607b90",
		configHash:  "4de34c86ce4234270d487c4dd36fdfefe4566bd9",
		host:        "127.11.10.2:45655",
		source:      "app_2",
		namespace:   "dev",
	},
	{
		buildCommit: "d3a8986332402061f9667aee033b33bd890bb96f",
		configHash:  "1c163de678aac3a8ce1582828fd5bea86773ba19",
		host:        "127.100.10.2:43355",
		source:      "app_2",
		namespace:   "prod",
	},
}

var errTemplates = []string{
	"read login request failed: %v",
	"authenticate password: %v",
	"read registration request failed: %v",
	"register user failed: %v",
	"read logout request failed: %v",
	"remove session token failed: %v",
	"read logout request failed: %v",
	"remove all tokens failed: %v",
	"authenticate token failed: %v",
	"read create catalog request failed: %v",
	"create catalog failed: %v",
	"read find catalog request failed: %v",
	"find catalog by url failed: %v",
	"read update catalog request failed: %v",
	"get catalogs list failed: %v",
	"read update catalog request failed: %v",
	"update catalog failed: %v",
	"read confirm email request failed: %v",
	"confirm email failed: %v",
	"read find user confirm request failed: %v",
	"find user confirm failed: %v",
	"read resend confirm request failed: %v",
	"resend confirm failed: %v",
	"read create item request failed: %v",
	"create item failed: %v",
	"read find item request failed: %v",
	"find item failed: %v",
	"read list item request failed: %v",
	"get items list failed: %v",
	"read update item request failed: %v",
	"update item failed: %v",
	"read create notification addr request failed: %v",
	"create notification addr failed: %v",
	"read list notification addr request failed: %v",
	"list notification addr failed: %v",
	"read update notification addr request failed: %v",
	"update notification addr failed: %v",
	"read find order request failed: %v",
	"find order failed: %v",
	"read find order request failed: %v",
	"find order failed: %v",
	"read list orders request failed: %v",
	"list orders failed: %v",
	"read create order request failed: %v",
	"create order failed: %v",
	"read update order request failed: %v",
	"update order failed: %v",
	"read create shop domain request failed: %v",
	"hold shop domain failed: %v",
	"read create shop request failed: %v",
	"create shop failed: %v",
	"read find shop request failed: %v",
	"find shop failed: %v",
	"read list shops request failed: %v",
	"get shops list failed: %v",
	"read find shop request failed: %v",
	"update shop failed: %v",
	"read list measures request failed: %v",
	"get list measures failed: %v",
	"read exec request failed: %v",
	"close exec request body failed: %v",
	"wrap error: %v",
	"read failed: %v",
	"do failed: %v",
	"read failed: %v",
	"do failed: %v",
	"read failed: %v",
	"do failed: %v",
	"read failed: %v",
	"do failed: %v",
	"list active photos failed: %v",
	"list minio photos failed: %v",
	"generate url failed: %v",
	"process image failed: %v",
	"put image failed: %v",
	"remove image failed: %v",
	"find token key failed: %v",
	"update expire in redis failed: %v",
	"generate session key failed: %v",
	"set data to redis failed: %v",
	"generate sign failed: %v",
	"get data from redis failed: %v",
	"remove data from redis failed: %v",
	"get data from redis failed: %v",
	"create catalog failed: %v",
	"find catalog by url failed: %v",
	"create catalog failed: %v",
	"find shop by domain failed: %v",
	"find catalog failed: %v",
	"find shop by domain failed: %v",
	"get catalogs list failed: %v",
	"find catalogs info failed: %v",
	"update catalog failed: %v",
	"find catalog by id failed: %v",
	"find catalog by url failed: %v",
	"update catalog failed: %v",
	"find catalog failed: %v",
	"find item by url failed: %v",
	"create item failed: %v",
	"find shop by domain failed: %v",
	"find item failed: %v",
	"find shop by domain failed: %v",
	"get items list failed: %v",
	"find item by id failed: %v",
	"update item failed: %v",
	"find item by url failed: %v",
	"check catalog owner failed: %v",
	"find measures failed: %v",
	"find user failed: %v",
	"list notification addr failed: %v",
	"find notification addr failed: %v",
	"store notification addr failed: %v",
	"find user failed: %v",
	"find notification addr failed: %v",
	"find notification failed: %v",
	"update notification failed: %v",
	"find shop failed: %v",
	"confirm order failed: %v",
	"find shop failed: %v",
	"find order failed: %v",
	"get order list by shop id failed: %v",
	"get order items failed: %v",
	"find shop failed: %v",
	"store order item failed: %v",
	"find order by id failed: %v",
	"update order failed: %v",
	"create shop failed: %v",
	"find shop failed: %v",
	"hold shop domain failed: %v",
	"find shops by owner id failed: %v",
	"find shops info failed: %v",
	"find shop failed: %v",
	"update shop failed: %v",
	"authenticate token failed: %v",
	"find shop by domain failed: %v",
	"validate email confirmation failed: %v",
	"confirm user failed: %v",
	"find user failed: %v",
	"find user failed: %v",
	"find user by id failed: %v",
	"find user by id failed: %v",
	"send email confirmation failed: %v",
	"find user by email failed: %v",
	"generate password hash failed: %v",
	"create user failed: %v",
	"generate token failed: %v",
	"add user session failed: %v",
	"send email confirmation message failed: %v",
	"find user failed: %v",
	"check active sessions failed: %v",
	"generate token failed: %v",
	"add user session failed: %v",
	"find user by id failed: %v",
	"update user failed: %v",
	"get file stat failed: %v",
	"convert message failed: %v",
	"convert notification message failed: %v",
	"send order notification failed: %v",
	"convert confirmation message failed: %v",
	"send confirmation failed: %v",
	"store notification failed: %v",
	"list shop notification addr failed: %v",
	"send notification failed: %v",
	"store send info failed: %v",
	"remove image failed: %v",
}

var errText = []string{
	"some message 1",
	"some message 2",
	"some message 3",
	"some message 4",
	"some message 5",
	"some message 6",
}

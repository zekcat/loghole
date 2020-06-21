CREATE DATABASE logs;

CREATE TABLE logs.internal_logs (
  `time` DateTime DEFAULT now(),
  `nsec` UInt64,
  `namespace` LowCardinality(String), 
  `source` LowCardinality(String), 
  `host` LowCardinality(String),
  `trace_id` String,
  `message` String,
  `params` String,
  `params_string.keys` Array(String),
  `params_string.values` Array(String),
  `params_float.keys` Array(String),
  `params_float.values` Array(Float64),
  `build_commit` String,
  `config_hash` String
) ENGINE = MergeTree()
PARTITION BY (time)
ORDER BY (time, namespace, source)
SETTINGS index_granularity = 8192;

CREATE TABLE logs.internal_logs_buffer AS logs.internal_logs ENGINE = Buffer(logs, internal_logs, 16, 10, 100, 10000, 1000000, 10000000, 100000000);

SELECT * FROM internal_logs ARRAY JOIN `params.keys` 
    AS keys, arrayEnumerate(`params.keys`) 
    AS idx WHERE keys='n2' AND arrayElement(`params.values`, idx) LIKE '%v3%';
package domain

const (
	TypeKey    = "key"
	TypeColumn = "column"
)

const (
	JoinOr  = "OR"
	JoinAnd = "AND"
)

type Query struct {
	Params []*QueryParam `json:"params"`
	Limit  int64         `json:"limit"`
	Offset int64         `json:"offset"`
}

type QueryParam struct {
	Type     string `json:"type"`
	Join     string `json:"join"`
	Key      string `json:"key"`
	Value    string `json:"value"`
	Operator string `json:"operator"`
}

/*
SELECT * FROM (
    SELECT * FROM internal_logs ARRAY JOIN `params_string`
) ARRAY JOIN `params_float` WHERE (params_string.keys='ks1' AND `params_string.values` = 'vs3') AND (params_float.keys='kf1' AND `params_float.values` > 0)

 */
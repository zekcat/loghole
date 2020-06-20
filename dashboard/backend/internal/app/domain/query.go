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

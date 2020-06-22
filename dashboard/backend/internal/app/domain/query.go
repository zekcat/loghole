package domain

const (
	TypeKey    = "key"
	TypeColumn = "column"
)

const (
	OperatorIn      = "IN"
	OperatorNotIn   = "NOT IN"
	OperatorLike    = "LIKE"
	OperatorNotLike = "NOT LIKE"
)

type Query struct {
	Params []*QueryParam `json:"params"`
	Limit  int64         `json:"limit"`
	Offset int64         `json:"offset"`
}

type QueryParam struct {
	Type     string     `json:"type"`
	Key      string     `json:"key"`
	Value    ParamValue `json:"value"`
	Operator string     `json:"operator"`
}

func (p *QueryParam) IsTypeJSON() bool {
	return p.Type == TypeKey
}

type ParamValue struct {
	Item string   `json:"item"`
	List []string `json:"list"`
}

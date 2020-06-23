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
	Params []*QueryParam
	Limit  int64
	Offset int64
}

type QueryParam struct {
	Type     string
	Key      string
	Value    ParamValue
	Operator string
}

func (p *QueryParam) IsTypeJSON() bool {
	return p.Type == TypeKey
}

type ParamValue struct {
	Item string
	List []string
}

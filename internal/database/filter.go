package database

const DefaultLimit = 100

const (
	Eq   Operation = "eq"
	Gt   Operation = "gt"
	Lt   Operation = "lt"
	Like Operation = "like"
	In   Operation = "in"
)

type Operation string

type SQLFilter struct {
	Op    Operation
	Field string
	Value string
}

type Filter struct {
	Filters []SQLFilter
	Limit   int
	Offset  int
	OrderBy string // e.g., created_at DESC
}

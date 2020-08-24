package postgo

type Query interface{
	ToString() (string, error)
	GetStatement() (string, error)
	GetParameters() []interface{}
}

func Select(args ...string) *SelectQuery {
	if len(args) == 0 {
		return &SelectQuery{}
	}
	return &SelectQuery{
		cols: args,
	}
}
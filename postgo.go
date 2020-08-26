package postgo

import (
	"fmt"
)

func From(table string, columns ...string) *Query {
	return &Query{
		source: table,
		cols:   columns,
	}
}

func FromSub(sub *Query, columns ...string) *Query {
	return &Query{
		source: fmt.Sprintf("(%s)", sub.Statement()),
		cols:   columns,
	}
}

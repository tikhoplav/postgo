package postgo

import (
	"fmt"
)

type Query struct {
	From string
}

func (q *Query) ToString() string {
	return fmt.Sprintf("SELECT * FROM %s", q.From)
}
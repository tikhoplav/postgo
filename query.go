package postgo

import (
	"fmt"
	"strings"
)

type Query struct {
	source string
	alias string
	cols []string
}

func (q *Query) ToString() string {
	return fmt.Sprintf("%s;", q.Statement())
}

func (q *Query) Statement() string {
	source := q.sourceDef()
	cols := q.colsDef()
	return fmt.Sprintf("SELECT %s FROM %s", cols, source) 
}

func (q *Query) sourceDef() string {
	if q.alias != "" {
		return fmt.Sprintf("%s AS %s", q.source, q.alias)
	}
	return q.source
}

func (q *Query) colsDef() string {
	switch len(q.cols) {
	case 0:
		return "*"
	case 1:
		return q.cols[0]
	default:
		return strings.Join(q.cols, ", ")
	}
}

func (q *Query) As(alias string) *Query {
	q.alias = alias
	return q
}

func (q *Query) Select(column string) *Query {
	q.cols = append(q.cols, column)
	return q
}

func (q *Query) SelectAs(column string, alias string) *Query {
	col := fmt.Sprintf("%s AS %s", column, alias)
	q.cols = append(q.cols, col)
	return q
}
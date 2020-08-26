package postgo

import (
	"fmt"
	"strings"
)

// SelectQuery is a SQL query builder class for select statements.
// Do not instantiate SelectQuery object, use FROM function instead.
type SelectQuery struct {
	source string
	alias  string
	cols   []string
}

func (q *SelectQuery) ToSQL() string {
	// Place parameters values
	// Form finished SQL query
	return fmt.Sprintf("%s;", q.Statement())
}

func (q *SelectQuery) Statement() string {
	source := q.sourceDef()
	cols := q.colsDef()
	return fmt.Sprintf("SELECT %s FROM %s", cols, source)
}

// Returns ordered array of parameters appliable to the query.
// Order of this array should be preserved until query execution.
func (q *SelectQuery) Parameters() []interface{} {
	return make([]interface{}, 0)
}

func (q *SelectQuery) sourceDef() string {
	if q.alias != "" {
		return fmt.Sprintf("%s AS %s", q.source, q.alias)
	}
	return q.source
}

func (q *SelectQuery) colsDef() string {
	switch len(q.cols) {
	case 0:
		return "*"
	case 1:
		return q.cols[0]
	default:
		return strings.Join(q.cols, ", ")
	}
}

// Set alias for the source of the query.
// If query uses Joins, this alias should be used as
// prefix for each column name.
func (q *SelectQuery) As(alias string) *SelectQuery {
	q.alias = alias
	return q
}

// Adds column to the SELECT definition of the query.
// Raw SQL statement can be used as well as column name.
// If query uses Join, it is recomended to put table alias as
// prefixes for each item, in order to avoid ambiguous selects.
func (q *SelectQuery) Select(column string) *SelectQuery {
	q.cols = append(q.cols, column)
	return q
}

// Adds colum to the SELECT definition of the query with alias.
// It is not recomended to use user input as alias due to fact,
// that aliases added this way would not be escaped.
// If it is necessary to use user input as alias, please, provide
// proper escape for them, for example with fmt.Sprintf("%f", alias)
func (q *SelectQuery) SelectAs(column string, alias string) *SelectQuery {
	col := fmt.Sprintf("%s AS %s", column, alias)
	q.cols = append(q.cols, col)
	return q
}

// Package postgo provides PostgreSQL query builder.
// Based on PostreSQL version 12 documentation:
// https://www.postgresql.org/docs/12/sql-select.html
//
// This library designed to produce prepared SQL statements
// (with parameters and their values escaped), since it is a most
// advised and secure way to execute a query. Despite that, it also
// gives an ability to produce plain SQL (for debugging, for example).
//
// This package does not use or depends on any particular database driver,
// But acts as a raw SQL string builder for PostgreSQL.
package postgo

import (
	"fmt"
)

// Query interface describes all types of queries (SELECT, UPDATE, CREATE and INSERT).
// PostgreSQL allows to use result of some queries as source for other queries,
// this is a main reason of having generic interface.
//
// Statement() is used to generate parametrized SQL statement (with `$1` placeholders).
// Parameters() is used to get copy of array of values of query parameters.
// ToSQL() function is used to generate plain SQL query and not recomended to use in production.
type Query interface {
	Statement() string
	Parameters() []interface{}
	ToSQL() string
}

// Returns new instance of a select query.
// Raw SQL statement (closed in paranthesis) can be used as well as table name.
func From(table string, columns ...string) *SelectQuery {
	return &SelectQuery{
		source: table,
		cols:   columns,
	}
}

// Returns new instance of a select query.
// SQL statement retreived from subquery will be used as source.
// All subquery parameters will be copied to a new query.
// Subquery will not be modified and can be used in other queries after.
func FromSub(sub Query, columns ...string) *SelectQuery {
	return &SelectQuery{
		source: fmt.Sprintf("(%s)", sub.Statement()),
		cols:   columns,
	}
}

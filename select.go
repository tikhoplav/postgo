package postgo

import (
	"fmt"
	"strings"
)

type SelectQuery struct {
	cols []string
	from string
}

func (sq *SelectQuery) ToString() (string, error) {
	cols := "*"
	switch len(sq.cols) {
	case 0:
	case 1:
		cols = sq.cols[0]
	default:
		cols = strings.Join(sq.cols, ", ")
	}
	
	return fmt.Sprintf("SELECT %s FROM %s;", cols, sq.from), nil
}

func (sq *SelectQuery) From(args ...string) *SelectQuery {
	if len(args) == 1 {
		sq.from = args[0]
	} else {
		sq.from = fmt.Sprintf("%s AS %s", args[0], args[1])
	}	
	return sq
}
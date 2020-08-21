package postgo

import "testing"

func TestQuery(t *testing.T) {
    want := "SELECT * FROM users"
    query := &Query{
    	From: "users",
    }
    if got := query.ToString(); got != want {
        t.Errorf("ToSQL() = %q, want %q", got, want)
    }
}
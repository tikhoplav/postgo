package postgo

import (
	"testing"
	"context"
	"github.com/jackc/pgx/v4"
	"os"
)

var dbURL = os.Getenv("DB_URL")

func TestConnection(t *testing.T) {
	_, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		t.Errorf("Failed to connect to DB: %s", err.Error())
	}
}

func TestNativeSelect(t *testing.T) {
	conn, _ := pgx.Connect(context.Background(), dbURL)
	rows, err := conn.Query(context.Background(), "SELECT $1::int, $2", 1, "string")

	if err != nil {
		t.Errorf("Failed to make query: %s", err.Error())
	}

	for rows.Next() {
		var n int32
		var s string
		err = rows.Scan(&n, &s)
		if err != nil {
			t.Errorf("Failed to scan query result: %s", err.Error())
		}
		if n != 1 && s != "string" {
			t.Errorf("Wanted\t| Got:\n1 \t|%d\nstring\t|%s", n, s)
		}
	}
}
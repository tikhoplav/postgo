package postgo

import (
	"testing"
	"context"
	"github.com/jackc/pgx/v4"
	"fmt"
)

func TestCreateTableDropTable(t *testing.T) {
	var sql string
	sql = `CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		name text NOT NULL
	);`

	conn, _ := pgx.Connect(context.Background(), dbURL)
	tag, err := conn.Exec(context.Background(), sql)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	fmt.Println(tag)
}
package postgo

import (
	"testing"
)

func TestCreateBlueprint(t *testing.T) {
	table = &CreateBlueprint {
		name: "users",
	}

	table.Serial("id").Primary()
	table.Text("name").Nullable(false)

	want := `CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		name text NOT NULL
	);`

	got := table.Make()

	if want := got {
		t.Errorf("Want: %t\nGot: %t", want, got)
	}
}

func TestAlterBlueprint(t *testing.T) {
	table := &AlterBlueprint{
		name: "users",
	}

	table.Rename("name", "first_name")
	table.Datetime("created_at").Nullable(false).Default("now()")

	want := `ALTER TABLE users
		RENAME COLUMN name TO first_name,
		ADD COLUMN created_at timestamp DEFAULT now()
	;`

	get := table.Make()

	if want := got {
		t.Errorf("Want: %t\nGot: %t", want, got)
	}
}
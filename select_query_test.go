package postgo

import "testing"

func checkQuery(t *testing.T, want string, query *SelectQuery) {
	if got := query.ToSQL(); got != want {
		t.Errorf("\nWant: %q\n Got: %q", want, got)
	}
}

func TestFrom(t *testing.T) {
	want := "SELECT * FROM users;"
	query := From("users")
	checkQuery(t, want, query)
}

func TestAlias(t *testing.T) {
	want := "SELECT * FROM users AS u;"
	query := From("users").As("u")
	checkQuery(t, want, query)
}

func TestFromSubquery(t *testing.T) {
	want := "SELECT COUNT(*) AS total FROM (SELECT * FROM users);"
	sub := From("users")
	query := FromSub(sub, "COUNT(*) AS total")
	checkQuery(t, want, query)

	query = FromSub(sub).SelectAs("COUNT(*)", "total")
	checkQuery(t, want, query)
}

func TestFromSubWithAlias(t *testing.T) {
	want := "SELECT * FROM (SELECT * FROM users) AS u;"
	sub := From("users")
	query := FromSub(sub).As("u")
	checkQuery(t, want, query)
}

func TestColumn(t *testing.T) {
	want := "SELECT id FROM users;"
	query := From("users", "id")
	checkQuery(t, want, query)

	query = From("users").Select("id")
	checkQuery(t, want, query)
}

func TestColumns(t *testing.T) {
	want := "SELECT id, name FROM users;"
	query := From("users", "id", "name")
	checkQuery(t, want, query)

	query = From("users").Select("id").Select("name")
	checkQuery(t, want, query)
}

func TestAddColumnWithAlias(t *testing.T) {
	want := "SELECT id, name AS user_name, email AS login FROM users;"
	query := From("users", "id", "name AS user_name").SelectAs("email", "login")
	checkQuery(t, want, query)

	query = From("users").
		Select("id").
		SelectAs("name", "user_name").
		SelectAs("email", "login")
	checkQuery(t, want, query)
}

package postgo

import "testing"

func check(t *testing.T, want string, got string) {
	if got != want {
        t.Errorf("\nWant: %q\n Got: %q", want, got)
    }
}

func TestQuery(t *testing.T) {
    want := "SELECT * FROM users;"
    query := Select().From("users")
    got, _ := query.ToString()
    check(t, want, got)
}

func TestFromAs(t *testing.T) {
	want := "SELECT * FROM users AS u;"
	query := Select().From("users", "u")
	got, _ := query.ToString()
    check(t, want, got)
}

func TestSelectColumns(t *testing.T) {
	want := "SELECT name, email FROM users;"
	query := Select("name", "email").From("users")
	got, _ := query.ToString()
	check(t, want, got)
}
package postgo

type SchemaBuilder struct {
	sql string
}

func (sb *SchemaBuilder) make() string {
	return sb.sql
}
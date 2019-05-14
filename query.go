package somesql //import go.lsl.digital/gocipe/somesql

import (
	"database/sql"
	"time"
)

const (
	// AndCondition represents a condition added to the query via AND keyword
	AndCondition uint8 = iota
	// OrCondition represents a condition added to the query via OR keyword
	OrCondition

	// LangEN represents the english language
	LangEN string = "en"
	// LangFR represents the french language
	LangFR string = "fr"
)

const (
	// None represents a simple way of explicitly specifying no value
	None = ""
)

// Query represents a composable query
// Setters: Select(), Save(), Delete(), Where(), SetLang(), SetLimit(), SetOffset()
// Getters: GetLang(), GetLimit(), GetOffset(), GetTx(), IsInner(), AsSQL()
type Query interface {
	Insert(FieldValuer) Query
	Select(fields ...string) Query
	Update(FieldValuer) Query
	Delete() Query
	Where(Condition) Query
	SetLang(lang string) Query
	GetLang() string
	SetLimit(limit int) Query
	GetLimit() int
	SetOffset(offset int) Query
	GetOffset() int
	SetTx(tx *sql.Tx) Query
	GetTx() *sql.Tx
	SetInner(inner bool) Query
	IsInner() bool
	AsSQL() (string, []interface{})
	Exec(autocommit bool) error
}

// FieldValuer assigns a value to a field
type FieldValuer interface {
	ID(id string) FieldValuer
	CreatedAt(t time.Time) FieldValuer
	UpdatedAt(t time.Time) FieldValuer
	OwnerID(id string) FieldValuer
	Status(s string) FieldValuer
	Type(s string) FieldValuer
	Data(json string) FieldValuer
	UseDefaults() FieldValuer
	Set(field string, value interface{}) FieldValuer
	List() ([]string, []interface{})
}

// Condition represents a conditional clause in a statement
type Condition interface {
	ConditionType() uint8
	AsSQL(in ...bool) (string, []interface{})
}

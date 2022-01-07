package table

import (
	"strings"
)

type Field struct {
	ColumnName string
}

func (f *Field) EQ(val interface{}) (string, interface{}) {
	return f.where(" = ?"), val
}

func (f *Field) NEQ(val interface{}) (string, interface{}) {
	return f.where(" <> ?"), val
}

func (f *Field) GT(val interface{}) (string, interface{}) {
	return f.where(" > ?"), val
}

func (f *Field) GTE(val interface{}) (string, interface{}) {
	return f.where(" >= ?"), val
}

func (f *Field) LT(val interface{}) (string, interface{}) {
	return f.where(" < ?"), val
}

func (f *Field) LTE(val interface{}) (string, interface{}) {
	return f.where(" <= ?"), val
}

func (f *Field) IN(val interface{}) (string, interface{}) {
	return f.where(" IN ?"), val
}

func (f *Field) NIN(val interface{}) (string, interface{}) {
	return f.where(" NOT IN ?"), val
}

func (f *Field) LIKE(val string) (string, interface{}) {
	return f.where(" LIKE ?"), val
}

func (f *Field) BETWEEN(start interface{}, end interface{}) (string, interface{}, interface{}) {
	return f.where(" BETWEEN ? AND ?"), start, end
}

func (f *Field) FIS(val string) (string, interface{}) {
	var b strings.Builder
	b.WriteString("FIND_IN_SET(?, ")
	b.WriteString(f.ColumnName)
	b.WriteString(")")
	return b.String(), val
}

func (f *Field) IsNull() string {
	return f.where(" IS NULL")
}

func (f *Field) IsNotNull() string {
	return f.where(" IS NOT NULL")
}

func (f *Field) MAX(alias string) string {
	return f.funcWhere("MAX", alias)
}

func (f *Field) MIN(alias string) string {
	return f.funcWhere("MIN", alias)
}

func (f *Field) COUNT(alias string) string {
	return f.funcWhere("COUNT", alias)
}

func (f *Field) SUM(alias string) string {
	return f.funcWhere("SUM", alias)
}

func (f *Field) ASC() string {
	return f.ColumnName
}

func (f *Field) DESC() string {
	return f.where(" DESC")
}

func (f *Field) where(s string) string {
	var b strings.Builder
	b.WriteString(f.ColumnName)
	b.WriteString(s)
	return b.String()
}

func (f *Field) funcWhere(funcName string, alias string) string {
	if alias == "" {
		alias = strings.ToLower(funcName) + "_" + f.ColumnName
	}
	var b strings.Builder
	b.WriteString(funcName)
	b.WriteString("(")
	b.WriteString(f.ColumnName)
	b.WriteString(") AS ")
	b.WriteString(alias)
	return b.String()
}

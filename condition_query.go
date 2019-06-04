package somesql

import (
	"fmt"
)

var (
	andInQuery    = andOrInQuery(AndCondition, "IN")
	orInQuery     = andOrInQuery(OrCondition, "IN")
	andNotInQuery = andOrInQuery(AndCondition, "NOT IN")
	orNotInQuery  = andOrInQuery(OrCondition, "NOT IN")
)

// ConditionQuery defines conditions for a query
type ConditionQuery struct {
	Type     uint8
	Field    string
	Operator string
	Query    Accessor
	Lang     string
}

func andOrInQuery(conditionType uint8, operator string) func(lang, field string, query Accessor) ConditionQuery {
	return func(lang, field string, query Accessor) ConditionQuery {
		return ConditionQuery{
			Type:     conditionType,
			Field:    field,
			Operator: operator,
			Query:    query,
			Lang:     lang,
		}
	}
}

// AndInQuery returns a condition in the format IN(?,?,?) adjoined with AND
func AndInQuery(lang, field string, query Accessor) ConditionQuery {
	return andInQuery(lang, field, query)
}

// OrInQuery returns a condition in the format IN(?,?,?) adjoined with OR
func OrInQuery(lang, field string, query Accessor) ConditionQuery {
	return orInQuery(lang, field, query)
}

// AndNotInQuery returns a condition in the format NOT IN(?,?,?) adjoined with AND
func AndNotInQuery(lang, field string, query Accessor) ConditionQuery {
	return andNotInQuery(lang, field, query)
}

// OrNotInQuery returns a condition in the format NOT IN(?,?,?) adjoined with OR
func OrNotInQuery(lang, field string, query Accessor) ConditionQuery {
	return orNotInQuery(lang, field, query)
}

// ConditionType return the condition type (or / and)
func (c ConditionQuery) ConditionType() uint8 {
	return c.Type
}

// AsSQL returns part of SQL incuding the sub-query
func (c ConditionQuery) AsSQL(in ...bool) (string, []interface{}) {
	var (
		field string
	)

	if IsFieldMeta(c.Field) || IsWholeFieldData(c.Field) {
		field = fmt.Sprintf(`"%s"`, c.Field)
	} else {
		field = fmt.Sprintf(`"%s"->>'%s'`, GetLangFieldData(c.Lang), c.Field)
	}

	c.Query.ToSQL()

	innerSQL, innerVals := c.Query.GetSQL(), c.Query.GetValues()

	sql := fmt.Sprintf(`%s %s (%s)`, field, c.Operator, innerSQL)

	return sql, innerVals
}

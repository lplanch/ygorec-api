package util

import (
	"fmt"

	"gorm.io/gorm/clause"
)

type OrderByCase struct {
	Column clause.Column
	Values map[string]int
	Desc   bool
	Asc    bool
}

// Name where clause name
func (orderBy OrderByCase) Name() string {
	return "ORDER BY"
}

// Build build where clause
func (orderByCase OrderByCase) Build(builder clause.Builder) {
	builder.WriteString("CASE")
	for field, weight := range orderByCase.Values {
		builder.WriteString(" WHEN ")
		builder.WriteQuoted(orderByCase.Column)
		builder.WriteString(fmt.Sprintf(" LIKE '%s' THEN %d", field, weight)) // potential sql injection risk
	}
	builder.WriteString(" END")
	if orderByCase.Desc {
		builder.WriteString(" DESC")
	} else if orderByCase.Asc {
		builder.WriteString(" ASC")
	}
	builder.WriteString(", ")
	builder.WriteQuoted(orderByCase.Column)
	if orderByCase.Desc {
		builder.WriteString(" DESC")
	} else if orderByCase.Asc {
		builder.WriteString(" ASC")
	}
}

// MergeClause merge order by clauses
func (orderByCase OrderByCase) MergeClause(clause *clause.Clause) {
	clause.Expression = orderByCase
}

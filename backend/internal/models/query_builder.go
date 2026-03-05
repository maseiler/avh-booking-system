package models

import "strconv"

func buildSelectSQL(q *Query) string {
	stmt := "SELECT * FROM " + string(q.Table)

	if len(q.Filter) > 0 {
		stmt += " WHERE "
		for i, f := range q.Filter {
			stmt += f.Column + " " + f.Operator.SqlString() + " " + f.Value
			if i != len(q.Filter)-1 {
				stmt += " AND "
			}
		}
	}

	if q.Sort != nil {
		stmt += " ORDER BY " + q.Sort.Column + " " + q.Sort.Order.SqlString()
	}

	if q.Limit != nil {
		stmt += " LIMIT " + strconv.Itoa(*q.Limit)
	}

	return stmt
}

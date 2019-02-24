package mingo

// Object model
type Object map[string]interface{}

// Query model
type Query struct {
	Criteria Object
	compiled []func(Object) bool
}

// Test method evaluates the query by processing the expression's operators.
func (q *Query) Test(obj Object) bool {
	q.compile()

	for _, v := range q.compiled {
		if !v(obj) {
			return false
		}
	}
	return true
}

// compile method takes the query's criteria, breaks it down and processes
// the supported operators.
func (q *Query) compile() {
	if len(q.Criteria) == 0 {
		return
	}

	var whereOperator Object

	for k, v := range q.Criteria {
		if k == "$where" {
			whereOperator = Object{
				"field": k,
				"expr":  v,
			}
		} else if k == "$expr" {
			q.processOperator(k, k, v)
		} else if q.operatorInList(k, []string{"$and", "$or", "$nor"}) {
			q.processOperator(k, k, v)
		} else {
			expr := q.normalize(v)
			for k2, v2 := range expr {
				q.processOperator(k, k2, v2)
			}
		}

		if len(whereOperator) > 0 {
			q.processOperator(whereOperator["field"].(string), whereOperator["field"].(string), whereOperator["expr"].(Object))
		}
	}
}

// operatorInList method checks if the given operator is in the given slice
// of operators.
func (q *Query) operatorInList(operator string, list []string) bool {
	for _, v := range list {
		if v == operator {
			return true
		}
	}
	return false
}

// processOperator method applies the respective query operator on the sub
// expression.
func (q *Query) processOperator(field string, operator string, expr interface{}) {
	qo := QueryOperator{}

	switch operator {
	case "$eq":
		q.compiled = append(q.compiled, qo.eq(field, expr))
		break
	case "$ne":
		q.compiled = append(q.compiled, qo.ne(field, expr))
		break
	case "$gt":
		q.compiled = append(q.compiled, qo.gt(field, expr))
		break
	case "$gte":
		q.compiled = append(q.compiled, qo.gte(field, expr))
		break
	case "$lt":
		q.compiled = append(q.compiled, qo.lt(field, expr))
		break
	case "$lte":
		q.compiled = append(q.compiled, qo.lte(field, expr))
		break
	case "$and":
		switch expr.(type) {
		case []Object:
			q.compiled = append(q.compiled, qo.and(field, expr.([]Object)))
		}
		break
	}
}

// normalize method flattens down object expression values. Defaults to the
// equal operator.
func (q *Query) normalize(expr interface{}) Object {
	switch expr.(type) {
	// Object
	case Object:
		for k, v := range expr.(Object) {
			for _, v2 := range objectOperators {
				if k == v2 {
					return Object{v2: v}
				}
			}
		}
		return Object{
			"$eq": expr,
		}
	// Value
	default:
		return Object{
			"$eq": expr,
		}
	}
}

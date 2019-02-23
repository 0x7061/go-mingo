package mingo

// Object model
type Object map[string]interface{}

// Query model
type Query struct {
	Criteria Object
	compiled []func(Object) bool
}

// Compile method
func (q *Query) Compile() {
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

func (q *Query) operatorInList(operator string, list []string) bool {
	for _, v := range list {
		if v == operator {
			return true
		}
	}
	return false
}

// Test method
func (q *Query) Test(obj Object) bool {
	q.Compile()

	for _, v := range q.compiled {
		if !v(obj) {
			return false
		}
	}
	return true
}

func (q *Query) processOperator(field string, operator string, expr interface{}) {
	qo := QueryOperators{}

	switch operator {
	case "$and":
		q.compiled = append(q.compiled, qo.and(field, expr.([]Object)))
		break
	case "$eq":
		q.compiled = append(q.compiled, qo.eq(field, expr))
		break
	}
}

func (q *Query) normalize(expr interface{}) Object {
	// Primitive
	// TODO: Object expression
	return Object{
		"$eq": expr,
	}
}

// QueryOperators model
type QueryOperators struct{}

func (qo *QueryOperators) and(selector string, values []Object) func(Object) bool {
	var queries []Query

	for _, v := range values {
		q := Query{Criteria: v}
		q.Compile()
		queries = append(queries, q)
	}

	return func(obj Object) bool {
		for _, v := range queries {
			if !v.Test(obj) {
				return false
			}
		}
		return true
	}
}

func (qo *QueryOperators) eq(field string, expr interface{}) func(Object) bool {
	return func(obj Object) bool {
		if obj[field] == expr {
			return true
		}
		return false
	}
}

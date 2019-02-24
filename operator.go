package mingo

import (
	"reflect"
)

var (
	objectOperators = []string{
		"$ne",
		"$gt",
		"$gte",
		"$lt",
		"$lte",
	}
)

// QueryOperator model
type QueryOperator struct{}

// eq checks if the field of the given data matches the defined expression.
func (qo *QueryOperator) eq(field string, expr interface{}) func(Object) bool {
	return func(obj Object) bool {
		return obj[field] == expr
	}
}

// ne checks if the value of the given data doesn't match the expression.
func (qo *QueryOperator) ne(field string, expr interface{}) func(Object) bool {
	return func(obj Object) bool {
		return obj[field] != expr
	}
}

// gt checks if the value of the given data is greater than the value of the
// expression.
func (qo *QueryOperator) gt(field string, expr interface{}) func(Object) bool {
	return func(obj Object) bool {
		switch obj[field].(type) {
		case int:
			return reflect.TypeOf(obj[field]) == reflect.TypeOf(expr) && obj[field].(int) > expr.(int)
		case float32:
			return reflect.TypeOf(obj[field]) == reflect.TypeOf(expr) && obj[field].(float32) > expr.(float32)
		}
		return false
	}
}

// gte checks if the value of the given data is greater or equal than the
// value of the expression.
func (qo *QueryOperator) gte(field string, expr interface{}) func(Object) bool {
	return func(obj Object) bool {
		switch obj[field].(type) {
		case int:
			return reflect.TypeOf(obj[field]) == reflect.TypeOf(expr) && obj[field].(int) >= expr.(int)
		case float32:
			return reflect.TypeOf(obj[field]) == reflect.TypeOf(expr) && obj[field].(float32) >= expr.(float32)
		}
		return false
	}
}

// lt checks if the value of the given data is lower than the value of the
// expression.
func (qo *QueryOperator) lt(field string, expr interface{}) func(Object) bool {
	return func(obj Object) bool {
		switch obj[field].(type) {
		case int:
			return reflect.TypeOf(obj[field]) == reflect.TypeOf(expr) && obj[field].(int) < expr.(int)
		case float32:
			return reflect.TypeOf(obj[field]) == reflect.TypeOf(expr) && obj[field].(float32) < expr.(float32)
		}
		return false
	}
}

// lte checks if the value of the given data is lower or equal than the
// value of the expression.
func (qo *QueryOperator) lte(field string, expr interface{}) func(Object) bool {
	return func(obj Object) bool {
		switch obj[field].(type) {
		case int:
			return reflect.TypeOf(obj[field]) == reflect.TypeOf(expr) && obj[field].(int) <= expr.(int)
		case float32:
			return reflect.TypeOf(obj[field]) == reflect.TypeOf(expr) && obj[field].(float32) <= expr.(float32)
		}
		return false
	}
}

// and joins query clauses with a logical AND and returns all documents that
// match the conditions of both clauses.
func (qo *QueryOperator) and(selector string, values []Object) func(Object) bool {
	var queries []Query

	for _, v := range values {
		q := Query{Criteria: v}
		q.compile()
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

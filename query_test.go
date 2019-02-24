package mingo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOperatorInList(t *testing.T) {
	q := Query{}

	Convey("Given a list of operators", t, func() {
		o := []string{"$and", "$or", "$eq"}

		Convey("Where the operator is included", func() {
			So(q.operatorInList("$and", o), ShouldEqual, true)
		})
		Convey("Where the operator is not included", func() {
			So(q.operatorInList("$xor", o), ShouldEqual, false)
		})
	})
}

func TestNormalize(t *testing.T) {
	q := Query{}

	Convey("Given an expression", t, func() {
		Convey("That is value based", func() {
			expr := 10
			for k, v := range q.normalize(expr) {
				So(k, ShouldEqual, "$eq")
				So(v, ShouldEqual, 10)
			}
		})
		Convey("That is object based", func() {
			expr := Object{
				"$gt": 10,
			}
			for k, v := range q.normalize(expr) {
				So(k, ShouldEqual, "$gt")
				So(v, ShouldEqual, 10)
			}
		})
	})
}

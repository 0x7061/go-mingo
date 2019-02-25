package mingo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueryOperatorEq(t *testing.T) {
	Convey("Given an expression containing a comparison", t, func() {
		query := Query{Criteria: Object{
			"type.ranking": 10,
		}}
		Convey("Where the dataset matches the expression", func() {
			data := Object{
				"type": Object{
					"ranking": 10,
				},
			}
			So(query.Test(data), ShouldEqual, true)
		})
		Convey("Where the dataset does not match the expression", func() {
			data := Object{
				"type": "something",
			}
			So(query.Test(data), ShouldEqual, false)
		})
	})
}

func TestQueryOperatorNe(t *testing.T) {
	Convey("Given an expression containing a comparison", t, func() {
		query := Query{Criteria: Object{
			"type": Object{
				"$ne": "ranking",
			},
		}}
		Convey("Where the dataset matches the expression", func() {
			data := Object{
				"type": "something",
			}
			So(query.Test(data), ShouldEqual, true)
		})
		Convey("Where the dataset does not match the expression", func() {
			data := Object{
				"type": "ranking",
			}
			So(query.Test(data), ShouldEqual, false)
		})
	})
}

func TestQueryOperatorGt(t *testing.T) {
	Convey("Given an expression containing a comparison", t, func() {
		query := Query{Criteria: Object{
			"type": Object{
				"$gt": 20,
			},
		}}
		Convey("Where the dataset matches the expression", func() {
			data := Object{
				"type": 21,
			}
			So(query.Test(data), ShouldEqual, true)
		})
		Convey("Where the dataset does not match the expression", func() {
			data := Object{
				"type": 20,
			}
			So(query.Test(data), ShouldEqual, false)
		})
	})
}

func TestQueryOperatorGte(t *testing.T) {
	Convey("Given an expression containing a comparison", t, func() {
		query := Query{Criteria: Object{
			"type": Object{
				"$gte": 20,
			},
		}}
		Convey("Where the dataset matches the expression", func() {
			data := Object{
				"type": 20,
			}
			So(query.Test(data), ShouldEqual, true)
		})
		Convey("Where the dataset does not match the expression", func() {
			data := Object{
				"type": 19,
			}
			So(query.Test(data), ShouldEqual, false)
		})
	})
}

func TestQueryOperatorLt(t *testing.T) {
	Convey("Given an expression containing a comparison", t, func() {
		query := Query{Criteria: Object{
			"type": Object{
				"$lt": 20,
			},
		}}
		Convey("Where the dataset matches the expression", func() {
			data := Object{
				"type": 19,
			}
			So(query.Test(data), ShouldEqual, true)
		})
		Convey("Where the dataset does not match the expression", func() {
			data := Object{
				"type": 20,
			}
			So(query.Test(data), ShouldEqual, false)
		})
	})
}

func TestQueryOperatorLte(t *testing.T) {
	Convey("Given an expression containing a comparison", t, func() {
		query := Query{Criteria: Object{
			"type": Object{
				"$lte": 20,
			},
		}}
		Convey("Where the dataset matches the expression", func() {
			data := Object{
				"type": 20,
			}
			So(query.Test(data), ShouldEqual, true)
		})
		Convey("Where the dataset does not match the expression", func() {
			data := Object{
				"type": 21,
			}
			So(query.Test(data), ShouldEqual, false)
		})
	})
}

func TestQueryOperatorAnd(t *testing.T) {
	Convey("Given an expression containing an AND operator", t, func() {
		query := Query{Criteria: Object{
			"type": "ranking",
			"$and": []Object{
				Object{
					"score": Object{
						"$gt": 5,
					},
				},
			},
		}}
		Convey("Where the dataset matches the expression", func() {
			data := Object{
				"type":  "ranking",
				"score": 10,
			}
			So(query.Test(data), ShouldEqual, true)
		})
		Convey("Where the dataset does not match the expression", func() {
			data := Object{
				"type":  "ranking",
				"score": 5,
			}
			So(query.Test(data), ShouldEqual, false)
		})
	})
}

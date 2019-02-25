package mingo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResolve(t *testing.T) {
	Convey("Given a valid object", t, func() {
		obj := Object{
			"foo": Object{
				"bar": Object{
					"moo": "noo",
				},
			},
		}

		Convey("And a valid selector", func() {
			sel := "foo.bar.moo"

			Convey("Running resolve()", func() {
				Convey("Should return the desired value", func() {
					value := resolve(obj, sel)
					So(value, ShouldEqual, "noo")
				})
			})
		})
		Convey("And an invalid selector", func() {
			sel := "x.y.z"

			Convey("Running resolve()", func() {
				Convey("Should return nil", func() {
					value := resolve(obj, sel)
					So(value, ShouldEqual, nil)
				})
			})
		})
	})
}

package main

import (
	"go-mingo/mingo"
	"log"
)

func main() {
	query := mingo.Query{Criteria: mingo.Object{
		"type": "homework",
		"$and": []mingo.Object{
			mingo.Object{
				"score": 10,
			},
		},
	}}

	result := query.Test(mingo.Object{
		"type":  "homework",
		"score": 10,
	})

	log.Println(result)
}

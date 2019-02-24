# go-mingo
Minimal mingo port for Go projects.

## Install
```$ go get github.com/codepushr/go-mingo```

## Features
- Query Operators
  - $eq, $gt, $gte, $lt, $lte, $and

## Usage
Import go-mingo, create a query and run `Test()` on it.
```
import "github.com/codepushr/go-mingo"

func main() {
    q := mingo.Query{Criteria: Object{
        "type": "ranking",
        "$and": []mingo.Object{
            mingo.Object{
                "score": mingo.Object{
                    "$gt": 5,
                },
            },
        },
    }} 

    result := q.Test(mingo.Object{
        "type":  "ranking",
        "score": 10,
    })
}
```

## Roadmap
- Support Dot Notation for both _`<array>.<index>`_ and _`<document>.<field>`_ selectors

## License
MIT
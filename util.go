package mingo

import (
	"strings"
)

// resolve method takes an object and selector and traverses the object by
// the selector to return the final value.
func resolve(obj Object, selector string) interface{} {
	path := strings.Split(selector, ".")

	if len(path) > 1 {
		if obj[path[0]] != nil {
			newObj := Object{}

			switch obj[path[0]].(type) {
			case map[string]interface{}:
				for k, v := range obj[path[0]].(map[string]interface{}) {
					newObj[k] = v
				}
				path = append(path[:0], path[0+1:]...)
				return resolve(newObj, strings.Join(path, "."))
			case Object:
				newObj = obj[path[0]].(Object)
				path = append(path[:0], path[0+1:]...)
				return resolve(newObj, strings.Join(path, "."))
			}
			return nil
		}
		return nil
	}

	return obj[path[0]]
}

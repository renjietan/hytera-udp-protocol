package tools

import "reflect"

// Tern Mortal 2026/5/18 16:22 三元表达式
func Tern[T any, K any](boolVal bool, a T, b K) any {
	var result any
	if boolVal {
		result = a
	} else {
		result = b
	}
	val := reflect.ValueOf(result)
	if val.Kind() == reflect.Func {
		funcType := val.Type()
		numIn := funcType.NumIn()
		if numIn == 0 {
			results := val.Call(nil)
			if len(results) > 0 {
				return results[0].Interface()
			}
			return nil
		}
		return result
	}

	return result
}

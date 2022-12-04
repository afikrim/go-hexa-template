package replacerutil

import "fmt"

func Replace(target, value interface{}) interface{} {
	if value == nil {
		return target
	}
	if getType(target) != getType(value) {
		return target
	}
	if target == value || isEmpty(value) {
		return target
	}

	return value
}

func isEmpty(target interface{}) bool {
	switch getType(target) {
	case "string":
		if target == "" {
			return true
		}
	case "int":
		if target == 0 {
			return true
		}
	}

	return false
}

func getType(target interface{}) string {
	return fmt.Sprintf("%T", target)
}

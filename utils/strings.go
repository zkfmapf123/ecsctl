package utils

import (
	"reflect"
)

func NotExistReturnDefault(pre string, def string) string {
	if pre == "" {
		return def
	}

	return pre
}

func IncludeString(arr []string, s string) bool {
	for _, v := range arr {

		if v == s {
			return true
		}
	}

	return false
}

func ToMap[T any](object T) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.ValueOf(object)
	typ := reflect.TypeOf(object)

	if val.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}

package common

import (
	"reflect"
)

func CheckFieldExistsAndNotZero(v interface{}, fieldName string) (exists bool, isZero bool) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return false, true
	}
	return true, field.IsZero()
}

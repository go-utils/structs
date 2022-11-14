package structs

import "reflect"

// GetStructName - get struct name
func GetStructName(obj any) string {
	rt := reflect.TypeOf(obj)

	// NOTE: guard
	if rt == nil {
		return ""
	}

	switch rt.Kind() {
	case reflect.Ptr, reflect.Struct:
	default:
		// NOTE: guard
		return ""
	}

	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	return rt.Name()
}

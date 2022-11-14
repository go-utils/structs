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

// GetNilFields - get nil field names
func GetNilFields(obj any) []string {
	return getNilFields(obj, "")
}

func getNilFields(value any, parent string) []string {
	nilFields := make([]string, 0)

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Struct:
		// ok
	case reflect.Ptr:
		if rv.IsNil() {
			return append(nilFields, parent)
		}
	default:
		return nilFields
	}

	rv = reflect.Indirect(rv)
	rt := rv.Type()
	structName := rt.Name()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		fieldName := f.Name

		field := rv.FieldByName(fieldName)
		if !field.IsValid() {
			continue
		}

		key := structName + "." + fieldName
		if parent != "" {
			key = parent + "." + key
		}

		switch field.Kind() {
		case reflect.Ptr, reflect.Interface:
			if field.IsNil() {
				nilFields = append(nilFields, key)
				continue
			}
		}

		if f.Anonymous {
			nilFields = append(nilFields, getNilFields(field.Interface(), structName)...)
		}

		continue
	}

	return nilFields
}

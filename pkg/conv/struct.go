package conv

import (
	"fmt"
	"reflect"
	"regexp"
)

// StructToMap convert struct to map, only convert exported struct field
// map key is specified same as struct field tag `json` value
func StructToMap(value any) (map[string]any, error) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("data type %T not support, shuld be struct or pointer to struct", value)
	}

	result := make(map[string]any)

	fieldNum := t.NumField()
	pattern := `^[A-Z]`
	regex := regexp.MustCompile(pattern)
	for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		tag := t.Field(i).Tag.Get("json")
		if regex.MatchString(name) && tag != "" {
			//result[name] = v.Field(i).Interface()
			result[tag] = v.Field(i).Interface()
		}
	}

	return result, nil
}

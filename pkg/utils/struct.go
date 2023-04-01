package utils

import (
	"log"
	"reflect"

	"golang.org/x/exp/maps"
)

func ExtractTagMap(tag string, st interface{}) map[string]string {
	tagMap := map[string]string{}

	t := reflect.TypeOf(st)
	v := reflect.ValueOf(st)

	if t.Kind() == reflect.Pointer {
		if v.IsNil() {
			log.Fatalf("nil struct passed to ExtractTagMap")
		}
		v = v.Elem()
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct ||
			field.Type.Kind() == reflect.Pointer ||
			field.Type.Kind() == reflect.Interface {

			extractedMap := ExtractTagMap(tag, v.Field(i).Interface())
			maps.Copy(tagMap, extractedMap)
		}

		if tagVal, ok := field.Tag.Lookup(tag); ok {
			tagMap[tagVal] = v.Field(i).String()
		}
	}
	return tagMap
}

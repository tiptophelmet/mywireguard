package utils

import "regexp"

func StrCompose(template string, args map[string]string) string {
	r := regexp.MustCompile(`{(\w+)}`)
	return r.ReplaceAllStringFunc(template, func(m string) string {
		k := m[1 : len(m)-1]
		return args[k] // If key not in map, returns empty string.
	})
}

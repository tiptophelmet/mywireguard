package utils

type strmap map[string]string

func MergeStrMaps(fromMap strmap, toMap strmap) strmap {
	for k, v := range fromMap {
		toMap[k] = v
	}

	return toMap
}

package utils

type strmap map[string]string

func MergeStrMaps(map1 strmap, map2 strmap) strmap {
	for k, v := range map2 {
		map1[k] = v
	}

	return map1
}

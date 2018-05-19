package equals

func SliceEquals(res, exp []string) bool {
	if len(res) != len(exp) {
		return false
	}

	for i, _ := range exp {
		if res[i] != exp[i] {
			return false
		}
	}
	return true
}

func MapEquals(res, exp map[string]int) bool {
	for k := range exp {
		if res[k] != exp[k] {
			return false
		}
	}
	return true
}

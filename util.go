package logplus

func concat(sep string, strs ...string) string {
	var result = make([]byte, 0, 100)
	for i, _ := range strs {
		result = append(result, strs[i]...)
		if i < len(strs)-1 {
			result = append(result, sep...)
		}
	}
	return string(result)
}

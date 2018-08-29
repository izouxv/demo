package util

func IntToBool(i int32) (b bool) {
	if i == 0 {
		b = false
	} else {
		b = true
	}
	return b
}

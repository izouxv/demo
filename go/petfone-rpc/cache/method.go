package cache


func InterToMap(inter interface{}, flag bool) (map[int32]int32, bool) {
	return inter.(map[int32]int32), flag
}

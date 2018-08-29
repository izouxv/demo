package util

type Comparable interface {
	Compare(o interface{}) bool // true -> 大于，false > 小于等于
}

type SortSet struct {
	key   string
	value int64
}

func (this *SortSet) Compare(o interface{}) bool {
	s := o.(*SortSet)
	if s.value > this.value {
		return true
	}
	return false
}

func QuickSort(src []Comparable, first, last int) {
	flag := first
	left := first
	right := last

	if first >= last {
		return
	}

	for first < last {
		for first < last {
			if src[last].Compare(src[flag]) {
				last -= 1
				continue
			} else {
				tmp := src[last]
				src[last] = src[flag]
				src[flag] = tmp
				flag = last
				break
			}
		}

		for first < last {
			if !src[first].Compare(src[flag]) {
				first += 1
				continue
			} else {
				tmp := src[first]
				src[first] = src[flag]
				src[flag] = tmp
				flag = first
				break
			}
		}
	}
	QuickSort(src, left, flag-1)
	QuickSort(src, flag+1, right)
}

package util

type Pair struct {
	Key   string
	Value int64
}

type PairList []*Pair

func SortMapByValue(m map[string]int64, desc bool) PairList {
	p := make(PairList, 0)
	i := 0
	for k, v := range m {
		p = append(p, &Pair{k, v})
	}
	for i = 1; i < len(m); i++ {
		j := i
		for j > 0 {
			if desc {
				/*true 降序*/
				if p[j-1].Value < p[j].Value {
					p[j-1], p[j] = p[j], p[j-1]
				}
				j = j - 1
			} else {
				if p[j-1].Value > p[j].Value {
					p[j-1], p[j] = p[j], p[j-1]
				}
				j = j - 1
			}
		}
	}
	return p
}

package api

import (
	"testing"
	"fmt"
)

func TestUserServer_FindPassword(t *testing.T) {
	type test struct {
		Name string
	}
	m := make(map[string]*test)
	te := []test{
		{Name: "test01"}, {Name: "test02"}, {Name: "test03"},
	}
	aa := test{}
	for i := 0;i<len(te);i++ {
		fmt.Println(te[i])
		aa = te[i]
		m[te[i].Name] = &aa
	}
	fmt.Println(m)
	for k,v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
}
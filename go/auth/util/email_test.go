package util

import (
	"testing"
	"regexp"
	"fmt"
)

func TestSendMail(t *testing.T) {
	SendMail("lids@radacat.com", []string{"token:" + "11111111", "nickname:" + "lids"}, "61")
}

func TestURL(t *testing.T)  {

	match, err := regexp.Match(`^/v1.1/tenants/[0-9]+$`, []byte("/v1.1/tenants/1/roles/2"))
	fmt.Println(match)
	fmt.Println(err)
}
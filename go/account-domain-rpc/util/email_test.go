package util

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	SendMail("lids@radacat.com", []string{"token:" + "11111111", "nickname:" + "lids"}, RegisterMail)
}

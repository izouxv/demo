package util

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

type Session struct {
	Uid      int32
	Username string
}

func (this *Session) GenerateSession() string {
	s := strconv.Itoa(int(this.Uid))+this.Username + strconv.FormatUint(uint64(time.Now().Unix()), 10)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

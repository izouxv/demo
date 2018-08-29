package util

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
)

func GenerateSession(base int32) string {
	s := strconv.Itoa(int(base)) + strconv.Itoa(int(rand.Int31()))
	session := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", session)
}

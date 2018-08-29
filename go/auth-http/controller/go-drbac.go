package controller

import "net/http"

var (
	Dids           = make(map[int32]int64)
	DidCookieKey   = "did"
	TokenCookieKey = "token"
	UserContext    = "UserContext"
)

type UserInfoContext struct {
	Uid      int64
	Username string
	Did      int64
	Token    string
	NickName string
}

/*----------getCookie-----------*/
func GetCookie(r *http.Request, key string) string {
	cookie, err := r.Cookie(key)
	if err != nil {
		return ""
	}
	return cookie.Value
}

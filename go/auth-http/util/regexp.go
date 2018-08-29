package util

import "regexp"

var (
	NameRegexp = regexp.MustCompile("^[A-Za-z0-9\u4e00-\u9fa5]{1,30}$")
	//UsernameRegexp = regexp.MustCompile("^[a-zA-Z0-9]*@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9-_]{2,5})+$")
	UsernameRegexp = regexp.MustCompile("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*.\\w+([-.]\\w+)*$")
	NicknameRegexp = regexp.MustCompile("^[A-Za-z0-9\u4e00-\u9fa5]{1,10}$")
	MacRegexp = regexp.MustCompile("^[A-Fa-f0-9]{2}:[A-Fa-f0-9]{2}:[A-Fa-f0-9]{2}:[A-Fa-f0-9]{2}:[A-Fa-f0-9]{2}:[A-Fa-f0-9]{2}$")
	VersionoReq = regexp.MustCompile("^\\d+([.][0-9]+){2}$")
)
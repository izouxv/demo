package po

import (
	"cotx-http/pb"
)

//登录：返回用户信息
type LoginPo struct {
	//账户基本资料
	Uid			int32		`json:"uid"`
	Username	string		`json:"username"`
	Nickname	string		`json:"nickname"`
	Token		string		`json:"token"`
	LoginState	int32		`json:"loginState"`
	State		int32		`json:"state"`
	ErrorCode	int32		`json:"errorCode"`
	//账户拓展资料
	Phone		string		`json:"phone"`
	Email		string		`json:"email"`
	Gender		int32		`json:"gender"`
	Birthday	int64		`json:"birthday"`
	Avatar		int32		`json:"avatar"`
	Signature	string		`json:"signature"`
	Address		string		`json:"address"`
	Job			int32		`json:"job"`
	//账户认证
	Realname	string		`json:"realname"`
	IdCard		string		`json:"idCard"`
	Certify		int32		`json:"certify"`
	//账户
	Credit		int32		`json:"credit"`
	Point		int32		`json:"point"`
	Grade		int32		`json:"grade"`
}

//对login/sso进行属性处理
func (l *LoginPo) SetSsoReplyIntoLoginPo(ssoReply *pb.SsoReply)  {
	l.Uid		= ssoReply.Uid
	l.Username	= ssoReply.Username
	l.Nickname	= ssoReply.Nickname
	l.Token		= ssoReply.SessionName
	l.LoginState= ssoReply.LoginState
	l.State		= ssoReply.State
	l.ErrorCode	= ssoReply.ErrorCode
}

//对login/account进行属性处理
func (l *LoginPo) SetAccountReplyIntoLoginPo(accountReply *pb.AccountReply)  {
	l.Phone		= accountReply.Phone
	l.Email		= accountReply.Email
	l.Gender	= accountReply.Gender
	l.Birthday	= accountReply.Birthday
	l.Avatar	= accountReply.Avatar
	l.Signature	= accountReply.Signature
	l.Address	= accountReply.UserAddress
	l.Job		= accountReply.UserJobId
	l.Realname	= accountReply.Realname
	l.IdCard	= accountReply.IdentityCard
	l.Certify	= accountReply.IsCertification
	l.Credit	= accountReply.CreditValues
	l.Point		= accountReply.UserPoint
	l.Grade		= accountReply.UserGradeId
}

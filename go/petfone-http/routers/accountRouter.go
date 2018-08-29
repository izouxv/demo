package routers

import (
	"github.com/julienschmidt/httprouter"
	"petfone-http/controller"
)

func accountRouter(router *httprouter.Router) {
	//todo 半公开接口
	//手机验证码
	router.POST(uriMobile, authInterceptor(controller.MobileCode))
	//邮箱验证
	//注册
	router.POST(uriReg, authInterceptor(controller.Register))
	//登录
	router.POST("/", controller.Index)
	router.POST(uriSessions, authInterceptor(controller.Login))
	//手机重置密码
	router.POST(uriResetPwd, authInterceptor(controller.ResetPassword))
	//邮箱重置密码
	router.POST(uriResetMail, authInterceptor(controller.MailResetPassword))

	//todo 封闭接口
	router.DELETE(uriSessions, authInterceptor(controller.Logout))
	//用户信息
	router.PUT(uriUserinfo, authInterceptor(controller.UpdateUserInfo))
	router.GET(uriUserinfo, authInterceptor(controller.GetUserInfo))
	router.GET(uriAccounts, authInterceptor(controller.GetAccountByUserName))
	//修改密码
	router.PUT(uriUpdatePwd, authInterceptor(controller.UserPassword))
	//围栏
	//router.POST(uriFenceDy, authInterceptor(controller.SetPetfoneFence))
	router.PUT(uriFenceDy, authInterceptor(controller.UpdatePetfoneFence))
	router.GET(uriFenceDy, authInterceptor(controller.GetPetfoneFence))
	//通知
	router.POST(uriNotices, authInterceptor(controller.SetNotice))
	router.GET(uriNotices, authInterceptor(controller.GetNotice))
	router.PUT(uriNotice, authInterceptor(controller.UpdateNotice))
}



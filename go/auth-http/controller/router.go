package controller

import (
	"github.com/julienschmidt/httprouter"
)

func RouterMethod(router *httprouter.Router) {
	//todo Login
	//router.POST(  "/v1.1/tokens",                         							auth.LoginWithDid)			//登录域
	router.POST("/v1.1/tokens", Login)      //登录租户
	router.DELETE("/v1.1/tokens", LoginOut) //登出
	router.GET("/v1.1/tokens", GetUserInfo) //获取用户信息

	//todo 账号信息
	router.PUT("/v1.1/accounts/:accountId", UpdateUserInfoAndPassword) //修改用户信息（昵称和密码）
	router.PUT("/v1.1/accounts/:accountId/password", UpdatePassword)   //修改密码
	router.POST("/v1.1/message/password", FindPassword)                //忘记密码
	router.POST("/v1.1/message/password/:token", ResetPassword)        //重置密码
	router.GET("/v1.1/message/tenants/:tid", GetDidByTid)    	      //通过tid获取did
	router.PUT("/v1.1/accounts/:accountId/state", UpdateUserState)     //修改用户状态

	//todo 租户
	router.POST("/v1.1/domains/:did/utacl", AddUserTenantACL)            //添加用户租户acl关系
	router.POST("/v1.1/domains/:did/tenants", AddTenantByDomain)            //添加租户
	router.POST("/v1.1/domains/:did/tenants/tid", AddTenantByFather)        //添加子租户
	router.PUT("/v1.1/domains/:did/tenants/:tid", UpdateTenant)             //修改租户信息
	router.PUT("/v1.1/domains/:did/tenants/:tid/states", UpdateTenantState) //修改租户状态
	router.GET("/v1.1/domains/:did/tenants", GetTenants)                    //获取域下租户列表
	router.DELETE("/v1.1/domains/:did/tenants/:tid", DeleteTenant)          //删除指定租户
	router.POST("/v1.2/logos/tenants", GetTidByUrl)    					//通过url获取tid
	router.POST("/v1.1/message/domains/:did/tenants", InviteUnactivatedTenant) 		//再次邀请租户

	//todo 租户角色
	router.GET("/v1.1/tenants/:tid/modules", GetTenantModules)       //获取租户权限模块列表
	router.POST("/v1.1/tenants/:tid/roles", AddTenantRole)           //添加租户角色
	router.PUT("/v1.1/tenants/:tid/roles", UpdateTenantRole)         //修改租户角色信息
	router.GET("/v1.1/tenants/:tid/roles", GetTenantRoles)           //获取租户角色列表
	router.DELETE("/v1.1/tenants/:tid/roles/:rid", DeleteTenantRole) //删除租户角色
	router.GET("/v1.1/tenants/:tid/roles/:rid", GetTenantRoleByRid)  //通过rid获取租户角色信息

	//todo 域角色
	router.GET("/v1.1/domains/:did/modules", GetDomainModules)       //获取域权限模块列表
	router.POST("/v1.1/domains/:did/roles", AddDomainRole)           //添加域角色
	router.PUT("/v1.1/domains/:did/roles", UpdateDomainRole)         //修改域角色信息
	router.GET("/v1.1/domains/:did/roles", GetDomainRoles)           //获取域角色列表
	router.DELETE("/v1.1/domains/:did/roles/:rid", DeleteDomainRole) //删除域角色
	router.GET("/v1.1/domains/:did/roles/:rid", GetDomainRoleByRid)  //通过rid获取域角色信息

	//todo 租户用户
	router.POST("/v1.1/tenants/:tid/users", AddTenantUser)                    //邀请用户进入租户
	router.PUT("/v1.1/tenants", EnterTenant)                                  //同意进入租户
	router.PUT("/v1.1/tenants/:tid/users/:uid/roles", UpdateTenantUserRole)   //修改租户用户角色
	router.PUT("/v1.1/tenants/:tid/users/:uid/states", UpdateTenantUserState) //修改租户用户状态
	router.GET("/v1.1/tenants/:tid/users", GetTenantUsers)                    //获取租户用户列表
	router.DELETE("/v1.1/tenants/:tid/users/:uid", DeleteTenantUser)          //删除租户用户


	//todo 域用户
	router.POST("/v1.1/domains/:did/users", AddDomainUser)                    //邀请用户进入域
	router.PUT("/v1.1/domains", EnterDomain)                                  //同意进入域
	router.PUT("/v1.1/domains/:did/users/:uid/roles", UpdateDomainUserRole)   //修改域用户角色
	router.PUT("/v1.1/domains/:did/users/:uid/states", UpdateDomainUserState) //修改域用户状态
	router.GET("/v1.1/domains/:did/users", GetDomainUsers)                    //获取域用户列表
	router.DELETE("/v1.1/domains/:did/users/:uid", DeleteDomainUser)          //删除域用户

	//todo 日志
	router.GET("/v1.1/actionlogs",							        				GetActionLogs)
	router.GET("/v1.1/domains/:did/actionlogs",							        	GetAllActionLogsByDid)
	router.GET("/v1.1/tenants/:tid/actionlogs",							        	GetAllActionLogsByTid)

	//todo 服务
	router.POST("/v1.1/tenants/:tid/services",AddService)						//添加服务
	router.DELETE("/v1.1/tenants/:tid/services/:sid",DeleteService)			//删除服务
	router.PUT("/v1.1/tenants/:tid/services/:sid",UpdateService)				//修改服务
	router.GET("/v1.1/tenants/:tid/services/:sid",GetServiceBySid)			//通过Sid获取服务
	router.GET("/v1.1/tenants/:tid/services",GetServiceByTid)					//通过Tid获取服务

	//todo 计费策略
	router.POST("/v1.1/domains/:did/services/:sid/policy",AddPolicy)					//添加计费策略
	router.DELETE("/v1.1/domains/:did/services/:sid/policy/:pid",DeletePolicyByPid)	//通过pid删除计费策略
	router.DELETE("/v1.1/domains/:did/services/:sid/policy",DeletePolicyBySid)			//通过sid删除计费策略
	router.PUT("/v1.1/domains/:did/services/:sid/policy/:pid",UpdatePolicy)			//修改计费策略
	router.GET("/v1.1/domains/:did/services/:sid/policy/:pid",GetPolicyByPid)			//通过Pid获取计费策略
	router.GET("/v1.1/domains/:did/services/:sid/policy",GetPolicyBySid)				//通过Sid获取计费策略

	// todo 支付信息
	router.POST("/v1.1/domains/:did/alipay",AddAliPay)				//添加支付宝信息
	router.DELETE("/v1.1/domains/:did/alipay",DeleteAliPay)			//删除支付宝信息
	router.PUT("/v1.1/domains/:did/alipay",UpdateAliPay)				//修改支付宝信息
	router.GET("/v1.1/domains/:did/alipay",GetAliPay)					//查询支付宝信息
	router.POST("/v1.1/domains/:did/wechatpay",AddWechatPay)			//添加微信支付信息
	router.DELETE("/v1.1/domains/:did/wechatpay",DeleteWechatPay)		//删除微信支付信息
	router.PUT("/v1.1/domains/:did/wechatpay",UpdateWechatPay)		//修改微信支付信息
	router.GET("/v1.1/domains/:did/wechatpay",GetWechatPay)			//查询微信支付信息

	//todo 租户账户金额及交易记录
	router.GET("/v1.1/tenants/:tid/account",GetTenantAccount)				//查询租户账户余额
	router.GET("/v1.1/tenants/:tid/tradings",GetTradingRecord)			//查询租户交易记录
	router.PUT("/v1.1/tenants/:tid/account",UpdateTenantAccount)			//修改租户交易记录（充值）

}

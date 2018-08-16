package controllers

import (
	"github.com/astaxie/beego"
)

// mainTest API
type MainController struct {
	beego.Controller
}

// @Title getStaticBlock
// @Description get all the staticblock by key
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /staticblock/:key [get]
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

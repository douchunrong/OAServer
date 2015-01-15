package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get() {
	self.TplNames = "login.html"
}

func (self *LoginController) Post() {
	userName := self.Input().Get("userName")
	userPwd := self.Input().Get("userpwd")
	if userName == beego.AppConfig.String("adminUserName") {
		if userPwd == beego.AppConfig.String("adminUserPwd") {
			self.Redirect("/home", 301)
		}
	}

	return
}

package controllers

import (
	"OAServer/identity_oauth2_server/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/context"
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
	autoLogin := self.Input().Get("autoLogin") == "on"

	beego.Debug("开始访问数据库。。。。")
	user := models.FindUserByName(userName)
	beego.Debug("结束访问。。。。")
	if user != nil {
		if user.Password == userPwd {
			maxAge := 0
			if autoLogin {
				maxAge = 1<<31 - 1
			}
			self.Ctx.SetCookie("userName", userName, maxAge, "/")
			self.Ctx.SetCookie("userPwd", userPwd, maxAge, "/")
			self.Ctx.SetCookie("IsLogin", "yes", "/")
			self.Redirect("/home", 301)
		} else {
			self.Redirect("/", 301)
		}
		self.Redirect("/", 301)
	}

	return
}

// 12114882015

// lihe@rntd.cn

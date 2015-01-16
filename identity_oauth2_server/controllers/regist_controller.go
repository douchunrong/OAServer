package controllers

import (
	"OAServer/identity_oauth2_server/models"
	"github.com/astaxie/beego"
)

type RegistController struct {
	beego.Controller
}

func (self *RegistController) Get() {
	self.TplNames = "regist.html"
}

func (self *RegistController) Post() {
	userName := self.Input().Get("userName")
	userPwd := self.Input().Get("userPwd")
	userRePwd := self.Input().Get("userRePwd")

	user := models.FindUserByName(userName)
	if user != nil {
		return
	} else {
		if userPwd == userRePwd {
			self.Redirect("/home", 301)
		}
	}
	return
}

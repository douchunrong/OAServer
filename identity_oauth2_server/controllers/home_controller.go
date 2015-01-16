package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (self *HomeController) Get() {
	self.TplNames = "home.html"
	self.Data["IsHome"] = true
	self.Data["IsLogin"] = self.Ctx.GetCookie("IsLogin") == "yes"
}

func (self *HomeController) Post() {

}

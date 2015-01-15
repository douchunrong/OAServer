package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (self *HomeController) Get() {
	self.TplNames = "home.html"
}

func (self *HomeController) Post() {

}

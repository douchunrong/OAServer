package main

import (
	_ "OAServer/identity_oauth2_server/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}


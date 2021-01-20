package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "gopescms/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true


	beego.Run()
}


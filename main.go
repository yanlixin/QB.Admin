package main

import (
	"encoding/gob"

	"QB.Admin/models/system"
	_ "QB.Admin/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	//go get github.com/garyburd/redigo/redis
	//"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq" // import your used driver
)

func init() {
	gob.Register(&models.User{})
}
func main() {
	beego.SetStaticPath("/html", "static/html")
	beego.AddTemplateExt("html")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379,100,qubei.2015"
	beego.BConfig.WebConfig.TemplateLeft = "<%"
	beego.BConfig.WebConfig.TemplateRight = "%>"

	beego.Run()
}

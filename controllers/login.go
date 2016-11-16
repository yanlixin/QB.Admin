package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"QB.Admin/models/system"
	"QB.Admin/utils"
	"github.com/astaxie/beego"
	"github.com/gpmgo/gopm/modules/log"
)

//LoginController is the login controller
type LoginController struct {
	beego.Controller
}

//Index of LoginController
func (ctr *LoginController) Index() {
	ctr.TplName = "login.html"
	ctr.Data["AppTitle"] = "DEMP"
	ctr.Data["cdn"] = utils.CdnURL
	ctr.Data["xsrfdata"] = template.HTML(ctr.XSRFFormHTML())
}

//Login of LoginController
func (ctr *LoginController) Login() {
	ctr.StartSession()
	ctr.TplName = "login.html"
	ctr.Data["cdn"] = utils.CdnURL
	loginName := ctr.GetString("loginName")
	loginPwd := ctr.GetString("userPassword")
	ctr.Data["xsrfdata"] = template.HTML(ctr.XSRFFormHTML())
	isSucceed, user := models.Login(loginName, loginPwd)
	flash := beego.NewFlash()
	if !isSucceed {
		flash.Error("登录名称或登录密码错误！")
		flash.Store(&ctr.Controller)

	} else {
		log.Debug(fmt.Sprintf("%s", "Logined"))

		ctr.SetSession("CurrentUser", user)
		ctr.Ctx.Redirect(302, "/index")
	}
}

//Logout of LoginController
func (ctr *LoginController) Logout() {
	cookie := http.Cookie{Name: "qubeiyuan", Path: "/", MaxAge: -1}
	http.SetCookie(ctr.Ctx.ResponseWriter, &cookie)
	ctr.Ctx.Redirect(302, "/login")
}

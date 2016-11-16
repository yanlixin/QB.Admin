package controllers

import (
	"html/template"
	"time"

	"QB.Admin/models/system"
	"QB.Admin/utils"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

// NestPreparer implemented global settings for all other routers.
type NestPreparer interface {
	NestPrepare()
}

// baseRouter implemented global settings for all other routers.
type baseRouter struct {
	beego.Controller
	i18n.Locale
	CurrentUser *models.User
	isLogin     bool
	flash       beego.FlashData
}

// Prepare implemented Prepare method for baseRouter.
func (router *baseRouter) Prepare() {
	// page start time
	var sessionUser = router.GetSession("CurrentUser")
	router.Data["xsrfdata"] = template.HTML(router.XSRFFormHTML())
	router.Data["xsrf_token"] = router.XSRFToken()
	router.Data["PageStartTime"] = time.Now()

	// Setting properties.
	router.Data["AppDescription"] = utils.AppDescription
	router.Data["AppKeywords"] = utils.AppKeywords
	router.Data["AppName"] = utils.AppName
	router.Data["AppVer"] = utils.AppVer
	router.Data["AppUrl"] = utils.AppURL
	router.Data["AppLogo"] = utils.AppLogo
	router.Data["AvatarURL"] = utils.AvatarURL
	router.Data["IsProMode"] = utils.IsProMode
	router.Data["cdn"] = utils.CdnURL
	router.Data["AppTitle"] = utils.AppTitle
	router.Data["UserInfo"] = sessionUser
	router.Data["HasAdd"] = false
	router.Data["HasEdit"] = false
	router.Data["HasDelete"] = false
	router.Data["HasView"] = false

	if nil != sessionUser {
		router.CurrentUser = sessionUser.(*models.User)
	}
	//fmt.Println(v)}
	if app, ok := router.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
func (router *baseRouter) CheckActiveRedirect() bool {
	return false
}
func (router *baseRouter) FlashWrite(msg string, isPr string) {

}

//BaseAdminRouter implemented global settings for all other routers.
type BaseAdminRouter struct {
	baseRouter
}

//NestPrepare of BaseAdminRouter implemented global settings for all other routers.
func (router *BaseAdminRouter) NestPrepare() {
	if router.CheckActiveRedirect() {
		return
	}

	// if user isn't admin, then logout user
	if nil == router.CurrentUser {

		// write flash message
		router.FlashWrite("NotPermit", "true")

		//router.Redirect("/login", 302)
		if router.Ctx.Input.IsAjax() {
			router.Ctx.WriteString("<script type=\"text/javascript\">window.location.href='/index'</script>")
		} else {
			router.Redirect("/login", 302)
		}
		return
	}

	router.Data["IsAdmin"] = true
	/*
		if app, ok := this.AppController.(ModelPreparer); ok {
			app.ModelPrepare()
			return
		}
	*/
}

// Get of BaseAdminRouter implemented global settings for all other routers.
func (router *BaseAdminRouter) Get() {
	router.TplName = "Get.tpl"
}

//Post of BaseAdminRouter implemented global settings for all other routers.
func (router *BaseAdminRouter) Post() {
	router.TplName = "Post.tpl"
}

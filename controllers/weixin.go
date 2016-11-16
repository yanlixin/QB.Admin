package controllers

import (
	"fmt"

	"QB.Admin/models"
	"QB.Admin/summaries"
	"github.com/astaxie/beego"
)

type WeixinController struct {
	beego.Controller
}

func (ctl *WeixinController) Prepare() {
	ctl.EnableXSRF = false
}
func (ctl *WeixinController) Index() {
	models.DefaultWeixin.HandleFunc(ctl.Ctx)
}
func (ctl *WeixinController) GetAllMessages() {
	ctl.Data["HasAdd"] = true
	ctl.TplName = "weixin/message/msglist.html"
}
func (ctl *WeixinController) EditMessage() {
	ctl.TplName = "weixin/message/msgedit.html"
}
func (ctl *WeixinController) SendMessage() {
	openids, err := models.DefaultWeixin.GetUserList()
	if nil != err {
		fmt.Printf("%+v", err)
	}
	fmt.Println(openids)
	result := summaries.JsonResultSuccess("0", nil)
	ctl.Data["json"] = result
	ctl.ServeJSON()
}

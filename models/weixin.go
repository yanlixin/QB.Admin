package models

import (
	"QB.Admin/utils/weixin"
	"github.com/astaxie/beego/context"
)

func init() {
	DefaultWeixin = newWeixin()
}

var DefaultWeixin *Weixin

type Weixin struct {
	TOKEN     string
	APPID     string
	APPSECRET string
}

func newWeixin() *Weixin {
	result := &Weixin{}
	result.TOKEN = "my-token"
	result.APPID = "wx3ecfceaf1a46783c"
	result.APPSECRET = "9aa84b30039ca7148399c33071017fe6"
	return result
}

func echo(w weixin.ResponseWriter, r *weixin.Request) {
	txt := r.Content          // 获取用户发送的消息
	w.ReplyText(txt)          // 回复一条文本消息
	w.PostText("Post:" + txt) // 发送一条文本消息
}

// 关注事件的处理函数
func subscribe(w weixin.ResponseWriter, r *weixin.Request) {
	w.ReplyText("欢迎关注") // 有新人关注，返回欢迎消息
}

func (wx *Weixin) HandleFunc(ctx *context.Context) {
	mux := weixin.New(wx.TOKEN, wx.APPID, wx.APPSECRET)
	// 注册文本消息的处理函数
	mux.HandleFunc(weixin.MsgTypeText, echo)
	// 注册关注事件的处理函数
	mux.HandleFunc(weixin.MsgTypeEventSubscribe, subscribe)

	mux.ServeHTTP(ctx.ResponseWriter, ctx.Request)
}

// 获取用户OpenID列表
func (wx *Weixin) GetUserList() ([]string, error) {
	mux := weixin.New(wx.TOKEN, wx.APPID, wx.APPSECRET)

	return mux.GetUserList()

}

// 获取服务器IP列表
func (wx *Weixin) GetIpList() ([]string, error) {
	mux := weixin.New(wx.TOKEN, wx.APPID, wx.APPSECRET)

	return mux.GetIpList() // mux.GetUserList()

}

func (wx *Weixin) PostText(touser string, text string) error {
	mux := weixin.New(wx.TOKEN, wx.APPID, wx.APPSECRET)

	return mux.PostText(touser, text) // mux.GetUserList()

}

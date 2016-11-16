package models

import (
	"testing"

	"QB.Admin/models"
)

func TestGetUserList(t *testing.T) {
	openids, err := models.DefaultWeixin.GetUserList()
	if nil != err {
		t.Error(err)
	}
	if len(openids) < 1 {
		t.Error("except more then zero actual 0")
	}
}
func TestGetIpList(t *testing.T) {
	openids, err := models.DefaultWeixin.GetIpList()
	if nil != err {
		t.Error(err)
	}
	if len(openids) < 1 {
		t.Error("except more then zero actual 0")
	}
}
func TestPostText(t *testing.T) {
	userOpenID := "osRttt3f2HVlNVOrLN4LuIFfiU74"
	msgText := "这是我的测试消息，不需要关注，如果疑问请与DavidYan来信"
	err := models.DefaultWeixin.PostText(userOpenID, msgText)
	if nil != err {
		t.Error(err)
	}

}

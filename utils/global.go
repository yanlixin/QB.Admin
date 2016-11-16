package utils

import (
	"encoding/base64"

	"github.com/astaxie/beego/logs"
	"golang.org/x/crypto/scrypt"
)

/*
const (
	AppDescription = ""
	AppKeywords    = ""
	AppName        = "DEMP.P2P"
	AppUrl         = ""
	AppLogo        = ""
	AvatarUlr      = ""
	IsProMode      = ""
)
*/
var (
	AppDescription string
	AppKeywords    string
	AppName        string
	AppVer         string
	AppURL         string
	AppLogo        string
	AvatarURL      string
	IsProMode      bool
	AppTitle       string
	ScryptKey      string
	CdnURL         string
	Logger         *logs.BeeLogger
)

const TestMessageBool = `Excepted %t actual %t`
const TestMessageLength = `Excepted %d actual %d`

func init() {
	AppTitle = "DEMP"
	CdnURL = "http://115.28.228.246/static/assets"
	ScryptKey = "&feixiang&"
	Logger = logs.NewLogger(10000)
	Logger.SetLogger("console", "")
	Logger.SetLogger("file", `{"filename":"log.txt"}`)
	Logger.EnableFuncCallDepth(true)
}
func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
func GenPwd(loginPwd string, salt string) []byte {
	if "" == salt {
		salt = ScryptKey
	}
	dk, _ := scrypt.Key([]byte(loginPwd), []byte(salt), 16384, 8, 1, 64)
	return dk
}
func GenPwdStr(loginPwd string) string {
	result := GenPwd(loginPwd, ScryptKey)
	return string(base64Encode([]byte(result)))

}
func CurrentUserId() int {
	return 1
}

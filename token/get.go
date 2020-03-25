package token

import (
	"fmt"
	"subscribe/tools"
)

var (
	appid    = "wxc48229a30edb7009"
	secret   = "57892e84d84750eaad2be4ea2d9f7a09"
	tokenApi = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"
)

// Info token信息
type Info struct {
	Token     string
	ExpiresAt int64
}

//GetAccessToken 获取access_token
func (t *Info) GetAccessToken() {
	// time.Now().Add(time.Second * 7200).Unix(),
	reqResult, err := tools.HTTPGet(tokenApi)
	if err != nil {
		panic(err)
	}
	fmt.Println(reqResult)
}

package auth

import (
	"encoding/json"
	"fmt"
	"psm/tools"
	"time"
)

var tokenAPI = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

// Token 获取的token
var Token TokenContainer

// TokenContainer token容器
type TokenContainer struct {
	Value     string `json:"access_token"`
	ExpiresAt int64  `json:"expires_in"`
}

// TokenParams 获取token需要的参数
type TokenParams struct {
	AppID, Secret string
}

//Get 获取access_token
func (p *TokenParams) Get() string {
	now := time.Now()
	if Token.ExpiresAt > now.Unix() { //token 没过期
		return Token.Value
	}
	var (
		tokenResult []byte
		err         error
	)
	if tokenResult, err = tools.NewRequest(
		fmt.Sprintf(tokenAPI, p.AppID, p.Secret), "").HTTPGet(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(tokenResult, &Token); err != nil {
		panic(err)
	}
	Token.ExpiresAt = now.Add(time.Second * 7200).Unix()
	return Token.Value
}

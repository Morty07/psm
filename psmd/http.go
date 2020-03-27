package psmd

import (
	"fmt"
	"net/http"
	"psm/auth"
	"psm/tools"

	"github.com/julienschmidt/httprouter"
)

var (
	appid        = "wxc48229a30edb7009"
	secret       = "57892e84d84750eaad2be4ea2d9f7a09"
	subscribeAPI = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=%s"
	data         = `{
		"touser": "opsM440GidAt7aZnoVczxZDNtxFo",
		"template_id": "6OU1apCDjY83m0CKXu0lG_nzAsF5JueWxVpTYP6vcFo",
		"data": {
			"name1": {
				"value": "滴滴滴"
			},
			"phone_number2": {
				"value": "18643492169"
			},
			"time3": {
				"value": "2020-03-26"
			} ,
			"thing4": {
				"value": "广州蛙鸣智能"
			}
		}
	}`
)

type HTTPServer struct {
	Router http.Handler
}

func (s *HTTPServer) SendSubscribeMsg(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	p := auth.TokenParams{appid, secret}
	if res, err := tools.NewRequest(fmt.Sprintf(subscribeAPI, p.Get()), data).HTTPPost(); err == nil {
		fmt.Println(string(res))
		fmt.Fprintf(w, "send success")
	} else {
		fmt.Fprintf(w, string(res))
	}
}

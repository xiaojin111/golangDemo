package main

import (
	"golangDemo/duanxin/models"
	"log"
)

var (
	AccountSid   = "8aaf07086f17620f016f64a98f8e309e"
	AppIDNew     = "8a216da86f696570016f88f5ec3f1086"
	AccountToken = "0b846098fe864018ba64ba04d301a21c"
	SignName     = "美年大健康"
)

func main() {
	NewSMSCode("15629092120", "643105", []string{"11点", "hello"}...)
}
func NewSMSCode(telenumber string, templateCode string, templateParam ...string) (err error) {
	cloopen := &models.Cloopen{
		AccountSid:   AccountSid,
		AppID:        AppIDNew,
		AccountToken: AccountToken,
	}
	log.Println(templateParam)
	req := &models.Request{
		Mobile:       telenumber,
		TemplateCode: templateCode,
		Datas:        templateParam,
	}
	var valid bool
	valid, err = cloopen.Send(req)
	if !valid {
		log.Println("短信错误", err.Error())
		return
	}
	return
}

package sms

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	config2 "orange_message_service/app/components/config"
)

func GetClient() *dysmsapi.Client {
	config := config2.GetConfig()
	REGION_ID := config.GetString("aliyunSms.regionId")
	ACCESS_KEY_ID := config.GetString("aliyunSms.keyId")
	ACCESS_KEY_SECRET := config.GetString("aliyunSms.secret")
	client, err := dysmsapi.NewClientWithAccessKey(REGION_ID, ACCESS_KEY_ID, ACCESS_KEY_SECRET)
	if err != nil {
		panic(err)
	}
	return client
}

func Send(phone string, TemplateCode string, TemplateParam map[string]interface{}) bool {
	templates, _ := json.Marshal(TemplateParam)

	client := GetClient()
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = TemplateParam["sign_name"].(string) //签名不能为空
	request.TemplateCode = TemplateCode
	request.TemplateParam = string(templates)

	//这里需要解析下 判断返回状态码
	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error(), response)
		return false
	}
	return true
}
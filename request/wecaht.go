package request

import (
	"WeChatPusher/config"
	"WeChatPusher/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAccessToken() (err error, AccessToken string) {
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + config.APP_ID + "&secret=" + config.APP_SECRET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, ""
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//fmt.Println("GET response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("GET response Body:", string(body))
	weResp := model.WeResp{}
	err2 := json.Unmarshal(body, &weResp)
	if err2 != nil {
		return err2, ""
	}
	return nil, weResp.AccessToken
}

func PushMessage(AccessToken string, data model.Data) error {
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + AccessToken

	message := model.Message{
		ToUser:     config.TO_USER,
		TemplateId: config.TEMPLATE_Id,
		Url:        data.FxLink.Value,
		Data:       data,
	}
	//json序列化
	jsonStr, err1 := json.Marshal(&message)
	if err1 != nil {
		fmt.Println("error:", err1)
		return err1
	}

	req, err2 := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err2 != nil {
		return err2
	}
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err3 := client.Do(req)
	if err3 != nil {
		return err3
	}
	defer resp.Body.Close()
	return nil
	//fmt.Println("POST response Status:", resp.Status)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("POST response Body:", string(body))
}

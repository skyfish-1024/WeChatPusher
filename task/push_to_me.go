package task

import (
	"WeChatPusher/config"
	"WeChatPusher/db"
	"WeChatPusher/request"
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Message struct {
	ToUser     string `json:"touser"`
	TemplateId string `json:"template_id"`
	Url        string `json:"url"`
	Data       Data   `json:"data"`
}

type Data struct {
	Time  Raw `json:"time,omitempty"`
	Day   Raw `json:"day,omitempty"`
	Count Raw `json:"count,omitempty"`
	State Raw `json:"state,omitempty"`
}
type Raw struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type WeResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func pushToMe(state string) error {
	count, err := db.GetCount()
	if err != nil {
		return err
	}
	data := Data{
		Time:  Raw{time.Now().Format("2006-01-02 15:04:05"), config.Pink},
		Day:   Raw{time.Now().Weekday().String(), config.Pink},
		Count: Raw{count, config.Pink},
		State: Raw{state, config.Pink},
	}
	err1, AccessToken := request.GetAccessToken()
	if err1 != nil {
		return err1
	}
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + AccessToken
	message := Message{
		ToUser:     config.TO_ME,
		TemplateId: config.TEMPLATE_Id_TO_ME,
		Url:        "www.zzyozzy.top",
		Data:       data,
	}
	//json序列化
	jsonStr, err2 := json.Marshal(&message)
	if err2 != nil {
		return err2
	}

	req, err3 := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err3 != nil {
		return err3
	}
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err4 := client.Do(req)
	if err4 != nil {
		return err4
	}
	defer resp.Body.Close()
	return nil
}

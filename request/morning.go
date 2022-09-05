package request

import (
	"WeChatPusher/config"
	"WeChatPusher/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GoodMorning() string {
	url := "http://api.tianapi.com/zaoan/index?key=" + config.MORNING_KEY
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err1 := client.Do(req)
	if err1 != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//fmt.Println("GET response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("GET response Body:", string(body))
	morningResp := model.MorningResp{}
	err2 := json.Unmarshal(body, &morningResp)
	if err2 != nil {
		fmt.Println(err2)
	}
	return morningResp.NewsList[0].Content
}

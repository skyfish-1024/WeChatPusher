package request

import (
	"WeChatPusher/config"
	"WeChatPusher/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetWeather() model.WeatherResp {
	url := "https://devapi.qweather.com/v7/weather/3d?location=" + config.LOCATION_ID + "&key=" + config.WEATHER_KEY
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
	weatherResp := model.WeatherResp{}
	err2 := json.Unmarshal(body, &weatherResp)
	if err2 != nil {
		fmt.Println(err2)
	}
	return weatherResp
}

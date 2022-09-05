package task

import (
	"WeChatPusher/config"
	"WeChatPusher/db"
	"WeChatPusher/model"
	"WeChatPusher/request"
	"fmt"
	"strconv"
	"time"
)

func pusher(day string) (*model.DbData, error) {
	weatherResp := request.GetWeather()
	MorningPhase := request.GoodMorning()
	data := model.Data{
		Lover:           model.Raw{Value: config.NAME, Color: config.Blue},
		Sender:          model.Raw{Value: config.SENDER, Color: config.Blue},
		FxLink:          model.Raw{weatherResp.FxLink, config.Purple},
		Time:            model.Raw{time.Now().Format("2006-01-02 15:04:05"), config.Purple},
		WeekDay:         model.Raw{time.Now().Weekday().String(), config.Purple},
		Location:        model.Raw{config.LOCATION, config.Yellow},
		Weather:         model.Raw{weatherResp.Daily[0].TextDay, config.Red},
		MaxTemperature:  model.Raw{weatherResp.Daily[0].TempMax, config.Red},
		MinTemperature:  model.Raw{weatherResp.Daily[0].TempMin, config.Red},
		SunRise:         model.Raw{weatherResp.Daily[0].Sunrise, config.Black},
		SunSet:          model.Raw{weatherResp.Daily[0].Sunset, config.Black},
		MoonPhase:       model.Raw{weatherResp.Daily[0].MoonPhase, config.Purple},
		UvIndex:         model.Raw{weatherResp.Daily[0].UvIndex, config.Purple},
		Humidity:        model.Raw{weatherResp.Daily[0].Humidity, config.Purple},
		AcquaintanceDay: model.Raw{day, config.Black},
		Birthday:        model.Raw{config.BIRTHDAY, config.BrightPink},
		MorningPhase:    model.Raw{MorningPhase, config.Pink},
	}
	//fmt.Println(data)
	err, AccessToken := request.GetAccessToken()
	if err != nil {
		return &model.DbData{}, err
	}
	err = request.PushMessage(AccessToken, data)
	if err != nil {
		return &model.DbData{}, err
	}
	dbData := model.DbData{
		Lover:           data.Lover.Value,
		Sender:          data.Sender.Value,
		SendTime:        data.Time.Value,
		WeekDay:         data.WeekDay.Value,
		Location:        data.Location.Value,
		Weather:         data.Weather.Value,
		MaxTemperature:  data.MaxTemperature.Value,
		MinTemperature:  data.MinTemperature.Value,
		SunRise:         data.SunRise.Value,
		SunSet:          data.SunSet.Value,
		MoonPhase:       data.MoonPhase.Value,
		UvIndex:         data.UvIndex.Value,
		Humidity:        data.Humidity.Value,
		AcquaintanceDay: data.AcquaintanceDay.Value,
		Birthday:        data.Birthday.Value,
		MorningPhase:    data.MorningPhase.Value,
	}
	return &dbData, nil
}

func Push() error {
	err1, day := db.GetDay()
	if err1 != nil {
		fmt.Println(err1)
		return err1
	}
	intDay, _ := strconv.Atoi(day)
	strDay := strconv.Itoa(intDay + 1)
	data, err2 := pusher(strDay)
	if err2 != nil {
		return err2
	}
	err3 := db.SaveData(data)
	if err3 != nil {
		return err3
	}
	return nil
}

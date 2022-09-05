package model

import "gorm.io/gorm"

type DbData struct {
	gorm.Model
	Lover           string `gorm:"lover"`            //亲爱的
	Sender          string `gorm:"sender"`           //发送人
	SendTime        string `gorm:"send_time"`        //发送当前时间
	WeekDay         string `gorm:"week_day"`         //星期几
	Location        string `gorm:"location"`         //地区
	Weather         string `gorm:"weather"`          //天气描述
	MaxTemperature  string `gorm:"max_temperature"`  //最高温
	MinTemperature  string `gorm:"min_temperature"`  //最低温
	SunRise         string `gorm:"sun_rise"`         //日出时间
	SunSet          string `gorm:"sun_set"`          //日落时间
	MoonPhase       string `gorm:"moon_phase"`       //月相
	UvIndex         string `gorm:"uv_index"`         //紫外线强度指数
	Humidity        string `gorm:"humidity"`         //相对湿度，百分比数值
	AcquaintanceDay string `gorm:"acquaintance_day"` //相识天数string
	Birthday        string `gorm:"birthday"`         //生日
	MorningPhase    string `gorm:"morning_phase"`    //早安心语
}

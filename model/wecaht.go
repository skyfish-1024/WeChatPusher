package model

type Message struct {
	ToUser     string `json:"touser"`
	TemplateId string `json:"template_id"`
	Url        string `json:"url"`
	Data       Data   `json:"data"`
}

type Data struct {
	Lover           Raw `json:"lover"`            //亲爱的
	Sender          Raw `json:"sender"`           //发送人
	FxLink          Raw `json:"fx_link"`          //天气详情链接
	Time            Raw `json:"time"`             //当前时间
	WeekDay         Raw `json:"week_day"`         //星期几
	Location        Raw `json:"location"`         //地区
	Weather         Raw `json:"weather"`          //天气描述
	MaxTemperature  Raw `json:"max_temperature"`  //最高温
	MinTemperature  Raw `json:"min_temperature"`  //最低温
	SunRise         Raw `json:"sun_rise"`         //日出时间
	SunSet          Raw `json:"sun_set"`          //日落时间
	MoonPhase       Raw `json:"moon_phase"`       //月相
	UvIndex         Raw `json:"uv_index"`         //紫外线强度指数
	Humidity        Raw `json:"humidity"`         //相对湿度，百分比数值
	AcquaintanceDay Raw `json:"acquaintance_day"` //相识天数
	Birthday        Raw `json:"birthday"`         //生日
	MorningPhase    Raw `json:"morning_phase"`    //早安心语
}

type DataToMe struct {
	Time  string
	Count string
	State string
}
type Raw struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type WeResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

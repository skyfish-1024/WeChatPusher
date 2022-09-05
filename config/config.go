package config

//你女朋友的信息
const (
	NAME     = "小呆瓜"
	BIRTHDAY = "农历七月十八日"
	SENDER   = "编程小余"
)

//mysql数据库配置
const (
	HOST     = "你的服务器地址" //主机地址
	DbName   = "你的数据库名称" //数据库名称
	UserName = "你的连接用户名" //MySQL连接用户名
	Password = "你的连接密码"  //连接密码
)

//wechat
const (
	APP_ID     = "你的微信app id"
	APP_SECRET = "你的app secret"

	TO_USER     = "你女朋友的扫描id"
	TEMPLATE_Id = "给你女朋友的模板id"

	TO_ME             = "你自己扫的id"
	TEMPLATE_Id_TO_ME = "推送结果给你自己的模板id"
)

//和风天气
const (
	LOCATION    = "渝北区"          //自行更改
	LOCATION_ID = "106.59,29.67" //自行更改
	WEATHER_KEY = "和风天气的接口key"
)

//早安心语
const (
	MORNING_KEY = "天行数据的用户key"
)

//颜色
const (
	White      = "#FFFFFF"
	Red        = "#FF0000"
	Green      = "#7FFF00"
	Blue       = "#0000FF"
	Yellow     = "#E47833"
	Black      = "#000000"
	Purple     = "#9F5F9F"
	Pink       = "#BC8F8F"
	BrightPink = "#FF1CAE"
)

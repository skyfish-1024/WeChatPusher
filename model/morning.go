package model

type MorningResp struct {
	Code     int        `json:"code"`
	Msg      string     `json:"msg"`
	NewsList []NewsList `json:"newslist"`
}
type NewsList struct {
	Content string `json:"content"`
}

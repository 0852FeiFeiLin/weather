package entity

type Result struct {
	//城市
	City string `json:"city"`
	//天气结果集详细信息
	Realtime Realtime	`json:"realtime"`
}

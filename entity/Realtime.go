package entity

//结构体里面的详细数据
type Realtime struct {
	Info        string  `json:"info"`
	Wid         string   `json:"wid"`
	Temperature string	`json:"temperature"`
	Humidity    string	`json:"humidity"`
	Direct      string	`json:"direct"`
	Power       string	`json:"power"`
	Api         string	`json:"api"`
	/*注意：一一对应api文档里面的字段*/
}

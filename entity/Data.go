package entity

type Data struct {
	//返回码
	Error_code int `json:"error_code"`
	//返回说明
	Reason string `json:"reason"`
	//返回结果集
	Result Result	`json:"result"`

	/*注意：结构体里面的字段要大写*/
}

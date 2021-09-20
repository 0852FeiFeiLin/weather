package active

import (
	"HttpDemoCode/entity"
	_ "HttpDemoCode/entity"
	"fmt"
)


//用户输入城市方法
func UserInput() (re entity.Result) {
	fmt.Println("hello world")
	fmt.Println("请输入你想查询天气的城市：")
	//var city string
	var result = entity.Result{}
	fmt.Scan(&result.City)
	return result  //返回一个结构体
}

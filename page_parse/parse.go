package page_parse

import (
	"HttpDemoCode/entity"
	"encoding/json"
	"fmt"
)
//解析数据
func ParseJson(weather []byte,err error){  //包名记得大写
	var data entity.Data
	jsonerr := json.Unmarshal(weather, &data) //参数1：解析的json数据 参数2：存放的解析后数据结构体，
	// json.Unmarshal该方法自动转换为对应类型
	if jsonerr != nil {
		fmt.Errorf("解码JSON异常:%v", err)
	} else { //解码成功
		/*
			注意：解码成功就代表我们拉取到了该系统的所有资源，
				通过key --> value 方式获取放进结构体里面了
				1、参考api文档的参数列表，然后在字段那边json一一对应
		*/
		//data.Error_code = data.["error_code"]

		fmt.Println(data.Error_code) //返回码(该api返回码为0，代表成功)
		fmt.Println(data.Reason)     //返回说明
		fmt.Println(data.Result)     //结果集，结果集里忙有realtime

		if data.Error_code == 0 { //先判断返回码，为0成功，非0失败
			//请求成功
			fmt.Println("请求成功")
		} else {
			fmt.Printf("请求失败：%v_%v",data.Error_code,data.Reason)//返回响应码和返回说明
		}
}
}

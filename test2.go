package main

import (
	"HttpDemoCode/active"
	"HttpDemoCode/page_parse"
	"HttpDemoCode/request"
	"fmt"
)
//获取天气预报方法3
/*
	1、读取用户输入的city
	2、请求网页为聚合数据
	3、设置url，模拟客户端发起请求
	4、获得响应数据，解析数据
url格式为："http://apis.juhe.cn/simpleWeather/query?city=苏州&key=d7ab14f872f0f7ba59874b3d86b99b21"
*/
func main() {
	//准备工作
	apiurl := "http://apis.juhe.cn/simpleWeather/query?"
	//调用用户输入方法获得需求city
	result1 := active.UserInput()
	//city转换为字符串
	result2 := string(result1.City)
	//key
	key1 := "26c586345e411e596178dd3121844420"
	//拼接url
	url1 := apiurl + "city=" + result2 + "&key=" + key1
	fmt.Println("拼接后的url：", url1)

	/*另一种方式：
	param := url.Values{}   //存放参数，这是一个key - value 的键值对映射map
	param.Set("city",result1.City)
	param.Set("key",key1)  //设置map的值，后面传给get方法
	http.Get(apiurl,param)
	*/

	//发送请求 (请求方法，url，请求体)
	weather, err := request.Request("GET", url1, nil) //返回响应
	if err != nil {
		//创建一个error对象
		fmt.Errorf("请求异常:\r\n%v", err)
	} else {
		page_parse.ParseJson(weather,err)
		}
	}


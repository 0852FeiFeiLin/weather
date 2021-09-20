package main

import (
	"HttpDemoCode/active"
	"HttpDemoCode/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)
/*
	1、读取用户输入的city
	2、请求网页为聚合数据
	3、*/
func main() {
	apiurl := "http://apis.juhe.cn/simpleWeather/query?"

	result1 := active.UserInput()
	key1 := "26c586345e411e596178dd3121844420"
	 param := url.Values{}   //存放参数，这是一个key - value 的键值对映射map

	param.Set("city",result1.City)
	param.Set("key",key1)  //设置map的值，后面传给get方法

	//发送请求 (url拼接)
	weather, err := GetWeather(apiurl, &param)//返回响应
	//http.Get(apiurl,param)
	if err != nil{
		//创建一个error对象
		fmt.Errorf("请求异常:\r\n%v",err)
	}else {
		/*//创建一个map，用于存放解析后的数据   不，我不用这个，用结构体存
		var netreturn map[string]interface{}

			//realTime := data.(map[string]interface{})["realtime"]
			/*解释：这个语法是，data结果集数据下面的realtime字段，通过键值对的方式
					realtime是key -->映射value到realtime里，返回的value是interface类型的*/
		//进行解析 byte ---> string?
		var data  entity.Data
		jsonerr := json.Unmarshal(weather, &data) //参数1：解析的json数据 参数2：存放的解析后数据结构体，
		// json.Unmarshal该方法自动转换为对应类型
		if jsonerr != nil{
			fmt.Errorf("解码JSON异常:%v",err)
		}else{  //解码成功转为unicode
			/*
			注意：解码成功就代表我们拉取到了该系统的所有资源，
				通过key --> value 方式获取放进结构体里面了
				1、参考api文档的参数列表，然后在字段那边json一一对应
			*/
			//data.Error_code = data.["error_code"]

			fmt.Println(data.Error_code) //返回码(该api返回码为0，代表成功)
			fmt.Println(data.Reason)   //返回说明
			fmt.Println(data.Result)   //结果集，结果集里忙有realtime

			if data.Error_code == 0{   //先判断返回码，为0成功，非0失败
				//请求成功
				fmt.Println("请求成功")


			}else{
				//fmt.Printf("请求失败：%v_%v",errorCode,reason)//返回响应码和返回说明
				fmt.Println("请求失败")
			}

		}

	}
}
func GetWeather(apiurl string ,param *url.Values) (rs []byte ,err error){
	var Url *url.URL   //声明一个url包下面的Url结构体
	Url, err = url.Parse(apiurl) //将url进行自动解析并返回一个url
	if err != nil{
		fmt.Println(err.Error())
		fmt.Println("解析apiurl失败")
		return
	}
	//现在就是把所有的URl都放进了Url结构体里面
	Url.RawQuery = param.Encode() //Encode方法将(我们的url)值编码为url格式   为啥要用这个字段？？？？
	//发起请求
	resp, err := http.Get(Url.String())//返回响应数据了
	if err != nil {
		fmt.Printf("请求失败，失败状态码为：%d\n",resp.StatusCode)
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("请求成功,状态码%d\n",resp.StatusCode)
	//延迟关闭请求响应数据流，避免资源长占用cpu
	defer resp.Body.Close()

	//返回读取的响应数据byte类型
	return ioutil.ReadAll(resp.Body)

}
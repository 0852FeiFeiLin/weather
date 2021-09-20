package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main1() {
	// 接口请求URL
	apiUrl := "http://apis.juhe.cn/simpleWeather/query"

	// 初始化参数
	param := url.Values{}

	// 接口请求参数
	param.Set("city", "北京")                              // 要查询的城市名称/id，城市名称如：温州、上海、北京
	param.Set("key", "26c586345e411e596178dd3121844420") // 接口请求Key

	// 发送请求
	data, err := Get(apiUrl, param)
	if err != nil {
		// 请求异常，根据自身业务逻辑进行调整修改
		fmt.Errorf("请求异常:\r\n%v", err)
	} else {
		var netReturn map[string]interface{}
		jsonerr := json.Unmarshal(data, &netReturn) //解析的josn数据，接收解析的数据结构
		if jsonerr != nil {
			// 解析JSON异常，根据自身业务逻辑进行调整修改
			fmt.Errorf("请求异常:%v", jsonerr)
		} else {
			errorCode := netReturn["error_code"]
			reason := netReturn["reason"]
			data := netReturn["result"]
			// 当前天气信息
			realtime := data.(map[string]interface{})["realtime"]
			/*解释：这个语法是，data结果集数据下面的realtime字段，通过键值对的方式
			realtime是key -->映射value到realtime里，类型是一个interface的 */
			if errorCode.(float64) == 0 {
				// 请求成功，根据自身业务逻辑进行调整修改
				fmt.Printf("温度：%v\n湿度：%v\n天气：%v\n风向：%v\n风力：%v\n空气质量：%v",
					realtime.(map[string]interface{})["temperature"],
					realtime.(map[string]interface{})["humidity"],
					realtime.(map[string]interface{})["info"],
					realtime.(map[string]interface{})["direct"],
					realtime.(map[string]interface{})["power"],
					realtime.(map[string]interface{})["aqi"],
				)
			} else {
				// 查询失败，根据自身业务逻辑进行调整修改
				fmt.Printf("请求失败:%v_%v", errorCode.(float64), reason)
			}
		}
	}
}

// get 方式发起网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String()) //发送请求返回响应
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

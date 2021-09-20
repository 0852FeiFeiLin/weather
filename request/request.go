package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)
//另一种发起请求的方法

func Request(method ,url string,body io.Reader)(rs []byte,err error){
	//创建客户端对象
	cliend := http.Client{
		Timeout: 30 * time.Second,
	}
	//创建一个请求对象
	request, err := http.NewRequest(method, url, body)
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	//发送请求
	response, err := cliend.Do(request)
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	//读取响应数据
	bytes, err := ioutil.ReadAll(response.Body)

	return bytes,nil
}

func Request1(apiurl string ,param *url.Values)(rs []byte ,err error){
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

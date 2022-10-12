package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
//进行网络请求
func response(url string) {
	fmt.Println("Step1:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(response)
	fmt.Println("Step2:", url)
	//延迟调用  一般用于资源回收
	defer response.Body.Close()

	fmt.Println("Step3:", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//输出页面主体部分
	// fmt.Println(body)
	ioutil.WriteFile("./f33.jpg", body, 0666)
	fmt.Println("Step4:", len(body))
}

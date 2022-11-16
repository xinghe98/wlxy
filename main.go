package main

import (
	"fmt"
	"github.com/xinghe98/wlxy/login"
	"io/ioutil"
	"os"
)

func main() {
	err := os.Setenv("HTTP_PROXY", "http://127.0.0.1:8888")
	if err != nil {
		return
	}
	cap := login.GetCaphta("http://wlxy.jxnxs.com/app/captcha/captcha")
	client, _ := login.GetCookie("http://wlxy.jxnxs.com/app/user/single/userlogin/login", "110379", "/DFGws7yGmJIUmbuYMU+Mg==", cap)
	fmt.Println(client)
	res, err := client.Get("http://wlxy.jxnxs.com/app/home")
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(body)
	fmt.Printf("%S", bodyStr)

}

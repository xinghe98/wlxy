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
	caphta := login.GetCaphta("http://wlxy.jxnxs.com/app/captcha/captcha")
	client, _, err := login.GetCookie("http://wlxy.jxnxs.com/app/user/single/userlogin/login", login.Login{
		Username: "110379", Password: "/DFGws7yGmJIUmbuYMU+Mg==", Caphta: caphta,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Get("http://wlxy.jxnxs.com/app/home")
	if err != nil {
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(body)
	fmt.Printf("%S", bodyStr)

}

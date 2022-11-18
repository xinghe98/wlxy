package main

import (
	"fmt"
	"github.com/xinghe98/wlxy/course"
	"github.com/xinghe98/wlxy/login"
	"os"
)

func main() {
	err := os.Setenv("HTTP_PROXY", "http://127.0.0.1:8888")
	if err != nil {
		return
	}
	caphta := login.GetCaphta("http://wlxy.jxnxs.com/app/captcha/captcha")
	client, _, err := login.GetCookie("http://wlxy.jxnxs.com/app/user/single/userlogin/login",
		"110371", "/DFGws7yGmJIUmbuYMU+Mg==", caphta)
	if err != nil {
		fmt.Println(err)
		return
	}
	course.GetMyCourse(client)
}

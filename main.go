package main

import (
	"fmt"
	"github.com/xinghe98/wlxy/login"
)

func main() {

	cap := login.GetCaphta("http://wlxy.jxnxs.com/app/captcha/captcha")
	cookie := login.GetCookie("http://wlxy.jxnxs.com/app/user/single/userlogin/login", "110371", "/DFGws7yGmJIUmbuYMU+Mg==", cap)
	fmt.Println(cookie)

}

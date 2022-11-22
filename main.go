//go:generate goversioninfo
package main

import (
	"fmt"
	"github.com/xinghe98/wlxy/course"
	"github.com/xinghe98/wlxy/login"
	"github.com/xinghe98/wlxy/util"
	"os"
)

func main() {
	//err := os.Setenv("HTTP_PROXY", "http://127.0.0.1:8888")
	//if err != nil {
	//	return
	//}

	var username, basePassword string
	fmt.Print("请输入柜员号：")
	//goland:noinspection ALL
	fmt.Scanln(&username)
	fmt.Print("请输入密码：")
	//goland:noinspection ALL
	fmt.Scanln(&basePassword)
	password := util.EncodePassword(basePassword)
	sigin := login.Login{Username: username, Password: password}
	caphta := sigin.GetCaphta("http://wlxy.jxnxs.com/app/captcha/captcha")
	client, _, err := sigin.GetCookie("http://wlxy.jxnxs.com/app/user/single/userlogin/login", caphta)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%d.我要学习\n%d.退出\n", 1, 0)
	var choice int
	//goland:noinspection ALL
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		getCourseInfo := course.GetCourseInfo{Session: client}

		for {
			itmId := getCourseInfo.GetMyCourse()
			if itmId == 0 {
				os.Exit(0)
			}
			getCourseInfo.GetCourseDetail(itmId)
		}
	case 0:
		return
	}

}

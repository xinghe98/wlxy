package login

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	"github.com/xinghe98/wlxy/util"
	"golang.org/x/net/publicsuffix"
)

var client *http.Client

func GetCaphta(url string) string {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client = &http.Client{
		Jar: cookieJar,
	}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile("caphta.png", content, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Println("请输入验证码：")
	var caphta string
	fmt.Scanln(&caphta)
	return caphta
}

func GetCookie(uri string, username string, password string, caphta string) []*http.Cookie {
	data := make(map[string]string)
	data["usrSteUsrId"] = username
	data["userPassword"] = password
	data["usrCode"] = caphta
	b := util.MapToJson(data)
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(b))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	if err != nil {
		panic(err)
	}
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("登录成功")
		cookie := resp.Cookies()
		return cookie
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		data = util.JsonToMap[string](string(data))
		fmt.Println("登录失败")
		return nil
	}

}

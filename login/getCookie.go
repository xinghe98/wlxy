package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xinghe98/wlxy/util"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

type Login struct {
	Username string `json:"usrSteUsrId"`
	Password string `json:"userPassword"`
	Caphta   string `json:"usrCode"`
}

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
	fmt.Print("请输入验证码：")
	var caphta string
	fmt.Scanln(&caphta)
	return caphta
}

func GetCookie(uri string, info Login) (*http.Client, []*http.Cookie, error) {
	data := make(map[string]string)
	data["usrSteUsrId"] = info.Username
	data["userPassword"] = info.Password
	data["usrCode"] = info.Caphta
	b, _ := json.Marshal(data)
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(b))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	if err != nil {
		panic(err)
	}
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	response := util.JsonToMap(resp.Body)
	if response["code"] == float64(200) {
		fmt.Println("登录成功")
		cookie := resp.Cookies()
		return client, cookie, nil
	} else {
		return nil, nil, fmt.Errorf("登录失败,错误信息：%s", response["msg"])
	}

}

package login

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/xinghe98/wlxy/util"
	"golang.org/x/net/publicsuffix"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
)

type Login struct {
	Username string
	Password string
}

var client *http.Client

func (l *Login) GetCaphta(url string) string {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client = &http.Client{
		Jar: cookieJar,
	}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	path, err := os.Create("caphta.png")
	if err != nil {
		return ""
	}
	_, err = io.Copy(path, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print("请输入验证码：")
	var caphta string
	fmt.Scanln(&caphta)
	return caphta
}

func (l *Login) GetCookie(uri string, caphta string) (*http.Client, []*http.Cookie, error) {
	data := make(map[string]string)
	data["usrSteUsrId"] = l.Username
	data["userPassword"] = l.Password
	data["usrCode"] = caphta
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, _ := json.Marshal(&data)
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(b))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	if err != nil {
		panic(err)
	}
	resp, _ := client.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	response := util.JsonToMap(resp.Body)
	if response["code"] == float64(200) {
		fmt.Println("登录成功")
		cookie := resp.Cookies()
		return client, cookie, nil
	} else {
		return nil, nil, fmt.Errorf("登录失败,错误信息：%s", response["msg"])
	}

}

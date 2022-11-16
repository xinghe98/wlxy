package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

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
	fmt.Print("请输入验证码：")
	var caphta string
	fmt.Scanln(&caphta)
	return caphta
}

func GetCookie(uri string, username string, password string, caphta string) (*http.Client, []*http.Cookie) {
	data := make(map[string]string)
	data["usrSteUsrId"] = username
	data["userPassword"] = password
	data["usrCode"] = caphta
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
	var response map[string]interface{}
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &response)
	if response["code"] == float64(200) {
		fmt.Println("登录成功")
		cookie := resp.Cookies()
		return client, cookie
	} else {
		fmt.Println("登录失败")
		//fmt.Printf("%T\n", response["code"])
		fmt.Println(response["msg"])
		fmt.Println(response["code"])
		//fmt.Println(len(response["code"]))
		return nil, nil
	}

}

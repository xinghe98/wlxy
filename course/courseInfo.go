package course

import (
	"fmt"
	"github.com/xinghe98/wlxy/util"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//course_id:resId (大节课程id)
//student_id:student_id (学员id)
//lesson_id : mod_id (小节课程id)
//lesson_location : 'default', (小节课程位置)
//time : duration,
//start_time : start_time,
//tkh_id : tkh_id|app_tkh_id
//cmt_lrn_pass_ind 判断课程是否完成

type GetCourseInfo struct {
	Session *http.Client
}

// 专门用来post接口获取课程信息的接口
func (g GetCourseInfo) requests(method string, uri string, b string) *http.Request {
	request, err := http.NewRequest(method, uri, strings.NewReader(b))
	request.Header.Set("Accept", "application/json,text/javascript, */*; q=0.01")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	if err != nil {
		panic(err)
	}
	return request
}

// GetMyCourse 访问 http://wlxy.jxnxs.com/app/course/getMyCourse 获取未学课程列表
func (g GetCourseInfo) GetMyCourse() int {
	b := "pageNo=1&pageSize=10&appStatus=I&pdate=" + strconv.FormatInt(time.Now().Unix(), 10)
	request := g.requests("POST", "http://wlxy.jxnxs.com/app/course/getMyCourse", b)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, _ := g.Session.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	response := util.JsonToMap(resp.Body)
	//fmt.Println(response["rows"])
	courseRows := response["rows"].([]interface{})
	fmt.Println("未看课程列表：")
	for i := 0; i < len(courseRows); i++ {
		title := courseRows[i].(map[string]interface{})["item"].(map[string]interface{})["itm_title"]
		id := courseRows[i].(map[string]interface{})["item"].(map[string]interface{})["itm_id"]
		fmt.Printf("%d、%s\n", int(id.(float64)), title)
	}
	fmt.Printf("请输入需要看的课程编号（课程名前面的数字）:")
	var itemId int
	input, err := fmt.Scanln(&itemId)
	fmt.Println(input)
	if err != nil {
		panic(err)
	}
	return itemId
}

// GetCourseDetail [rows][?][item][itm_id] 获取课程itm_id
// 访问 http://wlxy.jxnxs.com/app/course/detailJson/+itm_id 获取课程详情
func (g GetCourseInfo) GetCourseDetail(itmId int) {
	fmt.Printf("正在获取课程信息...%s\n", strconv.Itoa(itmId))
	urlStr := "http://wlxy.jxnxs.com/app/course/detailJson/" + strconv.Itoa(itmId) + "?pdate=" + strconv.FormatInt(time.Now().Unix(), 10)
	request := g.requests("GET", urlStr, "")
	fmt.Println(request.URL)
	resp, _ := g.Session.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	response := util.JsonToMap(resp.Body)
	fmt.Println(response["coscontent"])
	// for循环内一次性调用直接快进视频的接口，将该课程内的所有视频都看完

}

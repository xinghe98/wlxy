package course

import (
	"fmt"
	"github.com/xinghe98/wlxy/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//course_id:resId (小节课程id)
//student_id:student_id (学员id)
//lesson_id : mod_id (小节课程id)
//lesson_location : 'default', (小节课程位置)
//time : duration,
//start_time : start_time,
//tkh_id : tkh_id|app_tkh_id
//cmt_lrn_pass_ind 判断课程是否完成

type CourseInfo struct {
	itmId    string //课程id
	itmTitle string //课程名称
}

// GetMyCourse 访问 http://wlxy.jxnxs.com/app/course/getMyCourse 获取课程列表
func GetMyCourse(session *http.Client) {
	b := "pageNo=1&pageSize=10&appStatus=I&pdate=" + strconv.FormatInt(time.Now().Unix(), 10)
	request, err := http.NewRequest("POST", "http://wlxy.jxnxs.com/app/course/getMyCourse", strings.NewReader(b))
	request.Header.Set("Referer", "http://wlxy.jxnxs.com/app/course/signup")
	request.Header.Set("Accept", "application/json,text/javascript, */*; q=0.01")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")

	if err != nil {
		panic(err)
	}
	resp, _ := session.Do(request)
	defer resp.Body.Close()
	response := util.JsonToMap(resp.Body)
	//fmt.Println(response["rows"])
	courserows := response["rows"].([]interface{})
	fmt.Println("未看课程列表：")
	for i := 0; i < len(courserows); i++ {
		title := courserows[i].(map[string]interface{})["item"].(map[string]interface{})["itm_title"]
		id := courserows[i].(map[string]interface{})["item"].(map[string]interface{})["itm_id"]
		fmt.Printf("%d、%s\n", int(id.(float64)), title)
	}
	fmt.Printf("请输入需要看的课程编号（课程名前面的数字）:")
	var input int
	fmt.Scanln(&input)
	fmt.Printf("正在获取课程信息...%d", input)
}

// [rows][?][item][itm_id] 获取课程itm_id
// 访问 http://wlxy.jxnxs.com/app/course/detailJson/+itm_id 获取课程详情

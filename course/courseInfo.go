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
//student_id:userEntId (学员id)
//lesson_id : mod_id (小节课程id)
//lesson_location : 'default', (小节课程位置)
//time : mod_required_time,
//start_time : start_time,
//tkh_id : tkh_id|app_tkh_id
//cmt_lrn_pass_ind 判断课程是否完成

type GetCourseInfo struct {
	Session *http.Client
}

// 专门用来post接口获取课程信息的接口
func (g *GetCourseInfo) requests(method string, uri string, b string) *http.Request {
	request, err := http.NewRequest(method, uri, strings.NewReader(b))
	request.Header.Set("Accept", "application/json,text/javascript, */*; q=0.01")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	if err != nil {
		panic(err)
	}
	return request
}

// GetMyCourse 访问 http://wlxy.jxnxs.com/app/course/getMyCourse 获取未学课程列表
func (g *GetCourseInfo) GetMyCourse() int {
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
func (g *GetCourseInfo) GetCourseDetail(itmId int) {
	fmt.Printf("正在获取课程信息...%s\n", strconv.Itoa(itmId))
	urlStr := "http://wlxy.jxnxs.com/app/course/detailJson/" + strconv.Itoa(itmId) + "?pdate=" + strconv.FormatInt(time.Now().Unix(), 10)
	request := g.requests("GET", urlStr, "")
	//fmt.Println(request.URL)
	resp, _ := g.Session.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	response := util.JsonToMap(resp.Body)
	//fmt.Println(response["coscontent"])
	resId := int(response["resId"].(float64))                                      //cos_id
	tkhId := int(response["app"].(map[string]interface{})["app_tkh_id"].(float64)) //tkh_id
	studentId := int(response["userEntId"].(float64))                              //student_id
	cmtList := response["ccr"].(map[string]interface{})["cmt_lst"].([]interface{}) //cmt_list
	//fmt.Printf("resId:%d,tkhId:%d,studentId:%d", resId, tkhId, studentId)
	for i := 0; i < len(cmtList); i++ {
		fmt.Println(i)
		content := cmtList[i].(map[string]interface{})["res"].(map[string]interface{})
		// 使用课程文件类型以及课程状态判断是否需要学习
		//这里含有状态的列表顺序不是下面含有课程时长的顺序
		resType := content["res_type"]
		statusPassed := cmtList[i].(map[string]interface{})["cmt_lrn_pass_ind"] // 课程状态
		cmtTitle := cmtList[i].(map[string]interface{})["cmt_title"]            // 课程标题
		//fmt.Println(resType, statusPassed)
		if resType == "DXT" && statusPassed == false {
			fmt.Println("这是个考试，暂时不支持")
		}
		if resType == "VOD" && statusPassed == false {
			fmt.Printf("正在加速观看..........%s\n", cmtTitle)
			modId := int(content["res_id"].(float64))                                                                                                                                                   //lesson_id
			requireTime := int(response["coscontent"].([]interface{})[i].(map[string]interface{})["resources"].(map[string]interface{})["mod"].(map[string]interface{})["mod_required_time"].(float64)) //time
			hour, minute, second := util.ResolveTime(requireTime + 120)
			fmt.Printf("本课程需要观看时间：%d小时%d分钟%d秒\n", hour, minute, second)
			timeStr := fmt.Sprintf("%s:%s:%s", strconv.Itoa(hour), strconv.Itoa(minute), strconv.Itoa(second))
			starttime := util.GenerateTime(requireTime + 360)
			fmt.Printf("resId:%d,tkhId:%d,modId:%d,studentId:%d,requiretime:%d\n", resId, tkhId, modId, studentId, requireTime)
			fmt.Println(response["coscontent"].([]interface{})[i].(map[string]interface{})["resources"].(map[string]interface{}))
			learnCourse := LearnCourse{
				CosId:       strconv.Itoa(resId),
				StudentId:   strconv.Itoa(studentId),
				LessonId:    strconv.Itoa(modId),
				RequireTime: timeStr,
				ModType:     resType.(string),
				TkhId:       strconv.Itoa(tkhId),
				StartTime:   starttime,
				GetCourseInfo: &GetCourseInfo{
					Session: g.Session,
				},
			}
			learnCourse.Learn()
		}
	}
	// for循环内一次性调用直接快进视频的接口，将该课程内的所有视频都看完
}

package course

import (
	"fmt"
	"net/http"
)

type LearnCourse struct {
	CosId         string
	StudentId     string
	LessonId      string
	RequireTime   string
	TkhId         string
	ModType       string
	StartTime     string
	GetCourseInfo *GetCourseInfo
}

func (l *LearnCourse) Learn() {
	req, err := http.NewRequest("POST", "http://wlxy.jxnxs.com/app/course/saveVod", nil)
	if err != nil {
		return
	}
	refer := fmt.Sprintf("http://wlxy.jxnxs.com/servlet/qdbAction?env=wizb&cmd=get_mod&mod_id=%s&mod_type=%s&tpl_use=lrn_vod.xsl&cos_id=%s&tkh_id=%s&test_style=&stylesheet=lrn_vod.xsl&url_failure=/htm/close_window.htm&page=0", l.LessonId, l.ModType, l.CosId, l.TkhId)
	//fmt.Println(refer)
	req.Header.Set("Referer", refer)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	q := req.URL.Query()
	q.Add("course_id", l.CosId)
	q.Add("student_id", l.StudentId)
	q.Add("lesson_id", l.LessonId)
	q.Add("lesson_location", "default")
	q.Add("time", l.RequireTime)
	q.Add("start_time", l.StartTime)
	q.Add("tkh_id", l.TkhId)
	q.Add("lesson_status", "I")
	req.URL.RawQuery = q.Encode()
	//fmt.Println(req.URL.String())
	resp, err := l.GetCourseInfo.Session.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("学习成功")
	} else {
		fmt.Println("学习失败,请联系开发者")
	}
}

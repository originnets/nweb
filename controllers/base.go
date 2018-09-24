package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Read(resp map[string]interface{} ){
	c.Data["json"] = resp
	c.ServeJSON()
}


func Md5(str string) (MD5PAW string) {
	DATA := []byte(str)
	MD5PAW = fmt.Sprintf("%x",md5.Sum(DATA))
	return
}

//func Mail()(){
//	auth := smtp.PlainAuth("", "953637695@qq.com", "password", "smtp.qq.com")
//	to := []string{"874560965@qq.com"}
//	nickname := "test"
//	user := "953637695@qq.com"
//	subject := "test mail"
//	content_type := "Content-Type: text/plain; charset=UTF-8"
//	body := "This is the email body."
//	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
//		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
//	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
//	if err != nil {
//		fmt.Printf("send mail error: %v", err)
//	}
//}
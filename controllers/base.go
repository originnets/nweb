package controllers

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"os"
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

//实现配置文件写入指定参数并写到指定文件夹
func GenConfFile(port int64, sname, root, logname string) (err error) {
	t, err1:= template.ParseFiles("template/vhost.tpl")
	if err1 != nil {
		err = err1
		return
	}

	actor := make(map[string]interface{})
	actor["Port"] = port
	actor["Server_name"] = sname
	actor["Root"] = root
	actor["Logname"] = logname

	filename := beego.AppConfig.String("path") + "/" + sname + ".conf"
	//os.O_WRONLY 只写  os.O_CREATE 如果指定文件不存在，就创建该文件 os.O_TRUNC 如果指定文件已存在，就将该文件的长度截为0,即清空文件
	f, err2 := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		err = err2
		return
	}
	defer f.Close()
	//写入文件中
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	if err3 := t.Execute(writer, actor); err3 != nil {
		err = err3
		return
	}
	err = nil
	return
}

//判断文件是否存在
func FileExistence(f string)(bool, error){
	_, err := os.Stat(f)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

//实现删除配置文件
func DelGenConfFile(sname string)(err error) {
	filename := beego.AppConfig.String("path") + "/" + sname + ".conf"
	status, err1 := FileExistence(filename)
	if status == true {
		if err1 := os.Remove(filename); err1 != nil {
			err = err1
			return
		}
		err = nil
		return
	}
	err = err1
	return
}

//实现移动配置文件
func MvGenConfFile(sname string)(err error) {
	filename := beego.AppConfig.String("path") + "/" + sname + ".conf"
	status, err2 := FileExistence(filename)
	if status == true {
		b_file := beego.AppConfig.String("bakpath") + "/" + sname + ".conf"
		beego.Info(b_file)
		if err1 := os.Rename(filename, b_file); err1 != nil {
			err = err1
			return
		}
	}
	err = err2
	return
}


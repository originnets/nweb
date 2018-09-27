package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type User struct {
	Id			int64
	Name		string	`orm:"size(30);unique"`
	Password 	string	`orm:"size(60)"`
	Ctime		time.Time `orm:"auto_now_add;type(datetime)"`
	Utime		time.Time `orm:"auto_now;type(datetime)"`
	Domain		[]*Domain `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Domain struct {
	Id			int64
	Sname		string `orm:"column(server_name);unique"`
	Port		int64  `orm:"column(port)"`
	Root		string `orm:"column(root)"`
	Logname		string `orm:"column(logname)"`
	Status		int64  `orm:"column(status);default(1)"`	//1.表示启用, 2.表示未启用
	User		*User  `orm:"rel(fk)"`    //设置一对多关系

}

func init() {

	//读取app.conf文件中的配置

	username := beego.AppConfig.String("username")
	hostname := beego.AppConfig.String("hostname")
	password := beego.AppConfig.String("password")
	dbname := beego.AppConfig.String("dbname")
	port := beego.AppConfig.String("port")

	//当port 为空指定默认端口3306
	if port == "" {
		port = "3306"
	}

	//设置连接串
	conn := username + ":" + password + "@tcp(" + hostname + ":" + port + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn, 30) //连接数据库

	orm.RegisterModel(new(User), new(Domain)) //注册表

	orm.RunSyncdb("default", false, false) //force 是否同步表结构 verbone是否创建表
}

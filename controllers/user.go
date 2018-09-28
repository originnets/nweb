package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"nweb/models"
)

type UserController struct {
	BaseController
}

//主页
func (c * UserController)GetIndex() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//注册
func (c *UserController) PostReg() {
	resp := make(map[string]interface{})
	defer c.Read(resp)

	//获取post数据
	regdata := make(map[string]string)
	json.Unmarshal(c.Ctx.Input.RequestBody, &regdata)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &regdata); err != nil {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}

	//验证数据
	if regdata["username"] == "" || regdata["password"] == "" {
		resp["code"] = models.RECODE_NODATA
		resp["meg"] = models.ReCodeText(models.RECODE_NODATA)
		return
	}

	//插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = regdata["username"]
	user.Password = Md5(regdata["password"])	//密码MD5加密处理
	if _,err := o.Insert(&user); err != nil {
		resp["code"] = models.RECODE_DBERR
		resp["meg"] = models.ReCodeText(models.RECODE_DBERR)
		return
	}

	resp["code"] = models.RECODE_OK
	resp["meg"] = models.ReCodeText(models.RECODE_OK)

}

//登陆
func (c *UserController) PostLogin() {
	resp := make(map[string]interface{})
	defer c.Read(resp)

	//从缓存中拿数据
	username := c.GetSession("username")
	if username != nil {
		resp["code"] = models.RECODE_OK
		resp["meg"] = models.ReCodeText(models.RECODE_OK)
		resp["cache"] = "缓存"
		return
	}


	if err3 := GenConfFile(8080,"test.com","/home/www","test"); err3 != nil {
		beego.Info("写入错误")
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}
	//if err := DelGenConfFile("test2.com"); err != nil {
	//	beego.Info("写入错误")
	//	resp["code"] = models.RECODE_DATAERR
	//	resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
	//	return
	//}
	MvGenConfFile("test1.com")

	//获取post数据
	logindata := make(map[string]string)
	json.Unmarshal(c.Ctx.Input.RequestBody, &logindata)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &logindata); err != nil {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}

	//验证数据
	if logindata["username"] == "" || logindata["password"] == "" {
		resp["code"] = models.RECODE_NODATA
		resp["meg"] = models.ReCodeText(models.RECODE_NODATA)
		return
	}

	o := orm.NewOrm()
	user := models.User{Name: logindata["username"] , Password:Md5(logindata["password"])}
	if err := o.Read(&user,"Name", "Password"); err != nil {
		resp["code"] = models.RECODE_DBERR
		resp["meg"] = models.ReCodeText(models.RECODE_DBERR)
		return
	}

	c.SetSession("username", user.Name)

	resp["code"] = models.RECODE_OK
	resp["meg"] = models.ReCodeText(models.RECODE_OK)

}

//退出登入
func (c *UserController) GetLogout() {
	resp := make(map[string]interface{})
	defer c.Read(resp)

	//删除
	c.DelSession("user_id")
	resp["code"] = models.RECODE_OK
	resp["meg"] = models.ReCodeText(models.RECODE_OK)
}

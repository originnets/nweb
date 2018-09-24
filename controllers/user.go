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
	user_id := c.GetSession("user_id")
	if user_id != nil {
		resp["code"] = models.RECODE_OK
		resp["meg"] = models.ReCodeText(models.RECODE_OK)
		resp["cache"] = "缓存"
		return
	}

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
		beego.Info(err)
		resp["code"] = models.RECODE_DBERR
		resp["meg"] = models.ReCodeText(models.RECODE_DBERR)
		return
	}

	c.SetSession("user_id", user.Id)

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

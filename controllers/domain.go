package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"nweb/models"
	"strconv"
)

type DomainController struct {
	BaseController
}

func (c *DomainController) GetList() {
	//数据json返回
	resp := make(map[string]interface{})
	defer c.Read(resp)

	username:= c.GetSession("username")
	if username == nil {
		resp["code"] = models.RECODE_SESSIONERR
		resp["meg"] = models.ReCodeText(models.RECODE_SESSIONERR)
		return
	}

	o := orm.NewOrm()
	domains := []*models.Domain{}
	qs := o.QueryTable("domain")
	_, err := qs.Filter("User__name", username).All(&domains)
	if err != nil {
		resp["code"] = models.RECODE_DBERR
		resp["meg"] = models.ReCodeText(models.RECODE_DBERR)
		return
	}
	resp["code"] = models.RECODE_OK
	resp["meg"] = models.ReCodeText(models.RECODE_OK)
	resp["data"] = domains
}

func (c *DomainController) PostAdd() {
	resp := make(map[string]interface{})
	defer c.Read(resp)

	//获取session
	username := c.GetSession("username")
	if username == nil {
		resp["code"] = models.RECODE_SESSIONERR
		resp["meg"] = models.ReCodeText(models.RECODE_SESSIONERR)
		return
	}

	//获取前端传过来的数据
	domaindata := make(map[string]string)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &domaindata); err != nil {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}

	server_name := domaindata["server_name"]
	port := domaindata["port"]
	root := domaindata["root"]
	status := domaindata["status"]
	logname := domaindata["logname"]

	if server_name == "" || port == "" || root == "" || logname == "" {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}

	o := orm.NewOrm()
	user := models.User{Name: username.(string)}
	if err := o.Read(&user, "Name"); err != nil {
		resp["code"] = models.RECODE_DBERR
		resp["meg"] = models.ReCodeText(models.RECODE_DBERR)
		return
	}

	domain := models.Domain{}
	domain.Sname = server_name
	newport, err := strconv.Atoi(port)
	if  err != nil {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}
	newstatus, err := strconv.Atoi(status)
	if  err != nil {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}
	if err := o.Read(&domain, "Sname"); err == nil {
		resp["code"] = models.RECODE_DATAEXIST
		resp["meg"] = models.ReCodeText(models.RECODE_DATAEXIST)
		return
	}
	domain.Port = int64(newport)
	domain.Status = int64(newstatus)
	domain.Root = root
	domain.Logname = logname
	domain.User = &user

	_, err = o.Insert(&domain)
	if err != nil {
		resp["code"] = models.RECODE_DBERR
		resp["meg"] = models.ReCodeText(models.RECODE_DBERR)
		return
	}


	resp["code"] = models.RECODE_OK
	resp["meg"] = models.ReCodeText(models.RECODE_OK)
	resp["data"] = domain

}

func (c *DomainController) GetDelete() {
	resp := make(map[string]interface{})
	defer c.Read(resp)

	username:= c.GetSession("username")
	if username == nil {
		resp["code"] = models.RECODE_SESSIONERR
		resp["meg"] = models.ReCodeText(models.RECODE_SESSIONERR)
		return
	}
	//获取id
	domain_id := c.Ctx.Input.Param(":id")
	if domain_id == "" {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}
	newdoamin_id , err := strconv.Atoi(domain_id)
	if err != nil {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}

	o := orm.NewOrm()
	user := models.User{Name:username.(string)}
	err = o.Read(&user, "Name")
	if err != nil {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}
	domain := models.Domain{}
	domain.User = &user
	domain.Id = int64(newdoamin_id)
	err = o.Read(&domain, "Id", "User")
	if err != nil {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}
	num, err := o.Delete(&domain)
	if err != nil {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}
	if num == 0 {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}
	resp["code"] = models.RECODE_OK
	resp["meg"] = models.ReCodeText(models.RECODE_OK)
}


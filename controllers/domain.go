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


//列出所有域名信息
func (c *DomainController) GetListDomain() {
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
	var domains []models.Domain
	qs := o.QueryTable("domain")
	_, err := qs.Filter("User__name", username).All(&domains)
	if err != nil {
		resp["code"] = models.RECODE_DBERR
		resp["meg"] = models.ReCodeText(models.RECODE_DBERR)
		return
	}
	resp["code"] = models.RECODE_OK
	resp["meg"] = models.ReCodeText(models.RECODE_OK)
	tmpList := make(map[int]interface{})
	listData := make(map[string]interface{})
	for index, domain := range domains {
		listData["id"] = domain.Id
		listData["port"] = domain.Port
		listData["server_name"] = domain.Sname
		listData["root"] = domain.Root
		listData["log_name"] = domain.Logname
		listData["status"] = domain.Status
		tmpList[index] = listData
	}
	resp["data"] = tmpList
}

//添加域名信息
func (c *DomainController) PostAddDomain() {
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

	//实现文件写入文件
	//t,_ := template.ParseFiles("template/vhost.tpl")
	//actor := make(map[string]interface{})
	//actor["Port"] = domain.Port
	//actor["Server_name"] = domain.Sname
	//actor["Root"] = domain.Root
	//actor["Logname"] = domain.Logname
	//f, _ := os.OpenFile("test.conf", os.O_RDWR|os.O_CREATE, 0766)
	//defer f.Close()
	//writer := bufio.NewWriter(f)
	//defer writer.Flush()
	//t.Execute(writer, actor)
	//t.Execute(os.Stdout, actor)


}

//删除域名信息
func (c *DomainController) GetDeleteDomain() {
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

//停止使用域名
func (c *DomainController) GetDiscontinuationDomain() {
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
	domain.Status = 2
	num, err := o.Update(&domain)
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

//启用/恢复使用域名
func (c *DomainController) GetRecoveryDomain() {
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
	domain.Status = 1
	num, err := o.Update(&domain)
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


//更改域名信息
func (c *DomainController) PostChangeDomain() {
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

	if server_name == "" && port == "" && root == "" && logname == "" {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
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
	user := models.User{Name: username.(string)}
	if err := o.Read(&user, "Name"); err != nil {
		resp["code"] = models.RECODE_DBERR
		resp["meg"] = models.ReCodeText(models.RECODE_DBERR)
		return
	}

	domain := models.Domain{}
	domain.Id = int64(newdoamin_id)
	if err := o.Read(&domain); err != nil {
		resp["code"] = models.RECODE_NODATA
		resp["meg"] = models.ReCodeText(models.RECODE_NODATA)
		return
	}
	if server_name != "" {
		domain.Sname = server_name
	}

	if port != "" {
		newport, err := strconv.Atoi(port)
		if  err != nil {
			resp["code"] = models.RECODE_DATAERR
			resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
			return
		}
		domain.Port = int64(newport)
	}

	if status != "" {
		newstatus, err := strconv.Atoi(status)
		if  err != nil {
			resp["code"] = models.RECODE_DATAERR
			resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
			return
		}
		domain.Status = int64(newstatus)
	}

	if root != "" {
		domain.Root = root
	}

	if logname != "" {
		domain.Logname = logname
	}

	if server_name == "" || port == "" || root == "" || logname == "" || status == "" {
		domain.User = &user
	}

	_, err = o.Update(&domain)
	if err != nil {
		resp["code"] = models.RECODE_DATAERR
		resp["meg"] = models.ReCodeText(models.RECODE_DATAERR)
		return
	}

	resp["code"] = models.RECODE_OK
	resp["meg"] = models.ReCodeText(models.RECODE_OK)
}


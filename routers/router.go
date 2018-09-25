package routers

import (
	"nweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.UserController{})
	beego.Router("/api/v1.0/user/reg", &controllers.UserController{}, "post:PostReg")
	beego.Router("/api/v1.0/user/login", &controllers.UserController{},"post:PostLogin")
	beego.Router("/api/v1.0/user/logout", &controllers.UserController{}, "get:GetLogout")

	beego.Router("/api/v1.0/domain/list", &controllers.DomainController{}, "get:GetList")
	beego.Router("/api/v1.0/domain/add", &controllers.DomainController{}, "post:PostAdd")
	beego.Router("/api/v1.0/domain/delete/:id", &controllers.DomainController{}, "get:GetDelete")
	beego.Router("/api/v1.0/domain/dis", &controllers.DomainController{})
	beego.Router("/api/v1.0/domain/change", &controllers.DomainController{})
}

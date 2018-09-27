package routers

import (
	"nweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.UserController{}, "get:GetIndex")
	beego.Router("/api/v1.0/user/reg", &controllers.UserController{}, "post:PostReg")
	beego.Router("/api/v1.0/user/login", &controllers.UserController{},"post:PostLogin")
	beego.Router("/api/v1.0/user/logout", &controllers.UserController{}, "get:GetLogout")

	beego.Router("/api/v1.0/domain/list", &controllers.DomainController{}, "get:GetListDomain")
	beego.Router("/api/v1.0/domain/add", &controllers.DomainController{}, "post:PostAddDomain")
	beego.Router("/api/v1.0/domain/delete/:id", &controllers.DomainController{}, "get:GetDeleteDomain")
	beego.Router("/api/v1.0/domain/dis/:id", &controllers.DomainController{}, "get:GetDiscontinuationDomain")
	beego.Router("/api/v1.0/domain/rec/:id", &controllers.DomainController{},"get:GetRecoveryDomain")
	beego.Router("/api/v1.0/domain/change/:id", &controllers.DomainController{}, "post:PostChangeDomain")
}

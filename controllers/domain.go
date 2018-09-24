package controllers

type DomainController struct {
	BaseController
}

func (c *DomainController) GetDomain() {
	//数据json返回
	resp := make(map[string]interface{})
	defer c.Read(resp)
}
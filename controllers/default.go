package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
func (c *MainController) TestPost(){
	//发送json ok
	//接收json ？
	//var ob models.Object
	//var err error
	//if err = json.Unmarshal(c.Ctx.Input.RequestBody, &ob); err == nil {
	//	objectid := models.AddOne(ob)
	//	c.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
	//} else {
	//	c.Data["json"] = err.Error()
	//}
	//c.ServeJSON()
	beego.Info("------",c.GetString("name"))
	beego.Info("------",c.GetString("sex"))
	beego.Info("------", string(c.Ctx.Input.RequestBody))//这是获取到的json二进制数据)

	jbody,_ := json.Marshal(string(c.Ctx.Input.RequestBody))
	jmap := c.Ctx.Input.Data()
	name := c.Ctx.Input.Param("name")
	beego.Info("------",jbody)
	beego.Info("------",jmap)
	beego.Info("------",name)
	//还可以直接用结构体封装
	//json.Unmarshal(body, &ob)//解析二进制json，把结果放进ob中
	//user := &User{Id: ob.UserName, Mobile: ob.Mobile}
	c.TplName = "index.tpl"

}
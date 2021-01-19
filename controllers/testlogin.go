package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type TestLoginController struct {
	beego.Controller
}

func (c *TestLoginController) Login() {
	name := c.Ctx.GetCookie("name")
	password := c.Ctx.GetCookie("password")
	if name != "" {

		c.Ctx.WriteString("Username:" + name + ", password:" + password)
	}
	c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_input">
									<input type="text" name="Username" />
									<input type="password" name="Password" />
									<input type="Submit" value="提交" />
									</form>
								</html>`)
}
func (c *TestLoginController) Post() {

}
func (c *TestLoginController) Get() {

}

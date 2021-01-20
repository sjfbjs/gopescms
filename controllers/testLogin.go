package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type TestLoginController struct {
	beego.Controller
}

func (c *TestLoginController) Login() {
	name := c.Ctx.GetCookie("name")
	password := c.Ctx.GetCookie("password")
	//name := c.GetSession("name")
	//password := c.GetSession("password")
	if name != "" {
		c.Ctx.WriteString("Username:" + name + ", password:" + password)
	} else {
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_login" method="POST">
									<input type="text" name="Username" />
									<input type="password" name="pwd" />
									<input type="Submit" value="提交" />
									</form>
								</html>`)
	}
}

type UserInfo struct {
	Username string
	Password string `form:"pwd"`
}

func (c *TestLoginController) Post() {
	u := UserInfo{}
	if err := c.ParseForm(&u); err != nil {
		//全局捕获异常然后退出
		fmt.Println(err)
	}

	c.Ctx.SetCookie("name", u.Username, 100, "/")
	c.Ctx.SetCookie("password", u.Password, 100, "/")
	c.SetSession("name", u.Username)
	c.SetSession("password", u.Password)
	c.Ctx.WriteString("username:" + u.Username + ",password:" + u.Password)

}
func (c *TestLoginController) Get() {

}

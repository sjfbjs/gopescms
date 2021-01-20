package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type User struct {
	Username string
	Password string
}

type TestInputController struct {
	beego.Controller
}

func (c *TestInputController) Post() {

	u := User{}
	if err := c.ParseForm(&u); err != nil {
		//全局捕获异常然后退出
		fmt.Println(err)
	}
	c.Ctx.WriteString("username:" + u.Username + ",password:" + u.Password)

}
func (c *TestInputController) Get() {
	//id :=c.GetString("id")
	//c.Ctx.WriteString("<html>"+id+"<br />")
	//name := c.GetString("name")
	//// name := c.Input().Get("name")
	//c.Ctx.WriteString(name+"</html>")
	name := c.GetSession("name")
	password := c.GetSession("password")
	//if name != nil {
	if nameString, ok := name.(string); ok && nameString != "" {
		c.Ctx.WriteString("Username:" + name.(string) + ", password:" + password.(string))
	} else {
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_input"  method="POST">
									<input type="text" name="Username" />
									<input type="password" name="Password" />
									<input type="Submit" value="提交" />
									</form>
								</html>`)
	}
}

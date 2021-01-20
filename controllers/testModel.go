package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"
)

type TestModelController struct {
	beego.Controller
}

// my_user 表名，有映射关系
// AxxByy  =>   axx_byy
type MyUser struct {
	Id       int64
	Username string
	Password string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.1.145:3306)/bee?charset=utf8", 30)
	orm.RegisterModel(new(MyUser))
}
func (c *TestModelController) Get() {
	//orm.RegisterModel(new(MyUser))
	//orm.RegisterDataBase("default","mysql","root:root@tcp(192.168.1.145:3306)/bee?charset=utf8",30)
	//insert
	o := orm.NewOrm()
	user := MyUser{Username: "shisuan", Password: "xxasdsd"}
	id, err := o.Insert(&user)
	if err != nil {

		fmt.Println(id, err)
		c.Ctx.WriteString(cast.ToString(id))
	} else {
		c.Ctx.WriteString(cast.ToString(id))

	}

	//update
	//delete
	//read one
	//read all

}

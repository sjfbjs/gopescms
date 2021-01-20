package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"gopescms/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test_input", &controllers.TestInputController{}, "get:Get;post:Post")
	beego.Router("/test_login", &controllers.TestLoginController{}, "get:Login;post:Post")
	beego.Router("/test_model", &controllers.TestModelController{}, "get:Get")

}

package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1", &controllers.MdtController{})
}

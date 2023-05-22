package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["server/controllers:MdtController"] = append(beego.GlobalControllerRouter["server/controllers:MdtController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["server/controllers:MdtController"] = append(beego.GlobalControllerRouter["server/controllers:MdtController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["server/controllers:MdtController"] = append(beego.GlobalControllerRouter["server/controllers:MdtController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:SerialNumber",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["server/controllers:MdtController"] = append(beego.GlobalControllerRouter["server/controllers:MdtController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:SerialNumber",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["server/controllers:MdtController"] = append(beego.GlobalControllerRouter["server/controllers:MdtController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:SerialNumber",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}

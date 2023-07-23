package routers

import (
	"methyl-headimage/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/upload", &controllers.UploadImageController{})
	beego.Router("/getpath/all", &controllers.GetImagePathController{})
	beego.Router("/getpath/inuse", &controllers.GetImagePathController{})
	beego.Router("/useimg", &controllers.UseHeadImageController{})
}

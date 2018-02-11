package routers

import (
	"ss-backend/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/product",
			&controllers.ProductController{},
			"get:Get",
		),
		beego.NSRouter("/product",
			&controllers.ProductController{},
			"post:Post",
		),
		beego.NSRouter("/pemesanan",
			&controllers.PemesananController{},
			"get:Get",
		),
		beego.NSRouter("/pemesanan",
			&controllers.PemesananController{},
			"post:Post",
		),
		beego.NSRouter("/pemesanan",
			&controllers.PemesananController{},
			"put:Put",
		),
		// beego.NSRouter("/pemasukan",
		// 	&controllers.UserController{},
		// 	"post:Post",
		// ),
		// beego.NSRouter("/pengeluaran",
		// 	&controllers.UserController{},
		// 	"get:Get",
		// ),
	)
	beego.AddNamespace(ns)

}

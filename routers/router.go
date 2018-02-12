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
		beego.NSRouter("/penjualan",
			&controllers.PenjualanController{},
			"get:Get",
		),
		beego.NSRouter("/penjualan",
			&controllers.PenjualanController{},
			"post:Post",
		),
		beego.NSRouter("/laporan_nilai_barang",
			&controllers.ReportController{},
			"get:GetNilaiBarang",
		),
		beego.NSRouter("/laporan_penjualan",
			&controllers.ReportController{},
			"get:GetPenjualan",
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

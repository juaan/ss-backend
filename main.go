package main

import (
	_ "ss-backend/routers"

	"ss-backend/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	registerSQLite()
}

func main() {
	beego.Run()
}

func registerSQLite() {
	errDriver := orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if errDriver != nil {
		beego.Warning("error while register driver", errDriver)
	}

	errDB := orm.RegisterDataBase("default", "sqlite3", "file:data.db")
	if errDB != nil {
		beego.Warning("error while register DB", errDB)
	}

	models.RegisterModel()
	models.CreateTableProduct()
	models.CreateTablePemesanan()

	models.MigrateDataProduct()

}

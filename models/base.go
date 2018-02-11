package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// RegisterModel to register table
func RegisterModel() {
	orm.RegisterModel(new(Product))
}

// CreateTableProduct create table product
func CreateTableProduct() {
	o := orm.NewOrm()
	o.Using("default")
	var prod Product

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		prod.TableName(),
		"(",
		"id INTEGER PRIMARY KEY,",
		"sku TEXT,",
		"nama_item TEXT,",
		"ukuran TEXT,",
		"warna TEXT,",
		"jumlah INTEGER);",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	res, err := o.Raw(sql).Exec()

	if err != nil {
		beego.Warning("error creating table product", err)
	}

	beego.Debug(res)

}

// MigrateDataProduct ...
func MigrateDataProduct() {
	var products []Product
	o := orm.NewOrm()

	gopath := os.Getenv("GOPATH")
	fl := gopath + "/src/ss-backend/seeders/data_item.json"

	raw, err := ioutil.ReadFile(fl)
	if err != nil {
		beego.Warning("failed read file seeder", err)
	}
	err = json.Unmarshal(raw, &products)
	if err != nil {
		beego.Warning("failed unmarshall seeders", err)
	}

	cnt, errMulti := o.InsertMulti(len(products), products)
	beego.Debug(cnt, errMulti)

}

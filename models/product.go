package models

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Product is struct for product in inventory
type Product struct {
	ID       int    `json:"id" orm:"column(id);auto"`
	SKU      string `json:"sku" orm:"column(sku)"`
	NamaItem string `json:"nama_item" orm:"column(nama_item)"`
	Ukuran   string `json:"ukuran" orm:"column(ukuran)"`
	Warna    string `json:"warna" orm:"column(warna)"`
	Jumlah   int    `json:"jumlah" orm:"column(jumlah)"`
}

// TableName return the table name
func (p *Product) TableName() string {
	return "product"
}

// Get All product ...
func (p *Product) Get() []Product {
	var products []Product
	o := orm.NewOrm()
	qb := []string{
		"SELECT *",
		"FROM",
		p.TableName(),
	}
	sql := strings.Join(qb, " ")

	count, err := o.Raw(sql).QueryRows(&products)
	if err != nil {
		beego.Warning("Failed get all data product", err)
	}
	beego.Debug("jumlah data = ", count)
	return products
}

package models

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Product is struct for product in inventory
type Product struct {
	ID       int64  `json:"id" orm:"column(id);auto"`
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
func (p *Product) GetAll() ([]Product, error) {
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
		return []Product{}, err
	}
	beego.Debug("jumlah data = ", count)
	return products, nil
}

// AddProduct New product ...
func (p *Product) AddProduct() error {
	o := orm.NewOrm()

	id, err := o.Insert(p)
	if err != nil {
		beego.Debug("error insert", err)
		return err
	}
	p.ID = id

	return nil
}

// UpdateProduct used for update product name , sku, ukuran, warna
func (p *Product) UpdateProduct() error {
	o := orm.NewOrm()
	qb, errQB := orm.NewQueryBuilder("mysql")
	if errQB != nil {
		beego.Warning("Query builder failed")
		beego.Warning(errQB)
		return errQB
	}

	qb.Update(p.TableName()).
		Set(
			"sku = ?",
			"nama_item = ? ",
			"ukuran = ? ",
			"warna = ? ",
			"jumlah = ?").Where("sku = ? ")
	sql := qb.String()

	res, errSQL := o.Raw(sql, p.SKU, p.NamaItem, p.Ukuran, p.Warna,
		p.Jumlah).Exec()
	if errSQL != nil {
		beego.Debug("error while updating product")
		beego.Debug(errSQL)
		return errSQL
	}
	_, errRow := res.RowsAffected()
	if errRow != nil {
		beego.Debug("error get rows affected")
		beego.Debug(errRow)
		return errRow
	}

	return nil
}

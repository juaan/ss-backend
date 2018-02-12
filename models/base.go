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
	orm.RegisterModel(new(Pemesanan))
	orm.RegisterModel(new(Penjualan))

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

// CreateTablePemesanan create table product
func CreateTablePemesanan() {
	o := orm.NewOrm()
	o.Using("default")
	var order Pemesanan

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		order.TableName(),
		"(",
		"id INTEGER PRIMARY KEY,",
		"sku TEXT,",
		"nama_item TEXT,",
		"no_kwitansi TEXT,",
		"jumlah_pesanan INTEGER,",
		"jumlah_diterima INTEGER,",
		"harga DECIMAL(17,5),",
		"catatan TEXT,",
		"waktu TIMESTAMP DEFAULT CURRENT_TIMESTAMP,",
		"total DECIMAL(17,5),",
		"status TEXT);",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	res, err := o.Raw(sql).Exec()

	if err != nil {
		beego.Warning("error creating table pemesanan", err)
	}

	beego.Debug(res)

}

// CreateTablePenjualan create table product
func CreateTablePenjualan() {
	o := orm.NewOrm()
	o.Using("default")
	var sell Penjualan

	qb := []string{
		"CREATE TABLE IF NOT EXISTS",
		sell.TableName(),
		"(",
		"id INTEGER PRIMARY KEY,",
		"id_pesanan TEXT,",
		"sku TEXT,",
		"nama_item TEXT,",
		"jumlah_keluar INTEGER,",
		"harga_jual DECIMAL(17,5),",
		"catatan TEXT,",
		"waktu TIMESTAMP DEFAULT CURRENT_TIMESTAMP,",
		"total DECIMAL(17,5));",
	}

	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	res, err := o.Raw(sql).Exec()

	if err != nil {
		beego.Warning("error creating table penjualan", err)
	}

	beego.Debug(res)

}

// MigrateDataProduct ...
func MigrateDataProduct(param string) {
	o := orm.NewOrm()

	gopath := os.Getenv("GOPATH")

	if param == "product" {
		var products []Product
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

	} else if param == "order" {
		var order []Pemesanan
		fl := gopath + "/src/ss-backend/seeders/data_pemesanan.json"

		raw, err := ioutil.ReadFile(fl)
		if err != nil {
			beego.Warning("failed read file order seeder", err)
		}

		err = json.Unmarshal(raw, &order)
		if err != nil {
			beego.Warning("failed unmarshall order seeders", err)
		}

		cnt, errMulti := o.InsertMulti(len(order), order)
		beego.Debug(cnt, errMulti)
	} else if param == "sell" {
		var sell []Penjualan
		fl := gopath + "/src/ss-backend/seeders/data_penjualan.json"

		raw, err := ioutil.ReadFile(fl)
		if err != nil {
			beego.Warning("failed read file sell seeder", err)
		}

		err = json.Unmarshal(raw, &sell)
		if err != nil {
			beego.Warning("failed unmarshall sell seeders", err)
		}

		cnt, errMulti := o.InsertMulti(len(sell), sell)
		beego.Debug(cnt, errMulti)
	}

}

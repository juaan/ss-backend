package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type (
	// Penjualan is struct for Barang kluar
	Penjualan struct {
		ID           int64     `json:"id" orm:"column(id);auto"`
		IDPesanan    string    `json:"id_pesanan" orm:"column(id_pesanan)"`
		SKU          string    `json:"sku" orm:"column(sku)"`
		NamaItem     string    `json:"nama_item" orm:"column(nama_item)"`
		JumlahKeluar int64     `json:"jumlah_keluar" orm:"column(jumlah_keluar)"`
		HargaJual    float64   `json:"harga_jual" orm:"column(harga_jual)"`
		Catatan      string    `json:"catatan" orm:"column(catatan)"`
		Waktu        time.Time `json:"waktu" orm:"column(waktu);auto_now_add;type(datetime)"`
		Total        float64   `json:"total" orm:"column(total)"`
	}

	// RequestUpdate ...
	RequestUpdate struct {
		Jumlah  int64  `json:"jumlah"`
		ID      int64  `json:"id"`
		Catatan string `json:"catatan"`
	}
)

// TableName return the table name
func (p *Penjualan) TableName() string {
	return "penjualan"
}

// GetAll record barang masuk ...
func (p *Penjualan) GetAll(query RequestGet) ([]Penjualan, error) {
	var brKeluar []Penjualan
	o := orm.NewOrm()
	like := "%" + query.Query + "%"
	qb := []string{
		"SELECT *",
		"FROM",
		p.TableName(),
		"WHERE (waktu >= ? AND waktu <= ?)",
		"AND nama_item LIKE ?",
	}
	sql := strings.Join(qb, " ")

	count, err := o.Raw(sql, query.FromDate, query.ToDate,
		like).QueryRows(&brKeluar)
	if err != nil {
		beego.Warning("Failed get all data product", err)
		return []Penjualan{}, err
	}
	beego.Debug("jumlah data = ", count)
	return brKeluar, nil
}

// AddPenjualan untuk memasukkan record barang masuk ...
func (p *Penjualan) AddPenjualan() error {
	o := orm.NewOrm()

	p.Total = float64(p.JumlahKeluar) * p.HargaJual

	id, err := o.Insert(p)
	if err != nil {
		beego.Debug("error insert", err)
		return err
	}
	p.ID = id

	// Update Product
	var req = RequestUpdate{
		Jumlah:  p.JumlahKeluar,
		ID:      id,
		Catatan: p.Catatan,
	}

	errUpdateProd := updateProduct(req, p.SKU, "jual")
	if errUpdateProd != nil {
		beego.Warning("ERror update product", errUpdateProd)
		return errUpdateProd
	}
	// Update Product

	return nil
}

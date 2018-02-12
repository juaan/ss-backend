package models

import (
	helper "ss-backend/helpers"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// ReportNilaiBarang is struct for product in inventory
type ReportNilaiBarang struct {
	ID        int64  `json:"id" orm:"column(id);auto"`
	SKU       string `json:"sku" orm:"column(sku)"`
	NamaItem  string `json:"nama_item" orm:"column(nama_item)"`
	Jumlah    int    `json:"jumlah" orm:"column(jumlah)"`
	HargaBeli int64  `json:"harga_beli" orm:"column(harga_beli)"`
	Total     int64  `json:"total" orm:"column(total)"`
}

// ReportPenjualan is struct for product in inventory
type ReportPenjualan struct {
	IDPesanan string    `json:"id_pesanan" orm:"column(id_pesanan)"`
	SKU       string    `json:"sku" orm:"column(sku)"`
	NamaItem  string    `json:"nama_item" orm:"column(nama_item)"`
	Jumlah    int       `json:"jumlah" orm:"column(jumlah)"`
	HargaJual int64     `json:"harga_jual" orm:"column(harga_jual)"`
	HargaBeli int64     `json:"harga_beli" orm:"column(harga_beli)"`
	Laba      int64     `json:"laba" orm:"column(laba)"`
	Total     int64     `json:"total" orm:"column(total)"`
	Waktu     time.Time `json:"waktu" orm:"column(waktu);type(datetime)"`
}

// TableName return the table name
// func (p *Product) TableName() string {
// 	return "product"
// }

// GetAll ReportNilaiBarang ...
func (p *ReportNilaiBarang) GetAll() ([]ReportNilaiBarang, error) {
	// var products Product
	var rep []ReportNilaiBarang
	o := orm.NewOrm()
	qb := []string{
		"SELECT prod.id, prod.sku, prod.nama_item, prod.jumlah," +
			" CAST(AVG(ord.harga) as INT) as harga_beli," +
			" CAST(prod.jumlah*AVG(ord.harga) as INT) AS total ",
		"FROM product prod",
		"LEFT JOIN pemesanan ord",
		"ON prod.sku=ord.sku",
		"GROUP BY prod.id, prod.sku ",
	}
	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	count, err := o.Raw(sql).QueryRows(&rep)
	if err != nil {
		beego.Warning("Failed get all data product", err)
		return []ReportNilaiBarang{}, err
	}
	beego.Debug("jumlah data = ", count)
	return rep, nil
}

// GetAll for REPORT laba
func (l *ReportPenjualan) GetAll(query RequestGet) ([]ReportPenjualan, error) {
	// var products Product
	var rep []ReportPenjualan
	o := orm.NewOrm()
	like := "%" + query.Query + "%"

	qb := []string{
		"SELECT se.waktu, se.id_pesanan, prod.sku, prod.nama_item, (se.jumlah_keluar) AS jumlah," +
			" se.harga_jual, " +
			" CAST((se.jumlah_keluar*se.harga_jual) as INT) AS total," +
			" CAST(AVG(ord.harga) as INT) as harga_beli," +
			" CAST((se.jumlah_keluar*se.harga_jual) - (se.jumlah_keluar*AVG(ord.harga)) as INT) AS laba",
		"FROM product prod",
		"LEFT JOIN pemesanan ord",
		"ON prod.sku=ord.sku",
		"LEFT JOIN penjualan se",
		"ON ord.sku=se.sku",
		"WHERE (se.waktu >= ? AND se.waktu <= ?)",
		"AND se.nama_item LIKE ?",
		"GROUP BY se.waktu, prod.sku ",
	}
	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	count, err := o.Raw(sql, query.FromDate, query.ToDate,
		like).QueryRows(&rep)
	if err != nil {
		beego.Warning("Failed get all data product", err)
		return []ReportPenjualan{}, err
	}
	beego.Debug("jumlah data = ", count)
	return rep, nil
}

// GetAllAndWriteCSV ReportNilaiBarang ...
func (p *ReportNilaiBarang) GetAllAndWriteCSV(path string) error {
	// var products Product
	var rep []ReportNilaiBarang
	o := orm.NewOrm()
	qb := []string{
		"SELECT prod.id, prod.sku, prod.nama_item, prod.jumlah," +
			" CAST(AVG(ord.harga) as INT) as harga_beli," +
			" CAST(prod.jumlah*AVG(ord.harga) as INT) AS total ",
		"FROM product prod",
		"LEFT JOIN pemesanan ord",
		"ON prod.sku=ord.sku",
		"GROUP BY prod.id, prod.sku ",
	}
	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	count, err := o.Raw(sql).QueryRows(&rep)
	if err != nil {
		beego.Warning("Failed get all data product", err)
		return err
	}
	beego.Debug("jumlah data = ", count)

	// Write CSV
	p.WriteCsv(path, rep)
	//

	return nil
}

// GetAllAndWriteCSV for REPORT laba
func (l *ReportPenjualan) GetAllAndWriteCSV(query RequestGet,
	path string) error {

	var rep []ReportPenjualan
	o := orm.NewOrm()
	like := "%" + query.Query + "%"

	qb := []string{
		"SELECT se.waktu, se.id_pesanan, prod.sku, prod.nama_item, (se.jumlah_keluar) AS jumlah," +
			" se.harga_jual, " +
			" CAST((se.jumlah_keluar*se.harga_jual) as INT) AS total," +
			" CAST(AVG(ord.harga) as INT) as harga_beli," +
			" CAST((se.jumlah_keluar*se.harga_jual) - (se.jumlah_keluar*AVG(ord.harga)) as INT) AS laba",
		"FROM product prod",
		"LEFT JOIN pemesanan ord",
		"ON prod.sku=ord.sku",
		"LEFT JOIN penjualan se",
		"ON ord.sku=se.sku",
		"WHERE (se.waktu >= ? AND se.waktu <= ?)",
		"AND se.nama_item LIKE ?",
		"GROUP BY se.waktu, prod.sku ",
	}
	sql := strings.Join(qb, " ")
	beego.Debug(sql)
	count, err := o.Raw(sql, query.FromDate, query.ToDate,
		like).QueryRows(&rep)
	if err != nil {
		beego.Warning("Failed get all data product", err)
		return err
	}
	beego.Debug("jumlah data = ", count)

	// Write CSV
	l.WriteCsv(path, rep)
	//

	return nil
}

// WriteCsv for nilai barang
func (p *ReportNilaiBarang) WriteCsv(path string,
	res []ReportNilaiBarang) error {

	w, err := helper.NewCsvWriter(path)
	if err != nil {
		beego.Debug(err)
		return err
	}

	w.Write([]string{
		"sku",
		"nama_item",
		"jumlah",
		"harga_beli",
		"total",
	})

	count := len(res)

	for i := 0; i < count; i++ {

		w.Write([]string{
			res[i].SKU,
			res[i].NamaItem,
			strconv.Itoa(res[i].Jumlah),
			strconv.Itoa(int(res[i].HargaBeli)),
			strconv.Itoa(int(res[i].Total)),
		})
	}

	w.Flush()
	return nil
}

// WriteCsv for nilai barang
func (l *ReportPenjualan) WriteCsv(path string, res []ReportPenjualan) error {

	w, err := helper.NewCsvWriter(path)
	// CheckErr("err@writeCsv", err)
	if err != nil {
		beego.Debug(err)
		return err
	}

	w.Write([]string{
		"id_pesanan",
		"waktu",
		"sku",
		"nama_item",
		"jumlah",
		"harga_jual",
		"total",
		"harga_beli",
		"laba",
	})

	count := len(res)

	for i := 0; i < count; i++ {
		w.Write([]string{
			res[i].IDPesanan,
			res[i].Waktu.String(),
			res[i].SKU,
			res[i].NamaItem,
			strconv.Itoa(res[i].Jumlah),
			strconv.Itoa(int(res[i].HargaJual)),
			strconv.Itoa(int(res[i].Total)),
			strconv.Itoa(int(res[i].HargaBeli)),
			strconv.Itoa(int(res[i].Laba)),
		})

	}
	w.Flush()

	return nil
}

package models

import (
	"testing"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/mattn/go-sqlite3"
	. "github.com/smartystreets/goconvey/convey"
)

func initPg() {
	errDriver := orm.RegisterDriver("sqlite3", orm.DRSqlite)
	if errDriver != nil {
		beego.Warning("error while register driver", errDriver)
	}

	errDB := orm.RegisterDataBase("default", "sqlite3", "file:inventory-test.db")
	if errDB != nil {
		beego.Warning("error while register DB", errDB)
	}

	RegisterModel()
	CreateTablePemesanan()
	CreateTableProduct()
	CreateTablePenjualan()
}

func init() {
	initPg()
}

func TestGetPemesanan(t *testing.T) {
	ResetDB()
	MigrateDataProduct("order")
	var p Pemesanan
	var query = RequestGet{
		FromDate: "2018-02-10",
		ToDate:   "2018-02-12",
	}

	res, err := p.GetAll(query)

	Convey("Get Pemesanan \n", t, func() {
		Convey("Err should nil", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Res length > 0", func() {
			So(len(res), ShouldNotEqual, 0)
		})
	})

}

func TestGetPenjualan(t *testing.T) {
	ResetDB()
	MigrateDataProduct("sell")
	var p Penjualan
	var query = RequestGet{
		FromDate: "2018-02-10",
		ToDate:   "2018-02-12",
	}

	res, err := p.GetAll(query)

	Convey("Get Penjualan \n", t, func() {
		Convey("Err should nil", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Res length > 0", func() {
			So(len(res), ShouldNotEqual, 0)
		})
	})

}

func TestGetProduct(t *testing.T) {
	ResetDB()
	MigrateDataProduct("product")
	MigrateDataProduct("order")
	MigrateDataProduct("sell")

	var p Product

	res, err := p.GetAll()

	Convey("Get Product \n", t, func() {
		Convey("Err should nil", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Res length > 0", func() {
			So(len(res), ShouldNotEqual, 0)
		})
	})

}

func TestGetReportNilaiBarang(t *testing.T) {
	ResetDB()
	MigrateDataProduct("product")
	MigrateDataProduct("order")

	var p ReportNilaiBarang

	res, err := p.GetAll()

	Convey("Get Report Nilai Barang \n", t, func() {
		Convey("Err should nil", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Res length > 0", func() {
			So(len(res), ShouldNotEqual, 0)
		})
	})

}

func TestGetReportPenjualan(t *testing.T) {
	ResetDB()
	MigrateDataProduct("product")
	MigrateDataProduct("sell")
	MigrateDataProduct("order")

	var p ReportPenjualan
	var query = RequestGet{
		FromDate: "2018-02-10",
		ToDate:   "2018-02-12",
	}

	res, err := p.GetAll(query)

	Convey("Get Report Penjualan \n", t, func() {
		Convey("Err should nil", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("Res length > 0", func() {
			So(len(res), ShouldNotEqual, 0)
		})
	})

}

func TestAddPemesanan(t *testing.T) {
	ResetDB()
	MigrateDataProduct("product")

	var p = Pemesanan{
		SKU:            "SSI-D00791015-LL-BWH",
		NamaItem:       "Zalekia Plain Casual Blouse",
		NoKwitansi:     "20180102-6974222",
		JumlahDiterima: 20,
		JumlahPesanan:  20,
		Harga:          79000,
		Catatan:        "2017/11/02 terima 50",
	}

	err := p.AddPemesanan()

	Convey("Add Pemesanan \n", t, func() {
		Convey("Err should nil", func() {
			So(err, ShouldEqual, nil)
		})
		Convey("status pemesana should equal sukses", func() {
			So(p.Status, ShouldEqual, "sukses")
		})
	})

}

func TestAddPenjualan(t *testing.T) {
	ResetDB()
	MigrateDataProduct("product")

	var p = Penjualan{
		SKU:          "SSI-D00791015-LL-BWH",
		NamaItem:     "Zalekia Plain Casual Blouse",
		JumlahKeluar: 2,
		HargaJual:    179000,
		Catatan:      "Pesanan ID-20180109-8537240",
	}

	err := p.AddPenjualan()

	Convey("Add Pemesanan \n", t, func() {
		Convey("Err should nil", func() {
			So(err, ShouldEqual, nil)
		})
	})

}

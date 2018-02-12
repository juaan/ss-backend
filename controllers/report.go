package controllers

import (
	"os"
	"ss-backend/models"
	"time"

	"github.com/astaxie/beego"
)

type (
	// ReportController ...
	ReportController struct {
		beego.Controller
	}
)

// GetNilaiBarang to get all nilai barang
func (c *ReportController) GetNilaiBarang() {
	var resp RespData
	var rep models.ReportNilaiBarang

	res, errGet := rep.GetAll()
	if errGet != nil {
		resp.Error = errGet
	} else {
		resp.Body = res
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		panic("ERROR OUTPUT JSON LEVEL MIDDLEWARE")
	}
	// c.TplName = "index.tpl"
}

// GetPenjualan is controller for get data penjualan
func (c *ReportController) GetPenjualan() {
	var resp RespData
	var rep models.ReportPenjualan

	var reqDt = models.RequestGet{
		FromDate: c.Ctx.Input.Query("fromDate"),
		ToDate:   c.Ctx.Input.Query("toDate"),
		Query:    c.Ctx.Input.Query("query"),
	}

	res, errGet := rep.GetAll(reqDt)
	if errGet != nil {
		resp.Error = errGet
	} else {
		resp.Body = res
	}

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		panic("ERROR OUTPUT JSON LEVEL MIDDLEWARE")
	}
	// c.TplName = "index.tpl"
}

// GetNilaiBarangCSV to get all nilai barang
func (c *ReportController) GetNilaiBarangCSV() {

	var rep models.ReportNilaiBarang

	// path for directory
	gopath := os.Getenv("GOPATH")
	dt := time.Now()
	nameFl := "report_nilai_barang" + dt.Format("20060102")
	path := gopath + "/src/ss-backend/storages/" +
		nameFl + ".csv"

	errGet := rep.GetAllAndWriteCSV(path)

	if errGet != nil {
		beego.Debug("Error get csv", errGet)
	}

	c.Ctx.Output.Download(path, nameFl+".csv")

	// c.TplName = "index.tpl"
}

// GetPenjualanCSV is controller for get data penjualan
func (c *ReportController) GetPenjualanCSV() {
	var rep models.ReportPenjualan

	var reqDt = models.RequestGet{
		FromDate: c.Ctx.Input.Query("fromDate"),
		ToDate:   c.Ctx.Input.Query("toDate"),
		Query:    c.Ctx.Input.Query("query"),
	}
	gopath := os.Getenv("GOPATH")
	dt := time.Now()
	nameFl := "report_penjualan" + dt.Format("20060102")
	path := gopath + "/src/ss-backend/storages/" +
		nameFl + ".csv"

	errGet := rep.GetAllAndWriteCSV(reqDt, path)
	if errGet != nil {
		beego.Debug("Error get csv", errGet)
	}

	c.Ctx.Output.Download(path, nameFl+".csv")
	// c.TplName = "index.tpl"
}

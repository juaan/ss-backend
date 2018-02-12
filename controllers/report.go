package controllers

import (
	"ss-backend/models"

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

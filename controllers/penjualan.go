package controllers

import (
	"encoding/json"
	"ss-backend/models"

	"github.com/astaxie/beego"
)

type (
	// PenjualanController ...
	PenjualanController struct {
		beego.Controller
	}
)

// Get all data product
func (c *PenjualanController) Get() {
	var resp RespData
	var sell models.Penjualan

	var reqDt = models.RequestGet{
		FromDate: c.Ctx.Input.Query("fromDate"),
		ToDate:   c.Ctx.Input.Query("toDate"),
		Query:    c.Ctx.Input.Query("query"),
	}

	beego.Debug(reqDt)
	res, errGet := sell.GetAll(reqDt)
	if errGet != nil {
		resp.Error = errGet
	} else {
		resp.Body = res
	}
	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		panic("ERROR OUTPUT JSON LEVEL MIDDLEWARE")
	}
}

// Post add new sell
func (c *PenjualanController) Post() {
	var resp RespData
	var sell models.Penjualan

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &sell)

	if err != nil {
		beego.Warning("unmarshall req body failed")
	}

	errAdd := sell.AddPenjualan()

	if errAdd != nil {
		resp.Error = errAdd
	} else {
		resp.Body = sell

	}
	err = c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		beego.Warning("failed giving output", err)
	}
	// c.TplName = "index.tpl"
}

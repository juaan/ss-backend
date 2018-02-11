package controllers

import (
	"encoding/json"
	"ss-backend/models"

	"github.com/astaxie/beego"
)

type (
	ProductController struct {
		beego.Controller
	}
	RespData struct {
		// ReqHeader JSONHeaderResp    `json:"rsHeader"`
		Body  interface{} `json:"body"`
		Error error       `json:"error"`
	}
)

// Get all data product
func (c *ProductController) Get() {
	var resp RespData
	var prod models.Product

	res, errGet := prod.GetAll()
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

// Post new porduct
func (c *ProductController) Post() {
	var resp RespData
	var prod models.Product

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &prod)

	if err != nil {
		beego.Warning("unmarshall req body failed")
	}

	errAddProd := prod.AddProduct()
	if errAddProd != nil {
		resp.Error = errAddProd
	} else {
		resp.Body = prod
	}

	err = c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		beego.Warning("failed giving output", err)
	}
	// c.TplName = "index.tpl"
}

// Put for update porduct
func (c *ProductController) Put() {
	var resp RespData
	var prod models.Product

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &prod)

	if err != nil {
		beego.Warning("unmarshall req body failed")
	}

	errAddProd := prod.UpdateProduct()
	if errAddProd != nil {
		resp.Error = errAddProd
	} else {
		resp.Body = prod
	}

	err = c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		beego.Warning("failed giving output", err)
	}
	// c.TplName = "index.tpl"
}

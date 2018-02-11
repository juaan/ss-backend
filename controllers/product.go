package controllers

import (
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

	res := prod.Get()
	resp.Body = res

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		panic("ERROR OUTPUT JSON LEVEL MIDDLEWARE")
	}
	// c.TplName = "index.tpl"
}

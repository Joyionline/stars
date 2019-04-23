package controllers

import (
	"stars/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.StarsService
}

// 获取数据
func (c *AdminController) Get() mvc.Result {

	return
}

func (c *AdminController) GetEdit() mvc.Result {

	return nil
}

func (c *AdminController) PostSave() mvc.Result {

	return nil
}

func (c *AdminController) GetDelete() mvc.Result {

	return nil
}

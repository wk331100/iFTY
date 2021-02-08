package models

import (
	"fmt"
	"github.com/wk331100/iFTY/system/db"
	errors "github.com/wk331100/iFTY/system/error"
	"github.com/wk331100/iFTY/system/helper"
)

type UserModel struct {
	BaseModel
}

func (this *UserModel) Instance() *UserModel {
	this.BaseModel.Table("user")
	return this
}

func (this *UserModel) GetTest() ([]helper.Map,errors.Code) {
	model := this.getInstance(db.SLAVE).Table(this.BaseModel.table)
	fmt.Println("GetTest: ")
	return model.PageSize(4).Page(1).Get()
}



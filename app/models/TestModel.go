package models

import (
	"fmt"
	"github.com/wk331100/iFTY/system/db"
	errors "github.com/wk331100/iFTY/system/error"
	"github.com/wk331100/iFTY/system/helper"
)

type TestModel struct {
	BaseModel
}


func (this *TestModel) Instance() *TestModel {
	this.BaseModel.Table("test")
	return this
}

//负载BaseModel中的List方法
func (this *TestModel) List(where helper.Map) ([]helper.Map,errors.Code) {
	model := this.getInstance(db.SLAVE).Table(this.BaseModel.table)
	if where["ids"] != nil {
		idSlice := helper.Explode(",", where["ids"].(string))
		delete(where, "ids")
		model.WhereIn("id", idSlice)
	}
	fmt.Println("asdasd")
	return model.Where(where).PageSize(2).Page(1).Get()
}
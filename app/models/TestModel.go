package models

import (
	"github.com/wk331100/iFTY/system/db"
	"github.com/wk331100/iFTY/system/helper"
)

type TestModel struct {
	table string
	Connector *db.Mysql
}

func (this *TestModel) getInstance() *db.Mysql{
	this.table = "test"
	if this.Connector == nil || !this.Connector.IsConnected() {
		this.Connector =  new(db.Mysql).Connect()
	}
	return this.Connector
}

func (this *TestModel) Insert(data helper.Map) int {
	return this.getInstance().Table(this.table).Insert(data)
}

func (this *TestModel) Update(data, where helper.Map) int {
	return this.getInstance().Table(this.table).Where(where).Update(data)
}

func (this *TestModel) Delete(data helper.Map) bool {
	return this.getInstance().Table(this.table).Where(data).Delete()
}

func (this *TestModel) List(where helper.Map) []helper.Map {
	model := this.getInstance().Table(this.table)
	if where["ids"] != nil {
		idSlice := helper.Explode(",", where["ids"].(string))
		delete(where, "ids")
		model.WhereIn("id", idSlice)
	}
	return model.Where(where).Get()
}

func (this *TestModel) Info(where helper.Map) helper.Map {
	return this.getInstance().Table(this.table).Where(where).First()
}

func (this *TestModel) Exist(where helper.Map) bool {
	info := this.Info(where)
	if len(info) > 0 {
		return true
	}
	return false
}

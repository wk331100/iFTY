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

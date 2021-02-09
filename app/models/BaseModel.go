package models

import (
	"github.com/wk331100/iFTY/system/db"
	errors "github.com/wk331100/iFTY/system/error"
	"github.com/wk331100/iFTY/system/helper"
)

type BaseModel struct {
	table string
	Master *db.Mysql
	Slave  *db.Mysql
}

func (this *BaseModel) getInstance(cluster string) *db.Mysql{
	switch cluster {
	case db.MASTER:
		if this.Master == nil || !this.Master.IsConnected() {
			this.Master =  new(db.Mysql).Connect()
		}
		return this.Master
	case db.SLAVE:
		if this.Slave == nil || !this.Slave.IsConnected() {
			this.Slave =  new(db.Mysql).ConnectCluster(db.SLAVE)
		}
		return this.Slave
	}
	return this.Master
}

func (this *BaseModel) Table(table string) {
	this.table = table
}

func (this *BaseModel) Insert(data helper.Map) (int,errors.Code) {
	return this.getInstance(db.MASTER).Table(this.table).Insert(data)
}

func (this *BaseModel) Update(data, where helper.Map) (int,errors.Code) {
	return this.getInstance(db.MASTER).Table(this.table).Where(where).Update(data)
}

func (this *BaseModel) Delete(data helper.Map) (bool,errors.Code) {
	return this.getInstance(db.MASTER).Table(this.table).Where(data).Delete()
}

func (this *BaseModel) List(where helper.Map) ([]helper.Map,errors.Code) {
	model := this.getInstance(db.SLAVE).Table(this.table)
	if where["ids"] != nil {
		idSlice := helper.Explode(",", where["ids"].(string))
		delete(where, "ids")
		model.WhereIn("id", idSlice)
	}
	return model.Where(where).Get()
}

func (this *BaseModel) Info(where helper.Map) (helper.Map,errors.Code) {
	return this.getInstance(db.SLAVE).Table(this.table).Where(where).First()
}

func (this *BaseModel) Exist(where helper.Map) (bool,errors.Code) {
	info,errorCode := this.Info(where)
	if len(info) > 0 {
		return true,errors.EMPTY
	}
	return false,errorCode
}


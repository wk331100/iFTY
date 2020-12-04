package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wk331100/iFTY/config"
	"github.com/wk331100/iFTY/system/helper"
)

const MASTER = "master"
const SLAVE = "slave"
const COND_EQ = "="
const COND_LESS = "<"
const COND_LESS_OR_EQ = "<="
const COND_GREATER = ">"
const COND_GREATER_OR_EQ = ">="

type Mysql struct {
	//mysql data source name
	Address string
	Connector       *sql.DB
	TableName 	string
	Filter []Filter
	Column []string
}

type Filter struct {
	key string
	value interface{}
	condition string
}

func (this *Mysql)Connect() *Mysql {
	return this.ConnectCluster(MASTER)
}

func (this *Mysql)ConnectCluster(cluster string) *Mysql {
	mysqlConfig := config.MysqlConfig
	if mysqlConfig[cluster] == nil {
		panic("Error Cluster !" )
	}
	config := mysqlConfig[cluster].(helper.Map)
	this.Address = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config["username"],config["password"],config["host"],config["port"],config["dbname"])
	conn, err := sql.Open("mysql", this.Address)
	this.Connector = conn

	if err != nil || conn.Ping() != nil {
		fmt.Println("Mysql Connecting Failed!")
		return nil
	}
	fmt.Println("Mysql Connected")
	return this
}

func (this *Mysql) IsConnected() bool {
	if this.Connector.Ping() != nil{
		return false
	}
	return true
}

func (this *Mysql)Table(table string) *Mysql {
	this.TableName = table
	return this
}

func (this *Mysql)Insert(insertData helper.Map) int {
	sql := "INSERT INTO " + this.TableName
	if len(insertData) <= 0 {
		return 0
	}

	i, keys := 0, make([]interface{}, len(insertData))
	mark := make([]interface{}, len(insertData))
	vals := make([]interface{}, len(insertData))
	for key, val := range insertData {
		keys[i] = key
		mark[i] = "?"
		vals[i] = val
		i++
	}
	columns := helper.Implode(",", keys)
	marks := helper.Implode(",", mark)
	sql += "(" + columns + ") VALUES (" + marks + ")"

	stmt, _ := this.Connector.Prepare(sql)
	defer stmt.Close()

	ret, err := stmt.Exec(vals...)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return 0
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		return int(LastInsertId)
	}
	return 0;
}

func (this *Mysql)Delete() bool {
	sql := "DELETE FROM `" + this.TableName + "` WHERE "
	if len(this.Filter) < 0 {
		return false
	}
	filter := []interface{}{}
	for _,item := range this.Filter  {
		filter = append(filter, fmt.Sprintf("`%s` %s '%v'",item.key, item.condition, item.value))
	}

	sql += helper.Implode(" AND " , filter)
	fmt.Println(sql)
	_,err := this.Connector.Exec(sql)
	if err != nil {
		fmt.Printf("Delete data error: %v\n", err)
		return false
	}
	return true
}




func (this *Mysql)Where(filter helper.Map) *Mysql {
	for key,val := range filter  {
		filterStruct := Filter{
			key:       key,
			value:     val,
			condition: COND_EQ,
		}
		this.Filter = append(this.Filter, filterStruct)
		fmt.Println(this.Filter)
	}
	return this
}

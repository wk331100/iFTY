package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wk331100/iFTY/system/helper"
)


type Mysql struct {
	//mysql data source name
	Address string
	Connector       *sql.DB
	TableName 	string
}

func (this *Mysql)Connect() *Mysql {
	fmt.Println("Mysql Connecting ……")
	this.Address = "root:Wk331100!@tcp(192.168.126.100:3306)/blog"
	conn, err := sql.Open("mysql", this.Address)
	this.Connector = conn

	if err != nil || conn.Ping() != nil {
		fmt.Println("Mysql Connecting Failed!")
		panic(err)
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
	DB := this.Connect()
	DB.TableName = table
	return DB
}

func (this *Mysql)Insert(insertData helper.Map) int {
	sql := "INSERT INTO " + this.TableName;
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
	fields := helper.Implode(",", keys)
	marks := helper.Implode(",", mark)
	sql += "(" + fields + ") VALUES (" + marks + ")"

	fmt.Println(sql)
	prepare, _ := this.Connector.Prepare(sql)
	defer prepare.Close()

	ret, err := prepare.Exec(vals...)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return 0
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
		return int(LastInsertId)
	}
	return 0;
}

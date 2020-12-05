package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wk331100/iFTY/config"
	"github.com/wk331100/iFTY/system/helper"
	"log"
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
	Column []interface{}
	pageSize int
	page int
}

type Filter struct {
	key string
	value interface{}
	condition string
}

func (this *Mysql) Connect() *Mysql {
	return this.ConnectCluster(MASTER)
}

func (this *Mysql) ConnectCluster(cluster string) *Mysql {
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

func (this *Mysql) Table(table string) *Mysql {
	appConfig := config.AppConfig
	this.TableName = table
	this.page = 1
	this.pageSize = appConfig["pageSize"].(int)
	return this
}

func (this *Mysql) Insert(insertData helper.Map) int {
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

func (this *Mysql) Update(updateData helper.Map) int {
	sql := "UPDATE `" + this.TableName + "`"
	if len(updateData) <= 0 {
		return 0
	}

	setter, setVals := this.parseSet(updateData)
	sql += " SET " + setter

	filter, vals := this.parseFilter()
	sql += " WHERE " + filter

	fmt.Println(sql)
	stmt,_ := this.Connector.Prepare(sql)
	defer stmt.Close()
	vals = helper.ArrayMerge(setVals, vals)
	fmt.Println(vals)
	ret, err := stmt.Exec(vals...)
	if err != nil {
		fmt.Printf("Delete data error: %v\n", err)
		return 0
	}
	affectedNum, _ := ret.RowsAffected()
	return int(affectedNum)
}



func (this *Mysql) Delete() bool {
	sql := "DELETE FROM `" + this.TableName + "` WHERE "
	if len(this.Filter) < 0 {
		return false
	}
	filter, vals := this.parseFilter()
	sql += filter
	fmt.Println(sql)
	stmt,_ := this.Connector.Prepare(sql)
	defer stmt.Close()
	_, err := stmt.Exec(vals...)
	if err != nil {
		fmt.Printf("Delete data error: %v\n", err)
		return false
	}
	return true
}

func (this *Mysql) Get() []helper.Map {
	sql, vals := this.buildQuerySQL()
	stmt,_ := this.Connector.Prepare(sql)
	defer stmt.Close()
	rows, err := stmt.Query(vals...)
	if err != nil {
		fmt.Printf("Select data error: %v\n", err)
		return nil
	}
	return this.parseResult(rows)
}

func (this *Mysql) First() helper.Map {
	this.pageSize = 1
	this.page = 1
	sql, vals := this.buildQuerySQL()
	fmt.Println(vals)
	stmt,_ := this.Connector.Prepare(sql)
	defer stmt.Close()
	fmt.Println("++++++++++++")
	fmt.Println(vals)
	rows, err := stmt.Query(vals...)
	fmt.Println("===========")
	if err != nil {
		fmt.Printf("Select data error: %v\n", err)
		return nil
	}
	results := this.parseResult(rows)
	if len(results) > 0 {
		return results[0]
	}
	return helper.Map{}
}

func (this *Mysql) buildQuerySQL() (string, []interface{}){
	sql := "SELECT "
	if len(this.Column) <= 0{
		sql += " * "
	} else {
		columnStr := helper.Implode(",", this.Column)
		sql += columnStr
	}
	filter,vals := this.parseFilter()
	sql += " FROM `" + this.TableName + "` WHERE " + filter + " " + this.parseLimit()
	fmt.Println(sql)
	return sql, vals
}


func (this *Mysql) parseResult(rows *sql.Rows) []helper.Map {
	cols,_ := rows.Columns()
	vals := make([]interface{}, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range vals{
		scans[i] = &vals[i]
	}
	var result = []helper.Map{}

	for rows.Next()  {
		err := rows.Scan(scans...)
		if err != nil{
			log.Fatalln(err)
		}

		row := helper.Map{}
		for k, v := range vals{
			key := cols[k]
			row[key] = v
		}
		result = append(result, row)
	}
	return result
}


func (this *Mysql) parseLimit() string {
	start := (this.page - 1) * this.pageSize
	return fmt.Sprintf("LIMIT %d,%d", start, this.pageSize)
}

func (this *Mysql) parseFilter() (string,[]interface{}) {
	filter := []interface{}{}
	vals := []interface{}{}
	for _,item := range this.Filter  {
		filter = append(filter, fmt.Sprintf("`%s` %s ?",item.key, item.condition))
		vals = append(vals, item.value)
	}
	return helper.Implode(" AND " , filter), vals
}

func (this *Mysql) parseSet(sets helper.Map) (string,[]interface{}) {
	keys := []interface{}{}
	vals := []interface{}{}
	for key,val := range sets  {
		keys = append(keys, fmt.Sprintf("`%s` = ?", key))
		vals = append(vals, val)
	}
	return helper.Implode(" , " , keys), vals
}

func (this *Mysql) Page (page int) *Mysql {
	this.page = page
	return this
}

func (this *Mysql) PageSize (pageSize int) *Mysql {
	this.pageSize = pageSize
	return this
}

func (this *Mysql) Where(filter helper.Map) *Mysql {
	for key,val := range filter  {
		filterStruct := Filter{
			key:       key,
			value:     val,
			condition: COND_EQ,
		}
		this.Filter = append(this.Filter, filterStruct)
	}
	return this
}

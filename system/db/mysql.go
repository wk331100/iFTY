package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wk331100/iFTY/config"
	errors "github.com/wk331100/iFTY/system/error"
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
const COND_IN = "in"

type Mysql struct {
	Address string
	Connector       *sql.DB
	TableName 	string
	Filter []Where
	Column []interface{}
	pageSize int
	page int
}

type Where struct {
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
	conf := mysqlConfig[cluster].(helper.Map)
	fmt.Println(conf)

	this.Address = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf["username"],conf["password"],conf["host"],conf["port"],conf["dbname"])
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
	fmt.Println(appConfig)
	this.pageSize = appConfig["pageSize"].(int)
	fmt.Println(this.pageSize)
	return this
}

func (this *Mysql) Insert(insertData helper.Map) (int,errors.Code) {
	sql := "INSERT INTO " + this.TableName
	if len(insertData) <= 0 {
		return 0,errors.EMPTY_DATA
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

	fmt.Println(sql)
	stmt, _ := this.Connector.Prepare(sql)
	defer stmt.Close()

	ret, err := stmt.Exec(vals...)
	if err != nil {
		fmt.Printf(errors.GetMessage(errors.INSERT_ERROR, err.Error()))
		return 0, errors.INSERT_ERROR
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		return int(LastInsertId),errors.EMPTY
	}
	return 0,errors.EMPTY
}

func (this *Mysql) Update(updateData helper.Map) (int,errors.Code) {
	sql := "UPDATE `" + this.TableName + "`"
	if len(updateData) <= 0 {
		return 0,errors.EMPTY_DATA
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
		fmt.Printf(errors.GetMessage(errors.UPDATE_ERROR, err.Error()))
		return 0,errors.UPDATE_ERROR
	}
	affectedNum, _ := ret.RowsAffected()
	return int(affectedNum),errors.EMPTY
}



func (this *Mysql) Delete() (bool,errors.Code) {
	sql := "DELETE FROM `" + this.TableName + "` WHERE "
	if len(this.Filter) < 0  {
		return false,errors.EMPTY_DATA
	}
	filter, vals := this.parseFilter()
	sql += filter
	fmt.Println(sql)
	stmt,_ := this.Connector.Prepare(sql)
	defer stmt.Close()
	_, err := stmt.Exec(vals...)
	if err != nil {
		fmt.Printf(errors.GetMessage(errors.DELETE_ERROR, err.Error()))
		return false,errors.DELETE_ERROR
	}
	return true,errors.EMPTY
}

func (this *Mysql) Get() ([]helper.Map,errors.Code) {
	sql, vals := this.buildQuerySQL()
	stmt,_ := this.Connector.Prepare(sql)
	defer stmt.Close()
	rows, err := stmt.Query(vals...)
	if err != nil {
		fmt.Printf(errors.GetMessage(errors.QUERY_ERROR, err.Error()))
		return nil,errors.QUERY_ERROR
	}
	return this.parseResult(rows),errors.EMPTY
}

func (this *Mysql) First() (helper.Map,errors.Code) {
	this.pageSize = 1
	this.page = 1
	sql, vals := this.buildQuerySQL()
	stmt,_ := this.Connector.Prepare(sql)
	defer stmt.Close()
	rows, err := stmt.Query(vals...)
	if err != nil {
		fmt.Printf(errors.GetMessage(errors.QUERY_ERROR, err.Error()))
		return nil,errors.QUERY_ERROR
	}
	results := this.parseResult(rows)
	if len(results) > 0 {
		return results[0],errors.EMPTY
	}
	return nil,errors.EMPTY_DATA
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
	sql += " FROM `" + this.TableName + "`"
	if filter != "" {
		sql += " WHERE " + filter
	}

	sql += " " + this.parseLimit()
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
			switch v.(type) {
			case []uint8:
					v := string(v.([]uint8)[:])
					row[key] = v
			default:
				row[key] = v
			}
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
		if item.condition != COND_IN {
			filter = append(filter, fmt.Sprintf("`%s` %s ?",item.key, item.condition))
			vals = append(vals, item.value)
		} else {
			marks := []string{}
			for _,val :=  range item.value.([]string){
				marks = append(marks, "?")
				vals = append(vals, val)
			}
			filter = append(filter, fmt.Sprintf("`%s` IN("+ helper.ImplodeString(",", marks)+")",item.key))
		}

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
		filterStruct := Where{
			key:       key,
			value:     val,
			condition: COND_EQ,
		}
		this.Filter = append(this.Filter, filterStruct)
	}
	return this
}

func (this *Mysql) WhereIn(key string, InValues []string) *Mysql {
	filterStruct := Where{
		key:       key,
		value:     InValues,
		condition: COND_IN,
	}
	this.Filter = append(this.Filter, filterStruct)
	return this
}

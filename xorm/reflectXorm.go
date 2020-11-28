package main

import (
	//"encoding/json"
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)
import "github.com/go-xorm/xorm"
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/gin-gonic/gin"

func main() {
	testExists()
	var g = gin.Default()
	g.POST("/query", query)
	g.Run(":80")
}

func testFind() {
	engine := getEngine()
	var users []User
	err := engine.NewSession().Table("user").Select("id,name,created_at").Find(&users)
	if err != nil {
		fmt.Println("get err")
		panic(err)
	}

	e := reflect.ValueOf(&users[0]).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Tag.Get("json")
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v %v %v\n", varName, varType, varValue)
	}
}

func getEngine() *xorm.Engine {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	return engine
}

const driverName = "mysql"
const dataSourceName = "root:@tcp(127.0.0.1:3306)/test"

type User struct {
	CreatedAt time.Time `json:"created_at"`
	Id        int       `json:"id"`
	Name      string    `json:"name"`
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Param struct {
	Where  [][]interface{} `json:"where"`
	Fields []string        `json:"fields"`
}

func query(ctx *gin.Context) {
	var param Param
	err := ctx.Bind(&param)
	if err != nil {
		ctx.String(200, "参数解析失败")
		log.Println(err)
		return
	}

	session := getEngine().NewSession()
	defer session.Close()
	if len(param.Fields) > 0 {
		session.Select(strings.Join(param.Fields, ","))
	}
	if len(param.Where) > 0 {
		err := whereHandler(session, param.Where)
		if err != nil {
			ctx.String(200, err.Error())
			return
		}
	}

	var rows []User
	err = session.Find(&rows)
	if err != nil {
		ctx.String(200, err.Error())
		return
	}

	var fieldMap = map[string]bool{}
	for _, field := range param.Fields {
		fieldMap[field] = true
	}

	var data []map[string]interface{}
	for _, row := range rows {
		var d = map[string]interface{}{}
		e := reflect.ValueOf(row)
		for i := 0; i < e.NumField(); i++ {
			varName := e.Type().Field(i).Tag.Get("json")
			if len(fieldMap) != 0 {
				if _, ok := fieldMap[varName]; !ok {
					continue
				}
			}
			varValue := e.Field(i).Interface()
			d[varName] = varValue
		}
		data = append(data, d)
	}

	ctx.JSON(200, Resp{
		Code: 200,
		Msg:  "",
		Data: data,
	})
}

func whereHandler(e *xorm.Session, wheres [][]interface{}) error {
	if len(wheres) == 1 {
		where := wheres[0]
		switch len(where) {
		case 0:
			return nil
		case 1:
			return fmt.Errorf("where参数异常")
		case 2:
			e = e.Where(fmt.Sprintf(" %s = ?", where[0]), where[1])
		case 3:
			e = e.Where(fmt.Sprintf(" %s %s ?", where[0], where[1]), where[2])
		default:
			return fmt.Errorf("where参数异常")
		}
		return nil
	}

	for _, where := range wheres {
		switch len(where) {
		case 0:
			continue
		case 1:
			return fmt.Errorf("where参数异常")
		case 2:
			e = e.Where(fmt.Sprintf(" %s = ?", where[0]), where[1])
		case 3:
			e = e.Where(fmt.Sprintf(" %s %s ?", where[0], where[1]), where[2])
		//case 4:
		//	e = e.Where(fmt.Sprintf("%s %s ?", where[0],where[1]), where[2])
		default:
			return fmt.Errorf("where参数异常")
		}
	}

	return nil
}

func InsertMap()  {
	m := map[string]interface{}{
		"name":"xxx",
	}
	effectedRows,err := getEngine().Table("user").Insert(m)
	log.Printf("effectedRows:%d, err:%s",effectedRows,err)
	os.Exit(0)
}

func UpdateIncr()  {
	//m := map[string]interface{}{
	//	"id":"id+1",
	//}
	effectedRows,err := getEngine().Exec("")
	log.Printf("effectedRows:%d, err:%s",effectedRows,err)
	os.Exit(0)
}

func testExists()  {
	var row = struct {
		Id int
	}{}
	ok,err := getEngine().Table("user").Where("id=100").Get(&row)
	log.Printf("exists:%d, err:%s",ok,err)
	os.Exit(0)
}

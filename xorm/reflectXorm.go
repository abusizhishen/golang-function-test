package main

import (
	"encoding/json"
	//"encoding/json"
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
	"xorm.io/core"
)
import "github.com/go-xorm/xorm"
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/gin-gonic/gin"

const driverName = "mysql"
const dataSourceName = "root:root@tcp(127.0.0.1:3306)/ebike"

func main() {
	//testDelete()
	//testInsertMany()
	//testSession()
	//testQueryIn()
	selectById()
	return
	//var g = gin.Default()
	//g.POST("/query", query)
	////g.Run(":80")
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
		Id int `json:"id"`
		Name string `json:"name"`
		LimitNum int `json:"limit_num"`
	}{}
	ok,err := getEngine().Table("city").Where("id=1").Where("limit_num=6").Where("name='蔡家坡经开区'").Get(&row)
	log.Printf("exists:%v, err:%v",ok,err)
	log.Println(row)
	os.Exit(0)
}

func testDelete()  {
	ok,err := getEngine().Delete(City{Id: "1"})
	log.Printf("exists:%v, err:%v",ok,err)
	os.Exit(0)
}

func testInsertMany()  {
	//var rows = []City{
	//	City{LimitNum: 998,Name: "1西雅"},
	//	City{LimitNum: 998,Name: "2西雅"},
	//	City{LimitNum: 998,Name: "3西雅"},
	//	City{LimitNum: 998,Name: "4西雅"},
	//}
	//
	//ok,err := getEngine().NewSession().Insert(rows)
	//log.Printf("exists:%v, err:%v",ok,err)
	//log.Println(rows)
	//os.Exit(0)
}

func testQueryIn()  {
	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine := getEngine()
	engine.SetLogger(xorm.NewSimpleLogger(f))
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)

	session := engine.NewSession()
	session.Begin()
	defer session.Rollback()

	session.Table("oneField").Insert(map[string]string{"name":"aa"})
	var data = map[string]string{}
	_,err = session.Table("oneField").OrderBy("id desc").Limit(1).Get(&data)
	fmt.Println(data,err)

	//session.Close()
	session.Table("oneField").Insert(map[string]string{"name":"ab"})
	session.Commit()

	_,err = session.Table("oneField").OrderBy("id desc").Limit(1).Get(&data)
	fmt.Println(data,err)


	os.Exit(0)
}

type City struct {
	Id string `json:"id" xorm:"autoincr"`
	Name Name `json:"name" xorm:"name"`
	LimitNum int `json:"limit_num"`
	CreatedAt time.Time `json:"created_at" xorm:"created_at"`
}

type Name struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func (d Name) String() ([]byte, error) {
	return json.Marshal(d)
}

func (d Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(d)
}

func (d *Name) UnmarshalJSON(data []byte) error {
	var jd Name
	if err := json.Unmarshal(data, &jd); err != nil {
		return err
	}
	*d = jd
	return nil
}

func (city City)TableName()string  {
	return "city"
}


var engine = getEngine()
func testSession()  {
	var sessions []*xorm.Session
	for i:=0;i<100;i++{
		session := engine.NewSession()
		sessions = append(sessions, session)
	}

	time.Sleep(time.Minute)

}

func selectById()  {
	session := engine.NewSession().Table("city")
	var citys []City

	session = session.Where("id = 353")
	err := session.Find(&citys)

	fmt.Println(citys,err)
}
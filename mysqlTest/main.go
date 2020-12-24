package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/names"
	"log"
	"reflect"
)

type User struct {
	User_id    int64  `xorm:"pk autoincr 'userId'"` //指定主键并自增
	Name       string `xorm:"unique"`               //唯一的
	Balance    float64
	Time       int64 `xorm:"updated"` //修改后自动更新时间
	Creat_time int64 `xorm:"created"` //创建时间
	//Version int `xorm:"version"` //乐观锁
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Println(err)
	}
	tbMapper := names.NewPrefixMapper(names.SnakeMapper{},
		"pro2_")
	engine.SetTableMapper(tbMapper)
	err = engine.Sync(new(User))
	if err != nil {
		log.Println(err)
	}
	log.Println(reflect.TypeOf(engine.GetTableMapper()))
	log.Println(reflect.TypeOf(engine.GetColumnMapper()))

}

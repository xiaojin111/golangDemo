package main

import (
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v2"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a, err := mongodbadapter.NewAdapter("mongodb://admin:password@127.0.0.1:27017/casbin_rule")
	if err != nil {
		log.Println(err)
		return
	}
	e, err := casbin.NewEnforcer("./casbinTest/conf/rbac.conf", a)
	if err != nil {
		log.Println(err)
		return
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(e.Enforce("superAdmin", "gy", "project", "read"))
	log.Println(e.Enforce("quyuan", "jn", "asse", "read"))

}

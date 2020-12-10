package main

import (
	"github.com/jinzhu/now"
	"log"
)

func main() {
	t1 := now.BeginningOfMonth().Format("2006-01-02 15:04:05")
	log.Println(t1)
	t2 := now.EndOfMonth().Format("2006-01-02 15:04:05")
	log.Println(t2)
	now.MustParse(now.EndOfMonth().Format("2006-01-02 15:04:05"))
}

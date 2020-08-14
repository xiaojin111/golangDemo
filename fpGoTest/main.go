package main

import (
	"github.com/TeaEntityLab/fpGo"
	"log"
)

func main() {
	var m fpGo.MaybeDef
	m = fpGo.Maybe.Just(1)
	log.Println(m.IsPresent())
	log.Println(m.IsNil())
	log.Println(m.IsValid())
}

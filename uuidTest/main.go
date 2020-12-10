package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
)

func main() {
	u1 := uuid.Must(uuid.NewV4(), nil)
	log.Println(fmt.Sprintf(`%s`, u1))
}

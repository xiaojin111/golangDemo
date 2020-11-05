package main

import (
	"golangDemo/ljTest/crypt"
	"log"
)

func main() {
	//公钥解密
	log.SetFlags(log.Lshortfile)
	data, err := crypt.RSA_Decrypt(`EIW01nY7IfEHjqMq1lafQ1jH97/vhPmDoWe6izQXQBtQ9Prf8LgYANRsX3Gt9BOO1G6EOPQKS3XFDF02ygDHnwBovdM4UY+SkYgtmQpZXf1IqbPnMRhPWEdPCDZOh9eYFfzb9WnLgHAo1cjxWqqVgiPYXn3D2EF7SpRnwD3k2fk=
`)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(data))
}

package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

func main() {
	str := "@" + "d2c85217-2fb1-4e0b-9043-42ecf437a158" + "MNWSC2019"
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串
	s := hex.EncodeToString(h.Sum(nil))
	log.Println(s)

}

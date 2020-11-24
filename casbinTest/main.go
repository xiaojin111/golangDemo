package main

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
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

	e.AddPolicy("admin", "/hh", "read", "hello")
	e.RemovePolicy("admin", "/hh", "read", "hello")
	e.SavePolicy()
	log.Println(e.GetPolicy())
	//for i := 0; i < 10; i++ {
	//	err = e.LoadPolicy()
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	time.Sleep(3 * time.Second)
	//	log.Println(e.Enforce("f", "gy", "project", "read"))
	//}
	//log.Println(GetMsgById("88595350","421083199405183214"))

}

type ReqData1 struct {
	HealthReportId string `json:"healthReportId"`
	IdCardNumber   string `json:"idCardNumber"`
	TimeStamp      string `json:"timeStamp"`
	Sign           string `json:"sign"`
}

func GetMsgById(id, identity string) (msg string, err error) {
	sign := ""
	timestamp := fmt.Sprintf("%v", time.Now().Unix())
	reqData1 := ReqData1{
		HealthReportId: id,
		IdCardNumber:   identity,
		TimeStamp:      timestamp,
	}
	strA := "healthReportId=" + id + "&" +
		"idCardNumber=" + identity + "&" +
		"timeStamp=" + timestamp + "&" +
		"key=" + "2d40f3ad448f451d9c639f6cdf1af0fe"
	log.Println(strA)
	h := md5.New()
	h.Write([]byte(strA))
	sign = fmt.Sprintf("%X", h.Sum(nil))
	reqData1.Sign = sign
	reqStr, _ := json.Marshal(reqData1)
	log.Println(string(reqStr))
	payload := strings.NewReader(string(reqStr))
	req, _ := http.NewRequest("POST", "http://openapi2.ihaozhuo.com/getReport", payload)
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("timestamp", timestamp)
	//req.Header.Add("sign", sign)
	res, _ := http.DefaultClient.Do(req)
	var data []byte
	data, err = ioutil.ReadAll(res.Body)
	//log.Println(string(data))
	byteKey := []byte("gHA53u0Y#f%0UCc8")
	d1 := gjson.GetBytes(data, "data")
	var bytePass []byte
	bytePass, err = base64.StdEncoding.DecodeString(d1.String())
	origData := AesDecryptECB(bytePass, byteKey)
	msg = string(origData)
	return
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

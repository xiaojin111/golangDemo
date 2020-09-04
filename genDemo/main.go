package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"gopkg.in/h2non/gentleman.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	appid := "20200825000551153"
	key := "Tf3xViKBalZK0GxnhCw9"
	salt := time.Now().Unix() * 1000
	salt = 1598343320442
	query := "apple"
	from := "en"
	to := "zh"
	sign := fmt.Sprintf("%s%s%d%s", appid, query, salt, key)
	log.Println(sign)
	md5Str := fmt.Sprintf(`%x`, md5.Sum([]byte(sign)))
	log.Println(md5Str)
	cli := gentleman.New()
	cli.URL("http://api.fanyi.baidu.com/api/trans/vip/translate")
	req := cli.Request()
	req.Method("GET")
	req.Params(map[string]string{
		"q":     query,
		"appid": appid,
		"salt":  fmt.Sprintf(`%d`, salt),
		"from":  from,
		"to":    to,
		"sign":  md5Str,
	})
	req.AddHeader("Content-Type", "application/json")
	res, err := req.Send()
	if err != nil {
		log.Println(err)
	}
	log.Println(res.String())

	req1, _ := http.NewRequest("GET",
		"http://api.fanyi.baidu.com/api/trans/vip/translate?"+
			"q=apple&appid=20200825000551153&salt=1598343320442&from=en&to=zh&sign=f6ecd958e4f7a5e0d7a09b8a0cd6c205", nil)

	res1, _ := http.DefaultClient.Do(req1)
	defer res1.Body.Close()
	body, _ := ioutil.ReadAll(res1.Body)

	log.Printf(`%s`, body)
	A := []rune(string(body))
	str := ""
	for i := 0; i < len(A); i++ {
		cc := fmt.Sprintf(`%c`, A[i])
		if cc != `\` {
			str += cc
		} else {
			ss := ""
			for j := 0; j < 6; j++ {
				ss += fmt.Sprintf(`%c`, A[i+j])
			}
			log.Println(ss)
			kk, _ := u2s(ss)
			str += kk
			i += 5
		}
	}
	log.Println(str)
}
func u2s(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return
}

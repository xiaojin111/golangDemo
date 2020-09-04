package main

import (
	"log"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://qrhealth.ihaozhuo.com/check?identity=41594894534856")

	log.Println(u.Query()["identity"])
}

package main

import "log"

func main() {

	log.Println(f())

}
func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

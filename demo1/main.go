package main

import "log"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	a = append(a[:6], a[6+1:]...)
	log.Println(a)
}

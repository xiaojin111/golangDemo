package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var sum int32

func main() {
	go func() {
		log.Println(111)
		for true {
			time.Sleep(100 * time.Millisecond)
			log.Println("--------", runtime.NumGoroutine())
		}
	}()
	time.Sleep(10 * time.Second)
	log.Println(runtime.NumGoroutine())
}

func A() {
	go func() {
		log.Println(111)
		for true {
			time.Sleep(100 * time.Millisecond)
			log.Println("--------", runtime.NumGoroutine())
		}
	}()
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(20, func(i interface{}) {
		time.Sleep(1000 * time.Millisecond)
		n := i.(int32)
		atomic.AddInt32(&sum, n)
		fmt.Printf("run with %d\n", n)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}

type Request struct {
	Param  []byte
	Result chan []byte
}

func B() {
	pool, _ := ants.NewPoolWithFunc(100000, func(payload interface{}) {
		request, ok := payload.(*Request)
		if !ok {
			return
		}
		reverseParam := func(s []byte) []byte {
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			return s
		}(request.Param)

		request.Result <- reverseParam
	})
	defer pool.Release()

	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		param, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "request error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		request := &Request{Param: param, Result: make(chan []byte)}

		// Throttle the requests traffic with ants pool. This process is asynchronous and
		// you can receive a result from the channel defined outside.
		if err := pool.Invoke(request); err != nil {
			http.Error(w, "throttle limit error", http.StatusInternalServerError)
		}

		w.Write(<-request.Result)
	})

	http.ListenAndServe(":8080", nil)
}

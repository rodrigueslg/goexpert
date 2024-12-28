package main

import (
	"fmt"
	"net/http"
	"sync"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()

		w.Write([]byte(fmt.Sprintf("Number: %d\n", number)))
	})

	http.ListenAndServe(":8080", nil)
}

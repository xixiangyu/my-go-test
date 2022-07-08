package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func test(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	fmt.Println(time.Now())
	uri := os.Getenv("URI")
	acc := os.Getenv("ACCOUNT")
	pass := os.Getenv("PASSWORD")
	result := strings.Join([]string{host, uri, acc, pass}, "\n")
	w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/test", test)
	http.ListenAndServe("0.0.0.0:7788", nil)
}

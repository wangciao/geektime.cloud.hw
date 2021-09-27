package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start server...")

	// 当访问 localhost/healthz 时，应返回ok
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		//fmt.Println("error...")
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	// 只回應 ok
	io.WriteString(w, "ok")
}

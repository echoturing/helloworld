package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("LEANCLOUD_APP_PORT")
	if port == "" {
		port = "80"
	}
	log.Print("listen on port:", port)
	http.DefaultServeMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(fmt.Sprintf("hello world:%s", time.Now().Format(time.RFC3339))))
	})
	http.ListenAndServe("0.0.0.0:"+port, http.DefaultServeMux)
}

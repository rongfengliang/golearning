package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rongfengliang/golearning/netdemo"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", netdemo.GetContentFromNet("http://www.cnblogs.com/rongfengliang"))
}
func HttpServer() {
	http.HandleFunc("/", index_handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

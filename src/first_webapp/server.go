package main

import (
	"fmt"
	"net/http"
)

/*
引数について
ResponseWriterはインターフェース　参考：https://golang.org/pkg/net/http/?#ResponseWriter
Requestは構造体のポインタ        参考：https://golang.org/pkg/net/http/?#Request
*/
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

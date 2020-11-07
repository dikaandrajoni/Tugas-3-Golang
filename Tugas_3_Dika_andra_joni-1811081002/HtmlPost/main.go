package main

import (
	"fmt"
	"net/http"

	fn "Tugas_3_Dika_andra_joni-1811081002/HtmlPost/function"
)

func main() {
	http.HandleFunc("/", fn.RouteIndexGet)
	http.HandleFunc("/process", fn.RouteSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

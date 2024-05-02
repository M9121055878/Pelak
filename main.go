package main

import (
	"fmt"
	"net/http"
)

func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Hello World! ", r.URL.Path, "</h1>")
}

func main() {
	http.HandleFunc("/", myHandlerFunc)
	fmt.Println("Server is Runing in : 3000 ...")
	http.ListenAndServe(":3000", nil)
}

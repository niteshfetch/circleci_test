package main

import (
	"fmt"
	"net/http"
	"os"
)

func mainHandler() http.HandlerFunc {
	return http.HandlerFunc(func (w http.REsponseWriter, r *http.REquest)  {
		fmt.Fprintf(w, "Hello World! ")
	})
	
	func main()  {
		http.HandlerFunc("/", mainHandler())
		http.ListenAndServe(":8080", nil)
	}
}
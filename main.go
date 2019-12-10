package main

import (
	"fmt"
	"go-coveralls/book"
	"log"
	"net/http"
)

// Route はこのAPIサーバのルーティング設定をしている
func Route() *http.ServeMux {
	m := http.NewServeMux()

	m.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Fprintf(w, "Hello, %s", r.FormValue("name"))
	})

	return m
}

func main() {
	m := Route()
	fmt.Printf(book.OneFeed())
	log.Fatal(http.ListenAndServe(":8080", m))
}

// Calc 足し算するやつ
func Calc(x, y int) int {
	return x + y
}

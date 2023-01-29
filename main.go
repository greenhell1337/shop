package main

import (
	"githab/greenhell1337/shop/cmd/web"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/TeaShop", web.Shop)
	http.ListenAndServe(":8080", mux)
}
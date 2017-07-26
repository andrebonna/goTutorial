package main

import (
	"net/http"
)

func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Users!!!"))
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Products!!!"))
}

func main() {

	handler := RegexpHandler{}

	apiHandler := RegexpHandler{}

	apiHandler.HandleFunc("/users", handleUsers)
	apiHandler.HandleFunc("/products", handleProducts)

	handler.Handle("/api.*", &apiHandler)
	handler.Handle(".*", http.FileServer(http.Dir("static")))

	http.Handle("/", &handler)
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

var consumo string

func main() {
	http.HandleFunc("/post_data", posting_data)
	http.HandleFunc("/get_data", getting_data)
	log.Fatal(http.ListenAndServe(":7080", nil))
}

func posting_data(w http.ResponseWriter, r *http.Request) {

	consumo = r.FormValue("consumo")
	fmt.Println(consumo)

}

func getting_data(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(consumo))

}

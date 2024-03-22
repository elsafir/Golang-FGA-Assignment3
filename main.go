package main

import (
	"fmt"
	"net/http"

	"Golang-FGA-Assignment3/controller"
	"Golang-FGA-Assignment3/helper"
	"Golang-FGA-Assignment3/model"
)

func main() {
	// Inisialisasi data cuaca awal
	initData := model.Weather{
		Status: model.Status{
			Water: 0,
			Wind:  0,
		},
	}
	helper.WriteFile(initData, "weather.json")

	// Mulai ticker untuk pembaruan cuaca
	helper.StartTicker()

	// Menangani permintaan HTTP
	http.HandleFunc("/", controller.StatusController)

	// Menyajikan file statis
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("public/assets"))))

	// Mulai server HTTP
	fmt.Println("Server dimulai di localhost:9000")
	http.ListenAndServe(":1000", nil)
}

package helper

import (
	"math/rand"
	"time"

	"Golang-FGA-Assignment3/model"
)

func StartTicker() {
	// Membuat ticker dengan interval 15 detik
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	// Menginisialisasi data cuaca awal
	weather := model.Weather{
		Status: model.Status{
			Water: 0,
			Wind:  0,
		},
	}

	// Fungsi anonim untuk memperbarui cuaca setiap kali ticker berdenting
	go func() {
		for {
			select {
			case <-ticker.C:
				// Memperbarui data cuaca secara acak
				weather.Status.Wind = rand.Intn(100) + 1
				weather.Status.Water = rand.Intn(100) + 1

				// Menulis data cuaca ke file JSON
				WriteFile(weather, "weather.json")
			}
		}
	}()
}

package helper

import (
	"math/rand"
	"time"

	"Golang-FGA-Assignment3/model"
)

func StartTicker() {
	Weather := model.Weather{}

	ticker := time.NewTicker(15 * time.Second) // Membuat ticker yang akan memicu setiap 15 detik
	defer ticker.Stop()                        // Memastikan ticker berhenti saat fungsi selesai

	for {
		select {
		case <-ticker.C:
			Weather.Status.Wind = rand.Intn(100) + 1
			Weather.Status.Water = rand.Intn(100) + 1

			WriteFile(Weather, "weather.json")
		}
	}
}

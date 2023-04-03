package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/robfig/cron"
)

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	// Start microservice to update weather.json file every 15 seconds
	c := cron.New()
	c.AddFunc("*/15 * * * * *", func() {
		// generate random numbers between 1-100 for water and wind
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// create Weather object with random values
		w := Weather{
			Water: water,
			Wind:  wind,
		}

		// encode Weather object to JSON
		j, err := json.Marshal(w)
		if err != nil {
			log.Fatal(err)
		}

		// write JSON to file
		err = ioutil.WriteFile("weather.json", j, 0644)
		if err != nil {
			log.Fatal(err)
		}
	})
	c.Start()

	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", "user=postgres password=root dbname=weathernew sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create weather table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS weather (id SERIAL PRIMARY KEY, water INTEGER, wind INTEGER, created_at TIMESTAMP DEFAULT NOW())")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/weather.json", func(w http.ResponseWriter, r *http.Request) {
		// Baca data dari file weather.json
		data, err := ioutil.ReadFile("weather.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading weather data"))
			return
		}

		// Unmarshal data dari file weather.json ke struct Weather
		var weather Weather
		err = json.Unmarshal(data, &weather)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error parsing weather data"))
			return
		}

		// Marshal struct Weather ke JSON
		jsonData, err := json.Marshal(weather)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error encoding weather data"))
			return
		}

		// Set header untuk response
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// Kirim response
		w.Write(jsonData)
	})

	// Start HTTP server to display weather status
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// read weather.json file and decode to Weather object
		data, err := ioutil.ReadFile("weather.json")
		if err != nil {
			log.Fatal(err)
		}
		var weather Weather
		err = json.Unmarshal(data, &weather)
		if err != nil {
			log.Fatal(err)
		}

		// insert latest weather data to database
		_, err = db.Exec("INSERT INTO weather (water, wind) VALUES ($1, $2)", weather.Water, weather.Wind)
		if err != nil {
			log.Fatal(err)
		}

		// determine weather status based on water and wind values
		waterStatus := "aman"
		if weather.Water >= 6 && weather.Water <= 8 {
			waterStatus = "siaga"
		} else if weather.Water > 8 {
			waterStatus = "bahaya"
		}

		windStatus := "aman"
		if weather.Wind >= 7 && weather.Wind <= 15 {
			windStatus = "siaga"
		} else if weather.Wind > 15 {
			windStatus = "bahaya"
		}

		// display weather status
		fmt.Fprintf(w, "Tinggi Air mencapai (%d meter) Status: %s \n", weather.Water, waterStatus)
		fmt.Fprintf(w, "Kecepatan Angin mencapai (%d/detik) Status: %s", weather.Wind, windStatus)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

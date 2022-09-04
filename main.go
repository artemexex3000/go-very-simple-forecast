package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"main/database/postgres"
	"main/src"
	"net/http"
)

type forecast struct {
	Name     *string
	Forecast *string
}

func main() {
	var currentCountry string = "Sumy"

	db, err := postgres.DBConnect(postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "artemexex3000",
		Password: "159753159753",
		DBName:   "postgres",
		SSL:      "disable",
	})
	if err != nil {
		log.Fatalf("Failed to initialize DB: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping DB: %s", err.Error())
	}

	fmt.Println("Loading the Kastil' to show info in the browser")

	fmt.Println("DONE: localhost:8080/")

	/*-------------------Handlers-------------------*/

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles(`index.html`)

		currentCountry = r.FormValue("city")

		nameVar, forecastVar := src.Forecast(currentCountry)

		p := forecast{
			Name:     &nameVar,
			Forecast: &forecastVar,
		}

		db.MustExec("INSERT INTO requests (request) VALUES ($1)", &nameVar)

		err := t.Execute(w, p)
		if err != nil {

		}
	}

	/*-------------------Handler usage-------------------*/

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

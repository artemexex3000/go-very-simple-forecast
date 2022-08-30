package main

import (
	"fmt"
	"html/template"
	"log"
	src "main/src"
	"net/http"
	"os"
)

type forecast struct {
	Name     *string
	Forecast *string
}

func main() {
	var currentCountry string = "Sumy"

	/*fmt.Print("Enter the city/country: ")

	_, err := fmt.Scan(&currentCountry)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}*/

	fmt.Println("Loading the Kastil' to show info in the browser")

	fmt.Println("DONE: localhost:8080/")

	/*-------------------Handlers-------------------*/

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles(`index.html`)

		/*if r.Method != "Post" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}*/

		currentCountry = r.FormValue("city")

		nameVar, forecastVar := src.Forecast(currentCountry)

		p := forecast{
			Name:     &nameVar,
			Forecast: &forecastVar,
		}

		err := t.Execute(w, p)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	}

	/*-------------------Handler usage-------------------*/

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

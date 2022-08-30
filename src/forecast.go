package src

import (
	"io"
	"io/ioutil"
	"net/http"
)

func Forecast(currentCountry string) (string, string) {
	var mainUrl = "http://api.weatherapi.com/v1/current.json?key=85ed58b85f24483490c145353222608&q=" + currentCountry

	resp, err := http.Get(mainUrl)
	if err != nil {

	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	name, forecast := ExtractValue(responseData)

	return string(name), string(forecast) + " С°"
}

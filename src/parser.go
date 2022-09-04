package src

import (
	"github.com/buger/jsonparser"
)

func ExtractValue(body []byte) ([]byte, []byte) {
	getLocation, _, _, err := jsonparser.Get(body, "location", "name")
	if err != nil {

	}

	getTemperature, _, _, err := jsonparser.Get(body, "current", "temp_c")
	if err != nil {

	}

	return getLocation, getTemperature
}

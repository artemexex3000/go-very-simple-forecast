package src

import (
	"fmt"
	"github.com/buger/jsonparser"
	"os"
)

func ExtractValue(body []byte) ([]byte, []byte) {
	getLocation, _, _, err := jsonparser.Get(body, "location", "name")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	getTemperature, _, _, err := jsonparser.Get(body, "current", "temp_c")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return getLocation, getTemperature
}

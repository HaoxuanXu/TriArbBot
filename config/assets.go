package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Assets struct {
	Coins map[string][]string `json:"coins"`
}

func GetAssets() Assets {

	// read the json file
	content, err := ioutil.ReadFile("./config/assets/assets.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// unmarshal the data into struct
	var payload Assets
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("error during unmarshal: ", err)
	}

	return payload
}

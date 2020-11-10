package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type ConfigModel struct {
	Database struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Port     string `json:"port"`
		Host     string `json:"host"`
	} `json:"database"`
}

func GetConfigFile() (dataConfig ConfigModel, err error) {
	file, err := os.Open("config.json")
	defer file.Close()

	if err != nil {
		return
	}

	jsonParse := json.NewDecoder(file)
	jsonParse.Decode(&dataConfig)

	return
}

func Listen(addr string, handler http.Handler) {
	var err error
	err = http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatal(err)
	}
}

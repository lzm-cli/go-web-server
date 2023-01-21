package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const BuildVersion = "0.0.1"

type config struct {
	Port     int `json:"port"`
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Name     string `json:"name"`
	} `json:"database"`
	Mixin struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		SessionID    string `json:"session_id"`
		PinToken     string `json:"pin_token"`
		PrivateKey   string `json:"private_key"`
		PIN          string `json:"pin"`
	} `json:"mixin"`
	Key   string          `json:"key"`
	Admin map[string]bool `json:"admin"`
}

var C config

func init() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("config.json open fail...", err)
		return
	}
	err = json.Unmarshal(data, &C)
	if err != nil {
		log.Println("config.json parse err...", err)
	}
	log.Println("config.json load success...")
}

// SPDX-License-Identifier: Apache-2.0

package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

func NewConfig() *Config {
	// Load default values
	c := &Config{
		Address:  "localhost:6379",
		Password: "", // no password set
		Database: 0,  // use default DB
	}

	return c
}

func (c *Config) ParseConfig() {

	file, err := os.Open("conf.json")
	if err != nil {
		log.Println("Could not open config file. Using default config.")
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println("Error loading config: " + err.Error())
		return
	}
	*c = configuration
}

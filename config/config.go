// SPDX-License-Identifier: Apache-2.0

package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
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

	file, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Println(err)
		log.Println("Could not open config file. Using default config.")
		return
	}

	configuration := Config{}
	err = yaml.Unmarshal(file, &configuration)
	if err != nil {
		log.Println("Error loading config: " + err.Error())
		return
	}
	*c = configuration
}

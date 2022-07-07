// SPDX-License-Identifier: Apache-2.0

package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type AmqpConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	Database DbConfig   `yaml:"database"`
	Amqp     AmqpConfig `yaml:"amqpBroker"`
}

func NewConfig() *Config {
	// Load default values
	c := &Config{}

	c.Database.Address = "localhost:6379"
	c.Database.Password = "" // no password set
	c.Database.Database = 0  // use default DB

	c.Amqp.Host = "localhost"
	c.Amqp.Port = 5672
	c.Amqp.User = ""     // no password set
	c.Amqp.Password = "" // no password set

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

// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sogno-platform/state-estimation-api/config"
	"github.com/sogno-platform/state-estimation-api/routes"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	routes.RegisterEndpoints(api)
	return r
}

func main() {
	c := config.NewConfig()
	c.ParseConfig()

	// Print configuration values
	configJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("Configuration: %s\n", string(configJSON))

	r := setupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

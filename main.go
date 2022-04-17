// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sogno-platform/state-estimation-api/config"
	"github.com/sogno-platform/state-estimation-api/routes"
)

func setupRouter() *gin.Engine {
	config.Init()

	r := gin.Default()

	api := r.Group("/api")
	routes.RegisterEndpoints(api)
	return r
}

func main() {
	r := setupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// SPDX-License-Identifier: Apache-2.0

package estimation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterEstimationEndpoints(r *gin.RouterGroup) {

	r.GET("", GetEstimations)
	r.POST("", AddEstimation)
	r.GET("/:id", GetEstimationByID)
}

type job struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var jobs = []job{
	{ID: "1", Name: "test 1"},
	{ID: "2", Name: "test 2"},
}

func GetEstimations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jobs)
}

func AddEstimation(c *gin.Context) {
	var newJob job

	if err := c.BindJSON(&newJob); err != nil {
		return
	}

	jobs = append(jobs, newJob)
	c.IndentedJSON(http.StatusCreated, newJob)
}

func GetEstimationByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range jobs {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "job not found"})
}

// SPDX-License-Identifier: Apache-2.0

package estimation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterEstimationEndpoints(r *gin.RouterGroup) {

	r.GET("", GetEstimations)
	//r.POST("", AddEstimations)
	//r.GET("/:estimationID", GetEstimation)
}

type seJob struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var jobs = []seJob{
	{ID: "1", Name: "test 1"},
	{ID: "2", Name: "test 2"},
}

func GetEstimations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jobs)
}

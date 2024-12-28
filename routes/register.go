// SPDX-License-Identifier: Apache-2.0

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sogno-platform/state-estimation-api/estimation"
)

func RegisterEndpoints(r *gin.RouterGroup) {

	estimation.RegisterEstimationEndpoints(r.Group("/estimations"))
}

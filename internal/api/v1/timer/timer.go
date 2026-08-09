package timer

import "github.com/gin-gonic/gin"

func StartUserTimer(c *gin.Context) {

}

func StopUserTimer(c *gin.Context) {

}

func Routes(r *gin.RouterGroup) {
	actions := r.Group("/timer")

	actions.POST("/start", StartUserTimer)
	actions.POST("/end", StopUserTimer) //replace with query later
}

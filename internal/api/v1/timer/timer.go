package timer

import (
	"net/http"

	"github.com/ComputerScienceHouse/marks/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary      Starts a User's Timer
// @Description  Starts a User's Timer for the specified event. No-op if already started
// @Tags timer
// @Accept       json
// @Produce      json
// @Param        request  body      models.StartUserTimerInput  true  "Event Id"
// @Success      204 		"No Content"
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v2/timer/start [post]
func StartUserTimer(c *gin.Context) {
	c.JSON(http.StatusNoContent, models.StartUserTimerOutput{})
}

// @Summary      Resets a User's Timer
// @Description  Resets a User's Timer for the specified event. No-op if timer is not started
// @Tags timer
// @Accept       json
// @Produce      json
// @Param        request  body      models.ResetUserTimerInput  true  "Event Id"
// @Success      204 		"No Content"
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v2/timer/reset [post]
func ResetUserTimer(c *gin.Context) {
	c.JSON(http.StatusNoContent, models.ResetUserTimerOutput{})
}

// @Summary      Stops a User's Timer
// @Description  Stops a User's Timer for the specified event. No-op if timer is already stopped
// @Tags timer
// @Accept       json
// @Produce      json
// @Param        request  body      models.StopUserTimerInput  true  "Event Id"
// @Success      204 		"No Content"
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v2/timer/stop [post]
func StopUserTimer(c *gin.Context) {
	c.JSON(http.StatusNoContent, models.StopUserTimerOutput{})
}

func Routes(r *gin.RouterGroup) {
	timers := r.Group("/timer")

	timers.POST("/start", StartUserTimer)
	timers.POST("/reset", ResetUserTimer)
	timers.POST("/end", StopUserTimer)
}

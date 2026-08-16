package timer

import (
	"context"
	"net/http"
	"time"

	"github.com/ComputerScienceHouse/marks/internal/logging"
	"github.com/ComputerScienceHouse/marks/internal/models"
	"github.com/ComputerScienceHouse/marks/internal/redis"
	"github.com/ComputerScienceHouse/marks/internal/timer"
	"github.com/ComputerScienceHouse/marks/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary      Starts a User's Timer
// @Description  Starts a User's Timer for the specified event. No-op if timer is running
// @Tags timer
// @Accept       json
// @Produce      json
// @Param        request  body      models.StartUserTimerInput  true  "Event Id"
// @Success      204 		"No Content"
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v2/timer/start [post]
func StartUserTimer(c *gin.Context) {
	timer := timer.GetCurrentTime()

	user, err := users.GetCSHAuth(c)
	if err != nil {
		return
	}

	var input models.StartUserTimerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "v1/api/timer", "method": "StartUserTimer"}).Warning("failed to bind JSON")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Unable to process request!",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancel()

	if err := redis.StartUserTimer(ctx, timer, input.EventID, user.Uuid); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "v1/api/timer", "method": "StartUserTimer"}).Warning("failed to bind JSON")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Unable to process request!",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary      Resets a User's Timer
// @Description  Resets a User's Timer for the specified event. No-op if timer is not running
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
// @Description  Stops a User's Timer for the specified event. No-op if timer is not running
// @Tags timer
// @Accept       json
// @Produce      json
// @Param        request  body      models.StopUserTimerInput  true  "Event Id"
// @Success      204 		"No Content"
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v2/timer/stop [post]
func StopUserTimer(c *gin.Context) {
	timer := timer.GetCurrentTime()

	user, err := users.GetCSHAuth(c)
	if err != nil {
		return
	}

	var input models.StopUserTimerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "v1/api/timer", "method": "StopUserTimer"}).Warning("failed to bind JSON")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Unable to process request!",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancel()

	if err := redis.StopUserTimer(ctx, timer, input.EventID, user.Uuid); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "v1/api/timer", "method": "StopUserTimer"}).Warning("failed to bind JSON")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Unable to process request!",
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func Routes(r *gin.RouterGroup) {
	timers := r.Group("/timer")

	timers.POST("/start", StartUserTimer)
	timers.POST("/reset", ResetUserTimer)
	timers.POST("/end", StopUserTimer)
}

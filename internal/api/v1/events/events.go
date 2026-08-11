package events

import (
	"net/http"

	"github.com/ComputerScienceHouse/marks/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary      Creates a New Event Instance
// @Description  Creates a new Event instance for timer sessions
// @Tags events
// @Accept       json
// @Produce      json
// @Param        request  	body      models.CreateNewEventInput  true  "Name"
// @Success      204 		"No Content"
// @Failure      400      	{object}  models.ErrorResponse
// @Failure      401      	{object}  models.UnauthorizedResponse
// @Router       /api/v1/events/create [post]
func CreateNewEvent(c *gin.Context) {
	c.JSON(http.StatusNoContent, models.CreateNewEventOutput{})
}

// @Summary      Opens an Event
// @Description  Opens an Event to be eligible for timing.
// @Tags actions
// @Accept       json
// @Produce      json
// @Param        request  body      models.OpenEventInput  true  "Action Group"
// @Success      204      "No Content"
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v1/events/open [post]
func OpenEvent(c *gin.Context) {
	c.JSON(http.StatusNoContent, models.OpenEventOutput{})
}

// @Summary      Archives an Event
// @Description  Archives an Event to no longer be eligible for timing.
// @Tags actions
// @Accept       json
// @Produce      json
// @Param        request  body      models.ArchiveEventInput  true  "Action Group"
// @Success      204 	  "No Content"
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v1/events/archive [post]
func ArchiveEvent(c *gin.Context) {
	c.JSON(http.StatusNoContent, models.ArchiveEventOutput{})
}

// @Summary      Gets All Authorized Events
// @Description  Gets all open Events for standard users, or everything for Authorized Admins.
// @Tags actions
// @Accept       json
// @Produce      json
// @Param        request  body      models.ArchiveEventInput  true  "Action Group"
// @Success      200      {object}  models.GetEventsOutput
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v1/events/get [get]
func GetEvents(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetEventsOutput{})
}

func Routes(r *gin.RouterGroup) {
	events := r.Group("/events")

	events.POST("/create", CreateNewEvent)
	events.POST("/open", OpenEvent)
	events.POST("/archive", ArchiveEvent)
	events.GET("/get", GetEvents)
}

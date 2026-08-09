package v1

import (
	"github.com/ComputerScienceHouse/marks/internal/api/v1/events"
	"github.com/ComputerScienceHouse/marks/internal/api/v1/leaderboard"
	"github.com/ComputerScienceHouse/marks/internal/api/v1/timer"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")

	events.Routes(v1Group)
	leaderboard.Routes(v1Group)
	timer.Routes(v1Group)
}

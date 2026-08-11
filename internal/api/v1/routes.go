package v1

import (
	// "github.com/ComputerScienceHouse/marks/internal/api"
	"github.com/ComputerScienceHouse/marks/internal/api/v1/events"
	"github.com/ComputerScienceHouse/marks/internal/api/v1/leaderboard"
	"github.com/ComputerScienceHouse/marks/internal/api/v1/timer"

	// csh_auth "github.com/computersciencehouse/csh-auth/v2"
	"github.com/gin-gonic/gin"
)

// auth csh_auth.Auth for when auth is back up
func SetRoutes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")
	// v1Group.Use(api.CookieToAuthHeader())
	// v1Group.Use(auth.HeaderMiddleware())

	events.Routes(v1Group)
	leaderboard.Routes(v1Group)
	timer.Routes(v1Group)
}

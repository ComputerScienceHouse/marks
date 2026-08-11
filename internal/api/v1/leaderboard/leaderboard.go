package leaderboard

import (
	"net/http"

	"github.com/ComputerScienceHouse/marks/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary      Gets Top of Event Leaderboard
// @Description  Gets the Top 10 for an active event.
// @Tags leaderboard
// @Produce      json
// @Param        event    path      models.GetLeaderboardTopInput  true  "Event Id"
// @Success      200      {object}  models.GetLeaderboardTopOutput
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v1/leaderboard/top [get]
func GetLeaderboardTop(c *gin.Context) { //specifically needs an event, pulls from redis for speed
	_ = c.Query("event")

	c.JSON(http.StatusOK, models.GetLeaderboardTopOutput{})
}

// @Summary      Gets Data for a Leaderboard
// @Description  Gets Data from a leaderboard database, supports pagination with the cursor
// @Tags leaderboard
// @Produce      json
// @Param        request  path      models.GetLeaderboardGroupInput  false  "Event Id"
// @Success      200      {object}  models.GetLeaderboardGroupOutput
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v1/leaderboard/top [get]
func GetLeaderboardGroup(c *gin.Context) {
	_ = c.Query("event")

	c.JSON(http.StatusOK, models.GetLeaderboardGroupOutput{})
}

// @Summary      Removes an Entry
// @Description  Deletes a User timer for a specific event, updating leaderboards accordingly
// @Tags leaderboard
// @Accept       json
// @Produce      json
// @Param        request  body      models.RemoveLeaderboardEntryInput  true  "User / Event Data"
// @Success      200      {object}  models.RemoveLeaderboardEntryOutput
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.UnauthorizedResponse
// @Router       /api/v1/leaderboard/delete [delete]
func RemoveLeaderboardEntry(c *gin.Context) {
	c.JSON(http.StatusNoContent, models.RemoveLeaderboardEntryOutput{})
}

func Routes(r *gin.RouterGroup) {
	leaderboard := r.Group("/leaderboard")

	leaderboard.GET("/top", GetLeaderboardTop)
	leaderboard.GET("/group", GetLeaderboardGroup)
	leaderboard.DELETE("/delete", RemoveLeaderboardEntry) //decide if this should be post or not

}

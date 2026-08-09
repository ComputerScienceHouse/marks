package leaderboard

import "github.com/gin-gonic/gin"

func GetLeaderboardTop(c *gin.Context) { //specifically needs an event, pulls from redis for speed

}

func GetLeaderboardGroup(c *gin.Context) {

}

func Routes(r *gin.RouterGroup) {
	leaderboard := r.Group("/leaderboard")

	leaderboard.GET("/top", GetLeaderboardTop)
	leaderboard.GET("/group", GetLeaderboardGroup)
}

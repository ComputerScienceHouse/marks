package events

import "github.com/gin-gonic/gin"

func CreateNewEvent(c *gin.Context) {

}

func OpenEvent(c *gin.Context) {

}

func ArchiveEvent(c *gin.Context) {

}

func GetEvents(c *gin.Context) {

}

func Routes(r *gin.RouterGroup) {
	events := r.Group("/events")

	events.POST("/create", CreateNewEvent)
	events.POST("/open", OpenEvent)
	events.POST("/archive", ArchiveEvent)
	events.GET("/get", GetEvents)
}

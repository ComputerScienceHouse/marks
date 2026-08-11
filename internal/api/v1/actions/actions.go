package actions

import (
	"net/http"

	"github.com/ComputerScienceHouse/marks/internal/logging"
	"github.com/ComputerScienceHouse/marks/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const MODULE = "v1/api/actions" //stfu sonar

// GetActionGroup godoc
//
// @Summary      Gets a group of logged actions
// @Description  Gets a group of logged actions for auditing
// @Tags actions
// @Accept       json
// @Produce      json
// @Param        request  body      models.GetActionGroupInput  true  "Action Group"
// @Success      200      {object}  models.GetActionsGroupOutput
// @Failure      400      {object}  models.ErrorResponse
// @Failure      401      {object}  models.ErrorResponse
// @Router       /api/v1/actions/group [post]
func GetActionGroup(c *gin.Context) {

	var req models.GetActionGroupInput

	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": MODULE, "method": "GetActionGroup"}).Warning("failed to bind JSON")
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: "Unable to process request!",
		})
		return
	}

	resp := models.GetActionsGroupOutput{
		Actions: make([]models.GetActionGroupPart, 0),
	}

	c.JSON(http.StatusOK, resp)
}

func Routes(r *gin.RouterGroup) {
	actions := r.Group("/actions")

	actions.POST("/group", GetActionGroup) //replace with query later
}

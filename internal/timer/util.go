package timer

import (
	"time"

	"github.com/ComputerScienceHouse/marks/internal/models"
)

//current time as a centralized value, that way a start / stop cant be a different time annotation

func GetCurrentTime() models.TimerValue {
	return models.TimerValue(time.Now().UnixMilli())
}

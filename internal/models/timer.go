package models

type TimerStatus string
type TimerValue float64

const (
	TimerActive   TimerStatus = "active"
	TimerInactive TimerStatus = "inactive"
)

type TimerEntry struct {
	TimerStatus TimerStatus `json:"timer_status"`

	StartTime *TimerValue `json:"start_time"`
	EndTime   *TimerValue `json:"end_time"`
	Duration  *TimerValue `json:"duration"`
}

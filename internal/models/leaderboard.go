package models

type LeaderboardEntry struct {
	UserUUID string `json:"user_uuid"`
	Username string `json:"username"`

	EventID   string `json:"event_id"`
	EventName string `json:"event_name"`

	Duration TimerValue `json:"duration"`
}

package models

type GetLeaderboardTopInput string

type GetLeaderboardTopOutput struct {
	Entries   []GetLeaderboardEntry `json:"entries"`
	CanModify bool                  `json:"can_modify"`
}

type GetLeaderboardGroupInput string

type GetLeaderboardGroupOutput struct {
	Entries   []GetLeaderboardEntry `json:"entries"`
	CanModify bool                  `json:"can_modify"`
}

type RemoveLeaderboardEntryInput struct {
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
}

type RemoveLeaderboardEntryOutput struct{}

type GetLeaderboardEntry struct {
	Username  string `json:"username"`
	EventName string `json:"event_name"`
	Time      int64  `json:"time"`
}

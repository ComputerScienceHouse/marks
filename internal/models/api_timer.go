package models

type StartUserTimerInput struct {
	EventID string `json:"event_id"`
}
type StartUserTimerOutput struct{}

type ResetUserTimerInput struct {
	EventID string `json:"event_id"`
}
type ResetUserTimerOutput struct{}

type StopUserTimerInput struct {
	EventID string `json:"event_id"`
}
type StopUserTimerOutput struct{}

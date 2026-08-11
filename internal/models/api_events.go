package models

type CreateNewEventInput struct {
	Name string `json:"name"`
}
type CreateNewEventOutput struct{}

type OpenEventInput struct {
	ID string `json:"id"`
}
type OpenEventOutput struct{}

type ArchiveEventInput struct {
	ID string `json:"id"`
}
type ArchiveEventOutput struct{}

type GetEventsOutput []string

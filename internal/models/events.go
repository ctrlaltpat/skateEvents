package models

type Event struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"owner_id"`
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=3,max=500"`
	StartDate   string `json:"start_date" binding:"required,datetime=2006-01-02|datetime=2006-01-02T15:04:05Z07:00"`
	EndDate     string `json:"end_date" binding:"required,datetime=2006-01-02|datetime=2006-01-02T15:04:05Z07:00"`
	Location    string `json:"location" binding:"required,min=3,max=100"`
	Status      string `json:"status" binding:"required,oneof=draft live started ended completed cancelled archived"`
}

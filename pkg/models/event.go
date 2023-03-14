package models

type Event struct {
	Id           int64  `json:"id" gorm:"primaryKey"`
	DisciplineId int64  `json:"discipline_id"`
	Place        string `json:"place"`
	Date         string `json:"date"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Price        string `json:"price"`
	Limit        string `json:"limit"`
	Description  string `json:"description"`
}

package models

type Discipline struct {
	Id   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

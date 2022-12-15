package models

type User struct {
	ID        string `json:"id" gorm:"type:varchar(100);primaryKey"`
	FirstName string `json:"first_name" binding:"required" gorm:"type:varchar(255);not null"`
	LastName  string `json:"last_name" binding:"required" gorm:"type:varchar(255);not null"`
	Email     string `json:"email" binding:"email" gorm:"type:varchar(255);unique;not null"`
	Mobile    int64  `json:"mobile" binding:"mobile" gorm:"type:varchar(255);unique;not null"`
}

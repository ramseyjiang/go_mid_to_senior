package entity

type User struct {
	ID    string `json:"id" gorm:"type:varchar(100);primaryKey"`
	Name  string `json:"name" binding:"required" gorm:"type:varchar(255);not null"`
	Email string `json:"email" binding:"email" gorm:"type:varchar(255);not null"`
}

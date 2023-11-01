package model

type Users struct {
	ID       uint   `json:"title" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

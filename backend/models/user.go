package models

type UserDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// type UserDto struct {
// 	Username string `json:"username" binding:"required"`
// 	Token    string `json:"token" binding:"required"`
// }

type User struct {
	ID       uint   `gorm:"primaryKey;column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

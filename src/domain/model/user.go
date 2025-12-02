package model

type User struct {
	BaseModel
	Username string `gorm:"size:50;not null;unique"`
	Password string `gorm:"size:255;not null"`
	Email    string `gorm:"size:100;unique"`
	IsActive bool   `gorm:"default:true"`
}

func (User) TableName() string {
	return "users"
}

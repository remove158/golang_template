package entity

type User struct {
	ID       int16  `gorm:"primary_key:auto_increment"`
	Username string `gorm:"varchar(255)"`
	Password string `gorm:"varchar(255)"`
}

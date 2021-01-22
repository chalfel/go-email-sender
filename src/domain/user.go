package domain

// User struct
type User struct {
	Base
	Name  string `gorm: "type:varchar(255)"`
	Email string `gorm: "type:varchar(255);unique_index"`
}

package users

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"not null"`
}

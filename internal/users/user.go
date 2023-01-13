package users

// User example
type User struct {
	ID           uint64 `json:"id" gorm:"primarykey"`
	PasswordHash string `json:"-"`
	Email        string `json:"email"`
	Username     string `json:"username"`
}

func (User) TableName() string {
	return "user"
}

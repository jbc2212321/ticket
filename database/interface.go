package database

type UserDao interface {
	ExistUser(phone, category int64) bool
	AddUser(user *User) error
	CheckUser(phone, category int64, password string) bool
}

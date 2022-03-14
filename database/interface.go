package database

type UserDao interface {
	CheckUser(phone, category int64) bool
	AddUser(user *User) error
}

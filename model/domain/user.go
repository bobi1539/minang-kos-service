package domain

type User struct {
	Id          int64
	Email       string
	Password    string
	Name        string
	PhoneNumber string
	Role        Role
	BaseDomain
}

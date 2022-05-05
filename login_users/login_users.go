package loginusers

type User struct {
	Name     string
	Password string
}

type LoginUser interface {
	Login(User) string
}

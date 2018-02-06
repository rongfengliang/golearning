package interfaceapp

type UserLogin interface {
	Login() (token string, err error)
	Logout() string
}

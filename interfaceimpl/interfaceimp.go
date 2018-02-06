package interfaceimpl

import (
	"fmt"
)

type FirstAppUserLogin struct {
	UserName     string
	UserPassword string
	UserAppType  string
}

func (f FirstAppUserLogin) Login() (token string, err error) {
	if f.UserPassword == "dalong" && f.UserName == "dalong" {
		token = "aaaaaaaaa"
	} else {
		token = "create error"
		err = fmt.Errorf("%s", "create user token error")
	}
	return token, nil
}
func (f FirstAppUserLogin) Logout() string {
	fmt.Println(f.UserName, "is logout")
	return "ok"
}

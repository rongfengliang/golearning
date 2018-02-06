package panicapp

import "fmt"

type UserError struct{}
type LoginError struct{}
type OtherError struct{}

func PanicDemo(name string) string {
	if name == "" {
		panic("name must not empty")
	}
	return name
}

func RecoverDemo(name string) (title string, err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
		case UserError{}:
			err = fmt.Errorf("hapence user error")
		default:
			panic(p)
		}
	}()
	if name == "" {
		panic(LoginError{})
	} else if name == "dalong" {
		panic(UserError{})
	} else if name == "other" {
		panic(OtherError{})
	} else {
		title = name
	}
	return title, err
}

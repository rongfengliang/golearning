package structdemo

type User struct {
	X, Y        int
	UserVersion string
}
type Person struct {
	X, Y          int
	PersonVersion string
}
type Basic struct {
	X, Y int
}
type UserNew struct {
	BasicInfo   Basic
	UserVersion string
}
type PersonNew struct {
	BasicInfo     Basic
	PersonVersion string
}
type PersonApp struct {
	Basic
	PersonVersion string
}

func GetInfo() UserNew {
	var item UserNew
	item.BasicInfo.X = 44
	item.BasicInfo.Y = 34
	item.UserVersion = "v1"
	return item
}
func GetPersonInfo() PersonApp {

	var item PersonApp
	item.X = 44
	item.Y = 67
	item.PersonVersion = "v2"
	return item
}

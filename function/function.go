package function

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rongfengliang/golearning/interfaceapp"
)

type MyFunc struct {
}

func (fpp MyFunc) CreateUserToken(login interfaceapp.UserLogin) string {
	token, _ := login.Login()
	return token
}
func (fpp MyFunc) Logout(login interfaceapp.UserLogin) string {
	result := login.Logout()
	return result
}
func (foo MyFunc) ErrorLog(url string) ([]string, error) {
	if url == "" {
		return nil, fmt.Errorf("%s", "http url is nil")
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%s", "some error")
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("%s", "http request is error")
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return []string{string(b[:])}, nil
}

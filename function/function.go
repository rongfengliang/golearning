package function

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type MyFunc struct {
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

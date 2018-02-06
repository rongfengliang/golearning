package testapp

import (
	"testing"
)

func TestAppName(t *testing.T) {
	if Appdemo("dalong") {
		t.Log("is ok")
	} else {
		t.Error("is error")
	}
}

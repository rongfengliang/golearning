package tools

import (
	"fmt"

	"github.com/rongfengliang/golearning/packagedemo/log"
)

// print lvs init information
func LVSINIT() {
	fmt.Println("lvs init")
}

// print lvs detail information
func LVSDetailInfo() {
	log.Info("dalongdemo.....lvs")
}

// print lvs error infoarmation
func LVSAPPErrorInfo() {
	log.Error("dalongdemo....lvs debug ingo")
}

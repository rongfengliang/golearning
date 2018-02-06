package packagedemo

import (
	"fmt"

	"github.com/rongfengliang/golearning/packagedemo/tools"
)

func init() {
	fmt.Println("start server .......")
	tools.LVSINIT()
	tools.LVSDetailInfo()
	tools.LVSAPPErrorInfo()
}

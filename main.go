package main

import (
	"fmt"

	"github.com/rongfengliang/golearning/json"
	"github.com/rongfengliang/golearning/users"
)

const appversion = "1.0"

const (
	app  = "dalong"
	user = "dalongdemo"
	demo = "appdemo"
)

func appdemo() (name string, age int) {

	return "dalong", 444
}
func arraydemo() {
	var a = [3]string{"dalong", "appdemo", "ddd"}
	for i, v := range a {
		fmt.Println(i, v)
	}
	var inta [3]int = [3]int{1, 3, 4}
	for i, v := range inta {
		fmt.Println(i, v)
	}

	var appinfo = [...]int{4, 5, 5, 6, 7}
	for _, v := range appinfo {
		fmt.Println(v)
	}

}
func appenddemo() {

	var runes []rune

	for _, r := range "dalong rong" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}
func structdemo2() {
	user := users.Appdemo2{AppName: "dddd", AppIndex: 33}
	fmt.Printf(user.AppName)
}
func main() {
	// app := 3234
	// fmt.Print(app)
	// var user = users.User{Name: "dalofffffffffng", Age: 44}
	// fmt.Printf("dalong" + "世界")
	// fmt.Printf(user.Name)
	// resp, err := http.Get("http://www.baidu.com")
	// if err != nil {
	// 	fmt.Printf("there is some wrong for get data")
	// }
	// b, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	fmt.Printf("wrong convert data")
	// 	os.Exit(1)
	// }
	// fmt.Printf("%s", b)
	// fmt.Print(appversion)
	// server := server.Server{}
	// server.HttpServer()
	// a, b := appdemo()
	// fmt.Print(a, b)
	// p := new(server.Server)
	// p.HttpServer()
	// x := 123
	// y := fmt.Sprintf("%d", x)
	// fmt.Println(y, strconv.Itoa(x))
	// count := strings.Count("dadadadada", "da")
	// fmt.Print(count)
	// arraydemo()
	// appenddemo()
	// structdemo()
	/*
		 //  nested struct demo
			var info myapp.PersonApp = myapp.GetPersonInfo()
			fmt.Println(info.PersonVersion)

			appinfo := myapp.PersonApp{myapp.Basic{3, 6}, "appdemo"}

			fmt.Println(appinfo.PersonVersion)

			p := myapp.PersonApp{

				Basic: myapp.Basic{
					X: 44,
					Y: 666,
				},
				PersonVersion: "aaaaaa",
			}
			fmt.Println(p.PersonVersion)
	*/
	/*
		// json encode demo
	*/
	fmt.Print(movie.Json())

}

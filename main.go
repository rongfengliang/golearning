package main

import (
	"fmt"

	_ "github.com/rongfengliang/golearning/init"
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
	/*
		// http server demo
		server := server.Server{}
		server.HttpServer()
		a, b := appdemo()
		fmt.Print(a, b)
		p := new(server.Server)
		p.HttpServer()
	*/

	/*
		 // basic data type demo
		arraydemo()
		appenddemo()
		structdemo()
	*/

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
				fmt.Print(movie.Json())

				//decode json demo
		        info := movie.Json()
			fmt.Println(info)
			var items []movie.Movie = movie.Entity(info)
			for v, item := range items {
				fmt.Println(v, item.Color)
			}
	*/
	/*
			// template demo
		template.MyTemplate()

	*/

	/*
			 // func demo


			result, err := myfunc.ErrorLog("http://www.cnblogs.com/rongfengliang")
			if err != nil {
				fmt.Print("is wrong", err)
			}
			fmt.Print(result)

			myfunc := function.MyFunc{}
		result, err := myfunc.ErrorLog("")
		if err != nil {
			fmt.Print("is wrong", err)
		}
		fmt.Print(result)
	*/
	/*
			  // panic demo
			  panicapp.PanicDemo("")
			  // recover demo
			title, err := panicapp.RecoverDemo("dalongss")
		fmt.Println(title, err)
	*/

	/*
			 // method demo
			p := method.Point{X: 44, Y: 21}
		    result := p.Distance(method.Point{X: 3, Y: 6})
			result2 := p.Distance2(method.Point{X: 3, Y: 6})
			fmt.Println(result, "distance")
			fmt.Println(result2, "distance2")
	*/
	/*
		   //  method with point demo

		p := method.Point{X: 44, Y: 21}
		c := method.Cicle{Point: method.Point{X: 3, Y: 3}, Center: &method.Center{X: 5, Y: 7}}
		result := c.Length(p)
		result2 := c.Distance(p)
		fmt.Println(result)
		fmt.Println(result2)
		fmt.Println("call from point")
		c.ReturnX()
		c.ReturnX2()
	*/

}

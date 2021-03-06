package movie

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string
	Year  int  `json:"released"`
	Color bool `json:"color"`
}

func Entity(jsoninfo string) []Movie {
	// var items []struct {
	// 	Title string
	// 	Year  int  `json:"released"`
	// 	Color bool `json:"color"`
	// }
	var movies []Movie
	if err := json.Unmarshal([]byte(jsoninfo), &movies); err != nil {
		fmt.Println("some wrong")
	}
	fmt.Println(movies)
	return movies
}
func Json() string {
	movies := []Movie{
		{
			Title: "dalong",
			Year:  31,
			Color: false,
		},
		{
			Title: "dalongaaaaa",
			Year:  21,
			Color: false,
		},
		{
			Title: "dalowwwwng",
			Year:  45,
			Color: true,
		},
	}
	// b, err := json.Marshal(movies)
	b, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		fmt.Println("json encode is wrong")
	}
	result := string(b)
	return result
}

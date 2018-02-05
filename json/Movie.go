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

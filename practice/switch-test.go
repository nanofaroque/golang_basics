package main

import (
	"fmt"
	"net/http"
)

func main() {
	region, continent := location("Santa Monica")
	fmt.Printf("Matt lives in %s, %s", region, continent)

	client := &http.Client{}
	resp, err := client.Get("http://gobootcamp.com")
	fmt.Println(err)
	fmt.Println(resp)
}
func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Santa Monica", "LA", "Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent

}

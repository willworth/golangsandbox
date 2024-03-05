package main

import "fmt"

func main() {
	websites := map[string]string{ //key val types
		"Google":              "http://google.com",
		"Amazon Web Services": "http://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
	websites["LinkedIn"] = "http://linkedin.com"
	delete(websites, "Google")
	fmt.Println(websites)
}

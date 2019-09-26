package main

import "fmt"
import "net/url"

func main() {
	var urlString = "http://localhost:8080/movie?title=hello&code=29"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme) // http
	fmt.Printf("host: %s\n", u.Host)       // depeloper.com:80
	fmt.Printf("path: %s\n", u.Path)       // /hello
	var all = u.Query()
	var movie = u.Query()["title"][0] // john wick
	var code = u.Query()["code"][0]   // 27
	fmt.Printf("movie: %s, code: %s\n", movie, code)
	fmt.Printf("all: %s\n", all)
}

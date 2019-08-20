// file: http-nofollow-request.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// noFollowRedirect()
	indefiniteFollowRedirect()
}

func noFollowRedirect() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}

	resp, err := client.Get("http://www.jonathanmh.com")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("StatusCode:", resp.StatusCode)
	fmt.Println(resp.Request.URL)
}
func indefiniteFollowRedirect() {
	myURL := "http://www.jonathanmh.com"
	nextURL := myURL
	var i int
	for {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}

		resp, err := client.Get(nextURL)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("StatusCode:", resp.StatusCode)
		fmt.Println(resp.Request.URL)

		if resp.StatusCode == 200 {
			fmt.Println("Done!")
			break
		} else {
			nextURL = resp.Header.Get("Location")
			i++
		}
	}
}

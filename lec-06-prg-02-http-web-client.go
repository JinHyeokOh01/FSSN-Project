package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("## HTTP client started.")

	// GET request for directory
	fmt.Println("## GET request for http://localhost:8080/temp/")
	resp, err := http.Get("http://localhost:8080/temp/")
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	
	fmt.Println("## GET response [start]")
	fmt.Println(string(body))
	fmt.Println("## GET response [end]")

	// GET request for calculation
	fmt.Println("## GET request for http://localhost:8080/?var1=9&var2=9")
	resp, err = http.Get("http://localhost:8080/?var1=9&var2=9")
	if err != nil {
		fmt.Printf("Error making GET request: %v\n", err)
		return
	}
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	
	fmt.Println("## GET response [start]")
	fmt.Println(string(body))
	fmt.Println("## GET response [end]")

	// POST request
	fmt.Println("## POST request for http://localhost:8080/ with var1 is 9 and var2 is 9")
	data := url.Values{}
	data.Set("var1", "9")
	data.Set("var2", "9")
	
	resp, err = http.Post("http://localhost:8080",
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Printf("Error making POST request: %v\n", err)
		return
	}
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	
	fmt.Println("## POST response [start]")
	fmt.Println(string(body))
	fmt.Println("## POST response [end]")

	fmt.Println("## HTTP client completed.")
}
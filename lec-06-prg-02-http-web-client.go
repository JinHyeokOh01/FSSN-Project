package main

import(
	"fmt"
	"net/http"
	"io"
	"net/url"
)

func main() {
    fmt.Println("## HTTP client started.")

    // First GET request
    fmt.Println("## GET request for http://localhost:8080/temp/")
    resp, _ := http.Get("http://localhost:8080/temp/")
    fmt.Println("## GET response [start]")
    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
    fmt.Println("## GET response [end]")
    resp.Body.Close()

    // Second GET request with parameters
    fmt.Println("## GET request for http://localhost:8080/?var1=9&var2=9")
    resp, _ = http.Get("http://localhost:8080/?var1=9&var2=9")
    fmt.Println("## GET response [start]")
    body, _ = io.ReadAll(resp.Body)
    fmt.Println(string(body))
    fmt.Println("## GET response [end]")
    resp.Body.Close()

    // POST request
    fmt.Println("## POST request for http://localhost:8080/ with var1 is 9 and var2 is 9")
    postData := url.Values{
        "var1": {"9"},
        "var2": {"9"},
    }
    resp, _ = http.PostForm("http://localhost:8080/", postData)
    fmt.Println("## POST response [start]")
    body, _ = io.ReadAll(resp.Body)
    fmt.Println(string(body))
    fmt.Println("## POST response [end]")
    resp.Body.Close()

    fmt.Println("## HTTP client completed.")
}
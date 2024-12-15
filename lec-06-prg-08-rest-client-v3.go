package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func printResponse(num int, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	var key string
	for k := range result {
		key = k
		break
	}

	fmt.Printf("#%d Code: %d >> JSON: %s >> JSON Result: %s\n",
		num, resp.StatusCode, string(body), result[key])
}

func main() {
	baseURL := "http://127.0.0.1:5000/membership_api"
	client := &http.Client{}

	resp, _ := client.Get(baseURL + "/0001")
	printResponse(1, resp)

	data := url.Values{}
	data.Set("0001", "apple")
	resp, _ = client.Post(baseURL+"/0001", 
		"application/x-www-form-urlencoded", 
		strings.NewReader(data.Encode()))
	printResponse(2, resp)

	resp, _ = client.Get(baseURL + "/0001")
	printResponse(3, resp)

	data = url.Values{}
	data.Set("0001", "xpple")
	resp, _ = client.Post(baseURL+"/0001", 
		"application/x-www-form-urlencoded", 
		strings.NewReader(data.Encode()))
	printResponse(4, resp)

	data = url.Values{}
	data.Set("0002", "xrange")
	req, _ := http.NewRequest(http.MethodPut, baseURL+"/0002", 
		strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ = client.Do(req)
	printResponse(5, resp)

	data = url.Values{}
	data.Set("0002", "xrange")
	resp, _ = client.Post(baseURL+"/0002", 
		"application/x-www-form-urlencoded", 
		strings.NewReader(data.Encode()))

	data = url.Values{}
	data.Set("0002", "orange")
	req, _ = http.NewRequest(http.MethodPut, baseURL+"/0002", 
		strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ = client.Do(req)
	printResponse(6, resp)

	req, _ = http.NewRequest(http.MethodDelete, baseURL+"/0001", nil)
	resp, _ = client.Do(req)
	printResponse(7, resp)

	req, _ = http.NewRequest(http.MethodDelete, baseURL+"/0001", nil)
	resp, _ = client.Do(req)
	printResponse(8, resp)
}
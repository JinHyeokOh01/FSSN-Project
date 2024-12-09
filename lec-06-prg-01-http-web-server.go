package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MyHttpHandler struct{}

func (h *MyHttpHandler) printHttpRequestDetail(r *http.Request) {
	host := r.RemoteAddr
	hostParts := strings.Split(host, ":")
	
	// URL의 전체 경로를 가져옴 (쿼리 파라미터 포함)
	fullPath := r.URL.Path
	if r.URL.RawQuery != "" {
		fullPath = fullPath + "?" + r.URL.RawQuery
	}
	
	fmt.Printf("::Client address   : %s\n", hostParts[0])
	fmt.Printf("::Client port      : %s\n", hostParts[1])
	fmt.Printf("::Request command  : %s\n", r.Method)
	fmt.Printf("::Request line     : %s %s %s\n", r.Method, fullPath, r.Proto)
	fmt.Printf("::Request path     : %s\n", fullPath)
	fmt.Printf("::Request version  : %s\n", r.Proto)
}

func (h *MyHttpHandler) simpleCalc(para1, para2 int) int {
	return para1 * para2
}

func (h *MyHttpHandler) parameterRetrieval(msg string) []int {
	result := make([]int, 2)
	fields := strings.Split(msg, "&")
	
	for _, field := range fields {
		parts := strings.Split(field, "=")
		if len(parts) == 2 {
			val, _ := strconv.Atoi(parts[1])
			if strings.HasPrefix(parts[0], "var1") {
				result[0] = val
			} else if strings.HasPrefix(parts[0], "var2") {
				result[1] = val
			}
		}
	}
	return result
}

func (h *MyHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	
	switch r.Method {
	case "GET":
		fmt.Println("## do_GET() activated.")
		h.printHttpRequestDetail(r)

		if r.URL.RawQuery != "" {
			params := h.parameterRetrieval(r.URL.RawQuery)
			result := h.simpleCalc(params[0], params[1])
			
			fmt.Fprintf(w, "<html>GET request for calculation => %d x %d = %d</html>", 
				params[0], params[1], result)
			fmt.Printf("## GET request for calculation => %d x %d = %d.\n", 
				params[0], params[1], result)
		} else {
			fmt.Fprintf(w, "<html><p>HTTP Request GET for Path: %s</p></html>", r.URL.Path)
			fmt.Printf("## GET request for directory => %s.\n", r.URL.Path)
		}

	case "POST":
		fmt.Println("## do_POST() activated.")
		h.printHttpRequestDetail(r)

		r.ParseForm()
		var1, _ := strconv.Atoi(r.FormValue("var1"))
		var2, _ := strconv.Atoi(r.FormValue("var2"))
		result := h.simpleCalc(var1, var2)

		postData := fmt.Sprintf("var1=%s&var2=%s", r.FormValue("var1"), r.FormValue("var2"))
		fmt.Printf("## POST request data => %s.\n", postData)
		
		response := fmt.Sprintf("POST request for calculation => %d x %d = %d", var1, var2, result)
		fmt.Fprint(w, response)
		fmt.Printf("## POST request for calculation => %d x %d = %d.\n", var1, var2, result)
	}
}

func main() {
	serverName := "localhost"
	serverPort := 8080
	
	handler := &MyHttpHandler{}
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", serverName, serverPort),
		Handler: handler,
	}

	fmt.Printf("## HTTP server started at http://%s:%d.\n", serverName, serverPort)

	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}

	fmt.Println("HTTP server stopped.")
}
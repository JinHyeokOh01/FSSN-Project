package main

import (
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "strconv"
    "strings"
)

type MyHTTPHandler struct{}

func (h *MyHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        h.do_GET(w, r)
    case "POST":
        h.do_POST(w, r)
    }
}

func (h *MyHTTPHandler) printHTTPRequestDetail(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("::Client address  : %s\n", r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")])
    fmt.Printf("::Client port     : %s\n", r.RemoteAddr[strings.LastIndex(r.RemoteAddr, ":")+1:])
    fmt.Printf("::Request command : %s\n", r.Method)
    fmt.Printf("::Request line    : %s %s %s\n", r.Method, r.URL.Path, r.Proto)
    fmt.Printf("::Request path    : %s\n", r.URL.Path)
    fmt.Printf("::Request version : %s\n", r.Proto)
}

func (h *MyHTTPHandler) sendHTTPResponseHeader(w http.ResponseWriter) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "text/html")
}

func parameterRetrieval(msg string) []int {
    var result []int
    fields := strings.Split(msg, "&")
    for _, field := range fields {
        parts := strings.Split(field, "=")
        if len(parts) > 1 {
            if num, err := strconv.Atoi(parts[1]); err == nil {
                result = append(result, num)
            }
        }
    }
    return result
}

func simpleCalc(para1, para2 int) int {
    return para1 * para2
}

func (h *MyHTTPHandler) do_GET(w http.ResponseWriter, r *http.Request) {
    fmt.Println("## do_GET() activated.")
    h.printHTTPRequestDetail(w, r)
    h.sendHTTPResponseHeader(w)

    if strings.Contains(r.URL.Path, "?") {
        params := parameterRetrieval(r.URL.RawQuery)
        if len(params) >= 2 {
            result := simpleCalc(params[0], params[1])
            fmt.Fprintf(w, "<html>")
            fmt.Fprintf(w, "GET request for calculation => %d x %d = %d", params[0], params[1], result)
            fmt.Fprintf(w, "</html>")
            fmt.Printf("## GET request for calculation => %d x %d = %d\n", params[0], params[1], result)
        }
    } else {
        fmt.Fprintf(w, "<html><p>HTTP Request GET for Path: %s</p></html>", r.URL.Path)
        fmt.Printf("## GET request for directory => %s\n", r.URL.Path)
    }
}

func (h *MyHTTPHandler) do_POST(w http.ResponseWriter, r *http.Request) {
    fmt.Println("## do_POST() activated.")
    h.printHTTPRequestDetail(w, r)
    h.sendHTTPResponseHeader(w)

    if err := r.ParseForm(); err != nil {
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }

    var1 := r.FormValue("var1")
    var2 := r.FormValue("var2")
    num1, _ := strconv.Atoi(var1)
    num2, _ := strconv.Atoi(var2)
    result := simpleCalc(num1, num2)

    fmt.Fprintf(w, "POST request for calculation => %d x %d = %d", num1, num2, result)
    fmt.Printf("## POST request for calculation => %d x %d = %d\n", num1, num2, result)
}

func main() {
    serverName := "localhost"
    serverPort := 8080

    handler := &MyHTTPHandler{}
    
    fmt.Printf("## HTTP server started at http://%s:%d\n", serverName, serverPort)

    server := &http.Server{
        Addr:    fmt.Sprintf("%s:%d", serverName, serverPort),
        Handler: handler,
    }

    // Handle graceful shutdown
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)

    go func() {
        <-c
        fmt.Println("HTTP server stopped.")
        server.Close()
        os.Exit(0)
    }()

    if err := server.ListenAndServe(); err != nil {
        fmt.Printf("Server error: %v\n", err)
    }
}
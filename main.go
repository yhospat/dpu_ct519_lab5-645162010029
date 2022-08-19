package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}



func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.Handle("/about.html", http.FileServer(http.Dir("./www")))
    http.Handle("/myresearch.html", http.FileServer(http.Dir("./www")))
    http.Handle("/", http.FileServer(http.Dir("./www")))

    http.ListenAndServe(":8090", nil)
}

package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    fmt.Fprintf(w, "Hello Kubernetes again!\n\n"+
                   "running on Pod: %s", hostname)
}

func main() {
    address := ":8080"
    http.HandleFunc("/", handler)
    fmt.Printf("Go Web Server started on port %s \n"+
               "Check it out in your web browser on http://localhost%s \n"+
               "Press Ctrl+C to exit from server", address, address)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

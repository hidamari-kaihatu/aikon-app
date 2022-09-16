package main
import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello, World</h1>")
}
func main() {
    http.HandleFunc("/", helloHandler)
    // fmt.Println("Server Start")
    http.ListenAndServe(":8080", nil)
    // fmt.Println("Hello, World!!")
}
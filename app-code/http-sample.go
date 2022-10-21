package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Congratulations! Your Go application has been successfully deployed.")
}

func main() {
	fmt.Println("Starting basic go application on port 3000")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

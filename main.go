package main

import (
	"fmt"
	"net/http"
)

func ItsWorks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This apps works!!")

}

func ItsWorksOnV2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This apps works!!")

}

func main() {
	http.HandleFunc("/", ItsWorks)

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
)

func ItsWorksOnV1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This apps works!! - Version 1!")

}

func ItsWorksOnV2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This apps works on  - Version 2!")

}

func main() {
	http.HandleFunc("/v1", ItsWorksOnV1)
	http.HandleFunc("/v2", ItsWorksOnV2)

	http.ListenAndServe(":8080", nil)
}

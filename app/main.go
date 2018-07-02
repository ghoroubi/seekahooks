package main

import (
	"log"
	"net/http"
	"fmt"
	."webhook/lib"
)

func main() {
	http.HandleFunc("/payload", Chain(GetHooksfunc, Method("POST"), Logger()))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "welcome...")
	})
	log.Println(http.ListenAndServe(":8585", nil))
}

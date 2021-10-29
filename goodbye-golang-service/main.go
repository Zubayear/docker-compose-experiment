package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func goodbye(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("logs_container/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	resp := "Goodbye World!"
	log.Println("Method", r.Method, "Response", resp)
	fmt.Fprintf(w, "%s %s\n", resp, time.Now())
}

func main() {
	http.HandleFunc("/goodbye", goodbye)
	http.ListenAndServe(":8080", nil)
}

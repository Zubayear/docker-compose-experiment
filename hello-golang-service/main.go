package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// func createDir(name string) (string, error) {
// 	path, err := os.Getwd()
// 	if err != nil {
// 		log.Println("error msg", err)
// 	}

// 	outPath := filepath.Join(path, name)

// 	if _, err := os.Stat(outPath); os.IsNotExist(err) {
// 		return outPath, os.Mkdir(outPath, 0755)
// 	}
// 	return "", nil
// }

func greet(w http.ResponseWriter, r *http.Request) {
	// filename, err := createDir("logs_container")
	// if err != nil {
	// 	log.Println(err)
	// }

	file, err := os.OpenFile("logs_container/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	resp := "Hello World!"
	log.Println("Method", r.Method, "Response", resp)
	fmt.Fprintf(w, "%s %s\n", resp, time.Now())
}

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/hello", greet)
	fmt.Println(http.ListenAndServe(":8081", nil))
}

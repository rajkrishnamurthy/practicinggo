package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n", r.Header)
	fmt.Printf("Raw Query Values %v \n", r.URL.RawQuery)
	fmt.Printf("Query Values %v \n", r.URL.Query())

	// _, err := io.Copy(os.Stdout, r.Body)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
}

func main() {
	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

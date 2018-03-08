package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

// Not exported
type Person struct {
	Name string
	Age  int
}

func main() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	indexFile, err := os.Open("html/index.html")
	if err != nil {
		fmt.Println(err)
	}
	index, err := ioutil.ReadAll(indexFile)
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Client subscribed")
		myPerson := Person{
			Name: "Bill",
			Age:  0,
		}

		for {
			time.Sleep(2 * time.Second)
			if myPerson.Age < 40 {
				myJson, err := json.Marshal(myPerson)
				if err != nil {
					fmt.Println(err)
					return
				}
				err = conn.WriteMessage(websocket.TextMessage, myJson)
				if err != nil {
					fmt.Println(err)
					break
				}
				myPerson.Age += 2
			} else {
				conn.Close()
				break
			}
		}
		fmt.Println("Client unsubscribed")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(index))
	})
	http.ListenAndServe(":3000", nil)
}

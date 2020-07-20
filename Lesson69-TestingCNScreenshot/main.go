package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"

	socketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

func main() {
	var method string
	var input = make(map[string]interface{}, 0)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	targetURL := url.URL{
		Scheme: "http",
		Host:   "bankofamerica.com",
	}

	client, err := socketio.Dial(socketio.GetUrl("localhost", 8765, false),
		transport.GetDefaultWebsocketTransport())
	defer client.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = client.On("message", func(h *socketio.Channel, data map[string]interface{}) {
		// dataChannel := make(chan map[string]interface{})
		// defer close(dataChannel)
		// dataChannel <- data
		// log.Printf("--- Got chat message: %s", <-dataChannel)
		log.Printf("--- Got chat message: %s", data)
	})
	if err != nil {
		log.Fatal(err)
	}

	err = client.On(socketio.OnDisconnection, func(h *socketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	err = client.On(socketio.OnConnection, func(h *socketio.Channel) {
		log.Println("Connected")
	})
	if err != nil {
		log.Fatal(err)
	}

	method = "goto"
	input = map[string]interface{}{
		"url": targetURL.String(),
	}
	err = client.Emit(method, input)
	if err != nil {
		log.Fatal(err)
	}

	method = "getCookies"
	input = map[string]interface{}{
		"url": targetURL.String(),
	}
	err = client.Emit(method, input)
	if err != nil {
		log.Fatal(err)
	}

	// for element := range dataChannel {
	// 	log.Printf("Ranging over Channel=%v\n", element)
	// }

	fmt.Printf("The End")
}

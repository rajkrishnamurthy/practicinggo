package main

import (
	"fmt"
	"log"
	"time"

	"../asbclient"
)

func main() {
	log.Printf("Starting")

	var readMessages, writeMessages, deleteMessages = true, false, true

	const (
		azuresb_ns         = "cnSBTest"
		azuresb_policyname = "ManageTest"
		azuresb_policyval  = "bbAtG4S8Pg+XnlIJS5GjPMIdGBbAD4hAN/N+mPocwO4="
		azuresb_queue      = "q1"
	)

	// client := asbclient.New(asbclient.Queue, os.Getenv("sb_namespace"), os.Getenv("sb_key_name"), os.Getenv("sb_key_value"))
	// path := os.Getenv("sb_queue")

	client := asbclient.New(asbclient.Queue, azuresb_ns, azuresb_policyname, azuresb_policyval)
	path := azuresb_queue

	if writeMessages {
		i := 0
		for {

			log.Printf("Send: %d", i)
			err := client.Send(path, &asbclient.Message{
				Body: []byte(fmt.Sprintf("message %d", i)),
			})

			if err != nil {
				log.Printf("Send error: %s", err)
			} else {
				log.Printf("Sent: %d", i)
			}

			time.Sleep(time.Millisecond * 500)
			i++
		}
	}

	if readMessages {
		for {
			log.Printf("Peeking...")
			msg, err := client.PeekLockMessage(path, 30)

			if err != nil {
				log.Printf("Peek error: %s", err)
			} else {
				log.Printf("Peeked message: '%s'", string(msg.Body))
				if deleteMessages {
					err = client.DeleteMessage(msg)
					if err != nil {
						log.Printf("Delete error: %s", err)
					} else {
						log.Printf("Deleted message")
					}
				}
			}

			//time.Sleep(time.Millisecond * 200)
		}
	}
}

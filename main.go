package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	for {
		tweet, err := genTweet()
		if err != nil {
			log.Println(err)
			return
		}
		conn.WriteJSON(tweet)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	http.HandleFunc("/hello", reqHandler)
	http.ListenAndServe(":7000", nil)
}

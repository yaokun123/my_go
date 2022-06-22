package main

import (
	"github.com/jcuga/golongpoll"
	"log"
	"net/http"
)

func main() {
	// This uses the default/empty options. See section on customizing, and Options go docs.
	manager, err := golongpoll.StartLongpoll(golongpoll.Options{})

	// Expose pub-sub. You could omit the publish handler if you don't want
	// to allow clients to publish. For example, if clients only subscribe to data.
	if err == nil {
		http.HandleFunc("/events", manager.SubscriptionHandler)
		http.HandleFunc("/publish", manager.PublishHandler)
		http.ListenAndServe("127.0.0.1:8101", nil)
	} else {
		// handle error creating longpoll manager--typically this means a bad option.
		log.Fatal("init longpoll fail")
	}
}

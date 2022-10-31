package main

import (
	"log"
	"net/http"
	"socialchat/config"
	"socialchat/controllers"
	"socialchat/models"
)

/**
* Handle http request
 */
func handleMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world."))

	w.Header().Set("Access-Control-Allow-Origin", "*")
}

/**
* Chat response handler
 */
func handleRequest(w http.ResponseWriter, r *http.Request) {
	/*
	* All responses in JSON
	 */
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	token := r.FormValue("access-token")
	if token != "" {
		config.GetSession().SetCurrentUser(models.NewProfile(config.GetConnection()).GetByToken(token))
	}

	if config.GetSession().IsLoggedIn() {
		switch r.URL.Path {
		case "/messaging":
			if config.GetSession().IsLoggedIn() {
				cont := controllers.CreateMessaging()
				cont.HandleRequest(w, r)
			} else {
				w.WriteHeader(400)
			}
			break
		}
	} else {
		w.WriteHeader(403)
	}
}

func main() {
	// Main page
	http.HandleFunc("/", handleMessage)
	// Messaging (sending, waiting)
	http.HandleFunc("/messaging", handleRequest)

	sock := controllers.CreateMessagingSocket()
	go sock.Start("/messaging/ws")

	p2p := controllers.CreateMessagingP2P()
	go p2p.Start("/messaging/p2p")

	err := http.ListenAndServe(":81", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

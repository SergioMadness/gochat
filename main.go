package main

import (
	"chat/controllers"
	"fmt"
	"log"
	"net/http"
)

var messages map[string]chan string = make(map[string]chan string)

func handleMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world."))

	w.Header().Set("Access-Control-Allow-Origin", "*")
}

//func parseSignedMessage(message string) string {
//	var result string

//	operCodeParts := extractKeyCode(message)

//	result = operCodeParts[2]

//	return result
//}

//func extractKeyCode(str string) []string {
//	var result []string

//	r, _ := regexp.Compile("([a-zA-Z0-9]+)[[:blank:]]+([0-9]+)")
//	result = r.FindStringSubmatch(str)

//	return result
//}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.URL.Path {
	case "/registration":
		fmt.Println("Registration")
		cont := controllers.CreateRegistration()
		cont.HandleRequest(w, r)
		break
	case "/login":
		fmt.Println("Login")
		cont := controllers.CreateLogin()
		cont.HandleRequest(w, r)
		break
	case "/messaging":
		fmt.Println("Messaging")
		cont := controllers.CreateMessaging()
		cont.HandleRequest(w, r)
		break
	case "/push-message":
		fmt.Println("Push message")
		break
	}
}

func main() {
	http.HandleFunc("/", handleMessage)
	http.HandleFunc("/registration", handleRequest)
	http.HandleFunc("/login", handleRequest)
	http.HandleFunc("/messaging", handleRequest)

	err := http.ListenAndServe(":81", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

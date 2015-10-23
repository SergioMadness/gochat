package main

import (
	"chat/controllers"
	"chat/installer"
	"fmt"
	"log"
	"net/http"
	"os"
)

/**
* Handle http request
 */
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

/**
* Chat response handler
 */
func handleRequest(w http.ResponseWriter, r *http.Request) {
	/*
	* All responses in JSON
	 */
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
	case "/friends/online":
		fmt.Println("Get online friends")
		cont := controllers.CreateFriends()
		cont.GetOnlineUsers(w, r)
		break
	}
}

func consoleCommand(command string) {
	switch command {
	case "install":
		installer.Install()
		break
	case "uninstall":
		installer.Uninstall()
		break
	default:
		fmt.Println("Unknown command")
	}
}

func main() {
	command := os.Args[1]

	if command != "" {
		consoleCommand(command)
	} else {
		// Main page
		http.HandleFunc("/", handleMessage)
		// Registration
		http.HandleFunc("/registration", handleRequest)
		// Login
		http.HandleFunc("/login", handleRequest)
		// Messaging (sending, waiting)
		http.HandleFunc("/messaging", handleRequest)
		// Friends
		http.HandleFunc("/friends/online", handleRequest)

		err := http.ListenAndServe(":81", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}

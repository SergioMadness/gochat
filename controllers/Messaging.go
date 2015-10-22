package controllers

/**
* Messaging service
*/

import (
	"chat/config"
	"chat/models"
	"chat/models/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Messaging struct {
}

func CreateMessaging() *Messaging {
	return new(Messaging)
}

func (m Messaging) HandleRequest(w http.ResponseWriter, r *http.Request) {
	result := response.NewMessageResponse(0, "")

	message := r.FormValue("msg")
	fromStr := r.FormValue("from")
	toStr := r.FormValue("to")

	if fromStr == "" {
		w.WriteHeader(400)
		result.Result = 400
		result.ResultMessage = "Who are you?"
	} else {
		from, errFrom := strconv.Atoi(fromStr)
		to, errTo := strconv.Atoi(toStr)

		if errFrom == nil {
			fmt.Println("Get channels")

			chFrom := models.GetMessageChannelWrapper().GetChannel(from, true)
			chTo := models.GetMessageChannelWrapper().GetChannel(to, false)

			if message != "" {
				var messagePull []models.Message

				fmt.Println("Message not empty")
				if chTo != nil {
					fmt.Println("Undelivered")

					mess := models.NewMessage(config.GetConnection())
					messagePull = mess.GetUndeliveredMessages(to)

					var newMessage models.Message
					newMessage.From = from
					newMessage.To = to
					newMessage.Body = message
					messagePull = append(messagePull, newMessage)

					mess.RemoveUndeliveredMessages(to)

					fmt.Println("chan" + string(to) + " exists")

					fmt.Println("message: " + message)

					jsonResult, err := json.Marshal(messagePull)

					if err == nil {
						chTo <- string(jsonResult)
					}
				} else {
					messageDB := models.NewMessage(config.GetConnection())
					messageDB.Body = message
					messageDB.From = from
					messageDB.To = to
					messageDB.CreateDate = time.Now().Local().Unix()
					messageDB.IsRead = false
					messageDB.Save()
				}
			} else {
				fmt.Println("reconnect")
				io.WriteString(w, <-chFrom)
			}
		} else {
			if errFrom != nil {
				fmt.Println(errFrom.Error())
			}
			if errTo != nil {
				fmt.Println(errTo.Error())
			}
			w.WriteHeader(400)
		}
	}

	//	jsonResult, err := json.Marshal(result)

	//	if err != nil {
	//		w.WriteHeader(500)
	//	} else {
	//		w.Write(jsonResult)
	//	}
}

func (m Messaging) pushUndeliveredMessages(c chan string, to string) {

}

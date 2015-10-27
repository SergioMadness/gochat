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

func (m *Messaging) HandleRequest(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("msg")
	fromStr := r.FormValue("from")
	toStr := r.FormValue("to")

	if fromStr == "" {
		result := response.NewMessageResponse(0, "")

		w.WriteHeader(400)
		result.Result = 400
		result.ResultMessage = "Who are you?"
	} else {
		from, errFrom := strconv.Atoi(fromStr)
		to, errTo := strconv.Atoi(toStr)

		if errFrom == nil {
			fmt.Println("Get channels")

			chFrom := models.GetMessageChannelWrapper().GetChannel(from, false)
			if chFrom == nil {
				go m.notifyChannels(from)
				chFrom = models.GetMessageChannelWrapper().GetChannel(from, true)
			}
			chTo := models.GetMessageChannelWrapper().GetChannel(to, false)

			if message != "" {
				go m.sendMessage(w, to, from, message, chTo)
			} else {
				m.updateConnection(w, from, chFrom)
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
}

func (m *Messaging) sendMessage(w http.ResponseWriter, to, from int, message string, chTo chan string) {
	if chTo != nil {
		fmt.Println("Undelivered")

		messageResponse := response.NewMessageResponse(0, "")

		mess := models.NewMessage(config.GetConnection())
		messagePull := mess.GetUndeliveredMessages(to)

		var newMessage models.Message
		newMessage.From = from
		newMessage.To = to
		newMessage.Body = message
		messagePull = append(messagePull, newMessage)

		mess.RemoveUndeliveredMessages(to)

		messageResponse.Data = messagePull

		fmt.Println("chan" + string(to) + " exists")

		fmt.Println("message: " + message)

		jsonResult, err := json.Marshal(messageResponse)

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
}

func (m *Messaging) updateConnection(w http.ResponseWriter, id int, chFrom chan string) {
	fmt.Println("reconnect")

	io.WriteString(w, <-chFrom)
}

func (m *Messaging) notifyChannels(id int) {
	chanels := models.GetMessageChannelWrapper().GetChannels()

	profileO := models.NewProfile(config.GetConnection())

	profile := profileO.GetById(id)

	if profile != nil && profile.Id > 0 {
		newOnlineFriendResponse := response.NewOnlineFriendResponse(0, "")
		newOnlineFriendResponse.Data = profile
		jsonResult, err := json.Marshal(newOnlineFriendResponse)

		if err == nil {
			for key, channel := range chanels {
				if key != id {
					channel <- string(jsonResult)
				}
			}
		}
	}
}

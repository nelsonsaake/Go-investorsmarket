package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"projects/investorsmarket/chats/repo"
	"strconv"
	"time"
)

type createChatRequest struct {
	SenderId   uint64
	ReceiverId uint64
	Message    string
	SentAt     time.Time
}

func (ccReq createChatRequest) isValid() bool {

	if ccReq.SenderId == 0 || ccReq.ReceiverId == 0 {
		return false
	}

	if ccReq.SenderId == ccReq.ReceiverId {
		fmt.Println("Error sender and receiver are the same")
		return false
	}

	return true
}

func CreateChat(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		ccReq := createChatRequest{}
		err = json.Unmarshal(payload, &ccReq)
		if err != nil {
			fmt.Printf("Error parsing payload body: payload %v,\n err %v \n ", string(payload), err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		if !ccReq.isValid() {
			fmt.Println("Error: request is in valid")
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		chat, err := repo.CreateChat(ccReq.SenderId, ccReq.ReceiverId, ccReq.Message, ccReq.SentAt)
		if err != nil {
			fmt.Printf("Error creating chat: %v\n", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		w.Header().Add("Location", "/chats/"+fmt.Sprint(chat.ID))

		formatter.JSON(w, http.StatusCreated, chat)
	}
}

func GetChatCounterPartsOf(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		idstr := vars["id"]

		id, err := strconv.ParseUint(idstr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing id string: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		counterparts, err := repo.GetChatCounterPartsOf(id)
		if err != nil {
			fmt.Printf("Error getting chat counter parts of %d: \n\t %v\n", id, err)
		}
		
		formatter.JSON(w, http.StatusOK, counterparts)
	}
}

func GetChats(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		u1Str := vars["u1"]
		u2Str := vars["u2"]

		u1, err := strconv.ParseUint(u1Str, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing user 1 id string: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		u2, err := strconv.ParseUint(u2Str, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing user 2 id string: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		chats, err := repo.GetChats(u1,u2)
		if err != nil {
			fmt.Printf("Error getting chats of %d and %d: \n\t %v\n", u1, u2, err)
		}
		
		formatter.JSON(w, http.StatusOK, chats)
	}
}

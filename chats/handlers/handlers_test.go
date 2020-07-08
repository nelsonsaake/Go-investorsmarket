package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/matryer/is"
	"github.com/unrolled/render"

	"fmt"
	"projects/investorsmarket/models"
	"strings"
	"time"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(
		render.Options{
			IndentJSON: true,
		},
	)

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {

	mx.HandleFunc("/chats/counterpartsof/{id}", GetChatCounterPartsOf(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/chats/chats/{u1}/{u2}", GetChats(formatter)).Methods(http.MethodGet)
}

func TestCreateChat(t *testing.T) {

	return

	is := is.New(t)

	client := &http.Client{}

	formatter := render.New(render.Options{IndentJSON: true})
	server := httptest.NewServer(http.HandlerFunc(CreateChat(formatter)))
	defer server.Close()

	reqBody := []byte(
	`
		{
			"senderId": 3,
			"receiverId": 1,
			"message": "Hey",
			"sentAt": "` + time.Now().Format(time.RFC3339) + `"
		}
	`)
	fmt.Println()

	req, err := http.NewRequest(http.MethodPost, server.URL, bytes.NewBuffer(reqBody))
	is.NoErr(err) // creating request err

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	is.NoErr(err) // error sending request

	var (
		chat                           = models.Chat{}
		expectedStatus                 = http.StatusCreated
		expectedSenderId        uint64 = 3
		expectedReceiverId      uint64 = 1
		expectedMessage                = "Hey"
		expectedLocationContent        = "/chats/"
	)

	is.Equal(res.StatusCode, expectedStatus)

	if location, ok := res.Header["Location"]; !ok {
		t.Errorf("Error: location header not set")
		return
	} else {
		if len(location) == 0 {
			t.Errorf("Error: location empty")
			return
		} else {
			if !strings.Contains(location[0], expectedLocationContent) {
				t.Errorf("Error: location header does not have the right format; expected location header to contain: %v", expectedLocationContent)
			}
		}
	}

	payload, err := ioutil.ReadAll(res.Body)
	is.NoErr(err)
	defer res.Body.Close()

	err = json.Unmarshal(payload, &chat)
	is.NoErr(err)

	if chat.SenderId != expectedSenderId {
		t.Errorf("Error with the sender id: expected %d, received %d \n", expectedSenderId, chat.SenderId)
		return
	}

	if chat.ReceiverId != expectedReceiverId {
		t.Errorf("Error with the receiver id: expected %d, received %d \n", expectedReceiverId, chat.ReceiverId)
		return
	}

	if chat.Message != expectedMessage {
		t.Errorf("Error with the received message: expected %v, received %v \n", expectedMessage, chat.Message)
		return
	}
}

func TestGetChatCounterPartsOf(t *testing.T){

	var (
		url      = "/chats/counterpartsof/1"
		req, err = http.NewRequest(http.MethodGet, url, nil)
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating a new request for getting all the chat counter parts(other users) a particular user has chat(s) with: %v", err)
		return
	}
	server.ServeHTTP(res, req)
	
	var (
		expectedStatus = http.StatusOK
		counterParts = []models.User{}
		expectedName = "Nelson"
	)
	
	if res.Code != expectedStatus {
		t.Errorf("Error with the response code: expected %v, received %v ", expectedStatus, res.Code)
	}
	
	payload, err := ioutil.ReadAll(res.Body) 
	if err != nil {
		t.Errorf("Error reading the response body: %v\n", err)
		return
	}
	
	err = json.Unmarshal(payload, &counterParts)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return 
	}
	
	if len(counterParts) == 0 {
		t.Errorf("Error with the number of counters received: received nothing, expected at least one")
		return 
	}

	counterPart := counterParts[0]
	
	if counterPart.Name != expectedName {
		t.Errorf("Error with a counter part name: expected %s, received %s \n", expectedName, counterPart.Name)
		return
	}
}

func TestGetChats(t *testing.T) {
	
	var (
		url      = "/chats/chats/1/3"
		req, err = http.NewRequest(http.MethodGet, url, nil)
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating a new request for getting all chats between two users: %v", err)
		return
	}
	server.ServeHTTP(res, req)
	
	var (
		expectedStatus = http.StatusOK
		chats = []models.Chat{}
		expectedMessage = "Hello"
	)

	if res.Code != expectedStatus {
		t.Errorf("Error with the response code: expected %v, received %v ", expectedStatus, res.Code)
		return
	}
	
	payload, err := ioutil.ReadAll(res.Body) 
	if err != nil {
		t.Errorf("Error reading the response body: %v\n", err)
		return
	}

	err = json.Unmarshal(payload, &chats)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v\n", string(payload), err)
		return 
	}

	if len(chats) == 0 {
		t.Errorf("Error with the number of chats received: received nothing, expected at least one")
		return 
	}

	chat := chats[0]
	
	if chat.Message != expectedMessage {
		t.Errorf("Error with the first chat message: expected %s, received %s \n", expectedMessage, chat.Message)
		return
	}

	fmt.Println(string(payload))
}
package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
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

	mx.HandleFunc("/auths", CreateAuth(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/auths/ep", GetAuthGivenEP(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/auths/t1", GetAuthGivenToken(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/auths/ut", GetAuthGivenUserAndToken(formatter)).Methods(http.MethodPost)
}

func TestCreateAuth(t *testing.T) {

	return

	var (
		url     = "/auths"
		reqBody = []byte(`
		{
			"email": "nelsonsaakekofi@gmail.com",
			"password": "something"
		}`)
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	server.ServeHTTP(res, req)

	var (
		expectedStatus               = http.StatusCreated
		auth                         = authResponse{}
		expectedUserId        uint64 = 10
		expectedLeastTokenLen        = 32
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

	err = json.Unmarshal(payload, &auth)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return
	}

	if auth.UserId != expectedUserId {
		t.Errorf("Error : expected user id %d, received %d \n", expectedUserId, auth.UserId)
		return
	}

	if len(auth.Token) < expectedLeastTokenLen {
		t.Errorf("Error with token: code lenght too short!")
		return
	}
}

func TestGetAuthGivenEP(t *testing.T) {

	return

	var (
		url     = "/auths/ep"
		reqBody = []byte(`
		{
			"email": "nelsonsaakekofi@gmail.com",
			"password": "something"
		}`)
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	server.ServeHTTP(res, req)

	var (
		expectedStatus               = http.StatusOK
		auth                         = authResponse{}
		expectedUserId        uint64 = 10
		expectedLeastTokenLen        = 32
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

	err = json.Unmarshal(payload, &auth)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return
	}

	if auth.UserId != expectedUserId {
		t.Errorf("Error : expected user id %d, received %d \n", expectedUserId, auth.UserId)
		return
	}

	if len(auth.Token) < expectedLeastTokenLen {
		t.Errorf("Error with token: code lenght too short!")
		return
	}

	// fmt.Println(string(payload))
}

func TestGetAuthGivenToken(t *testing.T) {

	return

	var (
		url     = "/auths/t1"
		reqBody = []byte(`
		{
			"token": "m6VjDK9AbBi4UsAwKoL8IdQOuK6cnfImEHzt1sawYzxErGZ5FQbHDW869Dg5RCsuSwmDodMUcfjA1D6ve7IpGrjKFVYqIiSq88J0W7du8moh1rFbB6Zxrri8I5e9e5wC0gWRG"
		}`)
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	server.ServeHTTP(res, req)

	var (
		expectedStatus               = http.StatusOK
		auth                         = authResponse{}
		expectedUserId        uint64 = 10
		expectedLeastTokenLen        = 32
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

	err = json.Unmarshal(payload, &auth)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return
	}

	if auth.UserId != expectedUserId {
		t.Errorf("Error : expected user id %d, received %d \n", expectedUserId, auth.UserId)
		return
	}

	if len(auth.Token) < expectedLeastTokenLen {
		t.Errorf("Error with token: code lenght too short!")
		return
	}

	// fmt.Println(string(payload))
}

func TestGetAuthGivenUserAndToken(t *testing.T) {

	return

	var (
		url     = "/auths/ut"
		reqBody = []byte(`
		{
			"userId": 10,
			"token": "m6VjDK9AbBi4UsAwKoL8IdQOuK6cnfImEHzt1sawYzxErGZ5FQbHDW869Dg5RCsuSwmDodMUcfjA1D6ve7IpGrjKFVYqIiSq88J0W7du8moh1rFbB6Zxrri8I5e9e5wC0gWRG"
		}`)
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	server.ServeHTTP(res, req)

	var (
		expectedStatus               = http.StatusOK
		auth                         = authResponse{}
		expectedUserId        uint64 = 10
		expectedLeastTokenLen        = 32
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

	err = json.Unmarshal(payload, &auth)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return
	}

	if auth.UserId != expectedUserId {
		t.Errorf("Error : expected user id %d, received %d \n", expectedUserId, auth.UserId)
		return
	}

	if len(auth.Token) < expectedLeastTokenLen {
		t.Errorf("Error with token: code lenght too short!")
		return
	}

	fmt.Println(string(payload))
}

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
	"github.com/unrolled/render"

	"fmt"
	"projects/investorsmarket/models"
	"strings"
	_ "time"
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

	mx.HandleFunc("/changepasswords", CreateChangePassword(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/changepasswords", GetAllChangePasswords(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/changepasswords/g/1", GetChangePassword(formatter)).Methods(http.MethodPost)
}

func TestCreateChangePassword(t *testing.T) {

	return

	var (
		url      = "/changepasswords"
		reqBody  = []byte(`{"userId": 1}`)
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating a new requesting for change of password: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	server.ServeHTTP(res, req)

	var (
		expectedStatus                 = http.StatusCreated
		expectedLocationContent        = "/changepasswords/"
		cp                             = models.ChangePassword{}
		expectedActive                 = true
		expectedUserId          uint64 = 1
		expectedLeastCodeLen           = 32
	)

	if res.Code != expectedStatus {
		t.Errorf("Error with the response code: expected %v, received %v ", expectedStatus, res.Code)
		return
	}

	if location, ok := res.Header()["Location"]; !ok {
		t.Errorf("Error: location header not set")
		fmt.Println("Response header", res.Header())
		return
	} else {
		if len(location) == 0 {
			t.Errorf("Error: location header set but empty")
			return
		} else {
			if !strings.Contains(location[0], expectedLocationContent) {
				t.Errorf("Error: location header does not have the right format; expected location header to contain: %v", expectedLocationContent)
				return
			}
		}
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading the response body: %v\n", err)
		return
	}

	err = json.Unmarshal(payload, &cp)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return
	}

	if cp.Active != expectedActive {
		t.Errorf("Error with the created change passowrd record: expected is to be active but it's not!")
		return
	}

	if cp.UserId != expectedUserId {
		t.Errorf("Error with the created change passowrd record: expected user id %d, received %d \n", expectedUserId, cp.UserId)
		return
	}

	if len(cp.Code) < expectedLeastCodeLen {
		t.Errorf("Error with the created change passowrd record: code lenght too short!")
		return
	}
}

func TestGetChangePassword(t *testing.T) {

	return

	var (
		url      = "/changepasswords/g/1"
		code     = "2sfT9yYVVmtDALDjpAjhMOrc5UZhYQDMoesZQNaxyzyABCdv8sggiBL1Vna1Iwd1fKCccIWl4QXyuKcUtGz5bbbl"
		reqBody  = []byte(`{"userId": 1, "code": "` + code + `"}`)
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error getting change password record: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	server.ServeHTTP(res, req)

	var (
		expectedStatus        = http.StatusOK
		cp                    = models.ChangePassword{}
		expectedActive        = true
		expectedUserId uint64 = 1
		expectedCode          = code
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

	err = json.Unmarshal(payload, &cp)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return
	}

	if cp.Active != expectedActive {
		t.Errorf("Error with change passowrd record: expected is to be active but it's not!")
		return
	}

	if cp.UserId != expectedUserId {
		t.Errorf("Error with change passowrd record: expected user id %d, received %d \n", expectedUserId, cp.UserId)
		return
	}

	if cp.Code != expectedCode {
		t.Errorf("Error with change passowrd record: code lenght too short!")
		return
	}
}

func TestGetAllChangePasswords(t *testing.T) {

	var (
		url      = "/changepasswords"
		req, err = http.NewRequest(http.MethodGet, url, nil)
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error getting all change password records: %v", err)
		return
	}
	server.ServeHTTP(res, req)

	var (
		expectedStatus        = http.StatusOK
		cps                   = []models.ChangePassword{}
		expectedActive        = false
		expectedUserId uint64 = 1
		expectedCode          = "2sfT9JQWg8iy3cMiDabwz2xPBCyzwSyspNidI55EB5YORquRwO329Vwwe5r4DNH72MOE8bcGMK8j9tK7ohahoLX1"
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

	err = json.Unmarshal(payload, &cps)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return
	}

	if len(cps) == 0 {
		t.Errorf("Error: didn't get any change passowrd record")
		return
	}

	cp := cps[0]

	if cp.Active != expectedActive {
		t.Errorf("Error with the first change passowrd record received: expected is to be active but it's not!")
		return
	}

	if cp.UserId != expectedUserId {
		t.Errorf("Error with the first change passowrd record received: expected user id %d, received %d \n", expectedUserId, cp.UserId)
		return
	}

	if cp.Code != expectedCode {
		t.Errorf("Error with the first change passowrd record received: code lenght too short!")
		return
	}
}

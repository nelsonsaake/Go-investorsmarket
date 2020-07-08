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

	mx.HandleFunc("/users", CreateUser(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/users/{id}", UpdateUser(formatter)).Methods(http.MethodPut)
}

func TestCreateUser(t *testing.T) {
	
	return
	var (
		url      = "/users"
		reqBody  = []byte(`{"email": "nelsonsaake@gmail.com","password": "something"}`)
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	server.ServeHTTP(res, req)

	var (
		expectedStatus          = http.StatusCreated
		expectedLocationContent = "/users/"
		user                    = models.User{}
		expectedEmail           = "nelsonsaake@gmail.com"
		expectedPassword        = "something"
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

	err = json.Unmarshal(payload, &user)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return
	}

	if user.Email != expectedEmail {
		t.Errorf("\nexpected %s, received %s\n", expectedEmail, user.Email)
	}

	if user.Password != expectedPassword {
		t.Errorf("\nexpected %s, received %s\n", expectedPassword, user.Password)
	}
}

func TestUpdateUser(t *testing.T) {

	return
	var (
		url     = "/users/10"
		reqBody = []byte(`
			{
				"email": "nelsonsaakekofi@gmail.com",
				"password": "password",
				"picture": "/images/profiles/nelsonsaakekofi@gmail.com.jpg",
				"firstname": "Nelson",
				"surname": "Saake",
				"dateOfBirth": "",
				"gender": "male",
				"phoneNumber": "0548876758",
				"nationality": "Ghanaian",
				"occupation": "Student",
				"address": "PC Homes, Agric Hill",
				"country": "Ghana",
				"region": "Western",
				"city": "Tarkwa",
				"accName": "Nelson Kofi Saake",
				"accNumber": "12345678765432",
				"accBankName": "Access",
				"nkSurname": "Saake",
				"nkFirstname": "Rowland",
				"nkRelationship": "Bother",
				"nkEmail": "",
				"nkPhoneNumber": "",
				"nkAddress": ""
			}
		`)
		req, err = http.NewRequest(http.MethodPut, url, bytes.NewBuffer(reqBody))
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	server.ServeHTTP(res, req)

	var (
		expectedStatus    = http.StatusOK
		user              = models.User{}
		expectedEmail     = "nelsonsaakekofi@gmail.com"
		expectedPassword  = "something"
		expectedBankName  = "Access"
		expectedNkSurname = "Saake"
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

	err = json.Unmarshal(payload, &user)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return
	}

	if user.Email != expectedEmail {
		t.Errorf("\nexpected %s, received %s\n", expectedEmail, user.Email)
		return
	}

	if user.Password != expectedPassword {
		t.Errorf("\nexpected %s, received %s\n", expectedPassword, user.Password)
		return
	}

	if user.NkSurname != expectedNkSurname {
		t.Errorf("\nexpected %s, received %s\n", expectedNkSurname, user.NkSurname)
		return
	}

	if user.AccBankName != expectedBankName {
		t.Errorf("\nexpected %s, received %s\n", expectedBankName, user.AccBankName)
		return
	}
	
	fmt.Println(string(payload))
}

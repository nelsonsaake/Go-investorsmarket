package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"projects/investorsmarket/models"
	"strings"
	"testing"
)

func TestCreateOpportunity(t *testing.T) {

	// done with this
	 return

	//
	var (
		client    = &http.Client{}
		formatter = render.New(render.Options{IndentJSON: true})
		server    = httptest.NewServer(http.HandlerFunc(CreateOpportunity(formatter)))
		reqBody   = []byte(`
			{
				"name": "Stew Maker",
				"amount": 1005000.00,
				"industry": "Technology, Internet Of Things",
				"description": "Stew maker is like the rice cooker but makes stew. A complementary device for the rice cooker.",
				"userId" : 10,
				"picture": "/images/profiles/opportunities/1.jpg",
				"returns": 0.08,
				"duration": 1,
				"location": "Tarkwa"
			}
		`)
		req, err = http.NewRequest(http.MethodPost, server.URL, bytes.NewBuffer(reqBody))
	)
	defer server.Close()
	if err != nil {
		t.Errorf("Error creating a POST request for CreateOpportunity: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending POST request to CreateOpportunity: %v", err)
	}

	//
	var (
		expectedStatus                       = http.StatusCreated
		opportunity                          = models.Opportunity{}
		name                                 = "Stew Maker"
		amount                               = 1005000.00
		userId                        uint64 = 10
		expectedLocationHeaderContent        = "/opportunities/"
		expectedLoc = "Tarkwa"
	)

	if response.StatusCode != expectedStatus {
		t.Errorf("Error, expected %v, got %v ", expectedStatus, response.Status)
	}

	if location, ok := response.Header["Location"]; !ok {
		t.Errorf("Error, location header is not set.")
	} else {
		if len(location[0]) == 0 {
			t.Errorf("Error, location is empty")
		} else {
			if !strings.Contains(location[0], expectedLocationHeaderContent) {
				t.Errorf("Error with location header, expecting location header to contain %s, location header received : %s", expectedLocationHeaderContent, location[0])
			}
		}
	}

	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	defer response.Body.Close()

	err = json.Unmarshal(payload, &opportunity)
	if err != nil {
		t.Errorf("Error unmarshalling response: %v", err)
	}

	if name != opportunity.Name {
		t.Errorf("Error with the opportunity name: expected %v, got %v: ", name, opportunity.Name)
	}

	if amount != opportunity.Amount {
		t.Errorf("Error with the amount: expected %v, got %v: ", amount, opportunity.Amount)
	}

	if userId != opportunity.UserId {
		t.Errorf("Error with the user id: expected %v, got %v: ", userId, opportunity.UserId)
	}
	
	if expectedLoc != opportunity.Location {
		t.Errorf("Error with the location of opportunity: expected %v, got %v: ", expectedLoc, opportunity.Location)
	}
}

func TestGetAllOpportunities(t *testing.T) {

	return

	var (
		client    = &http.Client{}
		formatter = render.New(render.Options{IndentJSON: true})
		server    = httptest.NewServer(http.HandlerFunc(GetAllOpportunities(formatter)))
		reqBody   = []byte("")
		req, err  = http.NewRequest(http.MethodGet, server.URL, bytes.NewBuffer(reqBody))
	)
	defer server.Close()
	if err != nil {
		t.Errorf("Error creating a get request for GetAllOpportunities: %v", err)
		return
	}

	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending request to GetAllOpportunities: %v", err)
		return
	}

	// expecteds
	var (
		expectedStatus               = http.StatusOK
		opportunities                = []models.Opportunity{}
		expectedFirstOpportunityName = "Darvoc"
	)

	if response.StatusCode != expectedStatus {
		t.Errorf("Error with the status code, expected %v, got %v: ", expectedStatus, response.Status)
		return
	}

	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}
	defer response.Body.Close()

	err = json.Unmarshal(payload, &opportunities)
	if err != nil {
		t.Errorf("Error unmarshalling payload: %v", err)
		return
	}

	if len(opportunities) == 0 {
		t.Errorf("Error with the available opportunities, didn't get any opportunities")
		return
	}

	if opportunities[0].Name != expectedFirstOpportunityName {
		t.Errorf("Error with the first opportunity name: expected %s got %s", expectedFirstOpportunityName, opportunities[0].Name)
		return
	}

	fmt.Println("")
}

func TestGetOpportunityCreators(t *testing.T) {

	return

	var (
		client    = &http.Client{}
		formatter = render.New(render.Options{IndentJSON: true})
		server    = httptest.NewServer(http.HandlerFunc(GetOpportunityCreators(formatter)))
		reqBody   = []byte("")
		req, err  = http.NewRequest(http.MethodGet, server.URL, bytes.NewBuffer(reqBody))
	)
	defer server.Close()
	if err != nil {
		t.Errorf("Error creating a get request for GetOpportunityCreators: %v", err)
		return
	}

	var (
		expectedStatus = http.StatusOK
		creators       = []models.User{}
	)

	response, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending request to GetOpportunityCreators: %v", err)
		return
	}

	if response.StatusCode != expectedStatus {
		t.Errorf("Error with response status: expected %v, received %v \n", expectedStatus, response.Status)
		return
	}

	payload, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	err = json.Unmarshal(payload, &creators)
	if err != nil {
		t.Errorf("Error unmarshalling payload: %v", err)
		return
	}

	if len(creators) == 0 {
		t.Errorf("Error: no users found!")
		return
	}
}

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

	mx.HandleFunc("/opportunities/creators/{id}/history", GetCreatorHistory(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/opportunities/{id}", GetOpportunity(formatter)).Methods(http.MethodGet)
}

func TestGetCreatorHistory(t *testing.T) {

	return

	var (
		creatorId = "3"
		reqBody   = []byte("")
		reqURL    = "/opportunities/creators/" + creatorId + "/history"
		req, err  = http.NewRequest(http.MethodGet, reqURL, bytes.NewBuffer(reqBody))
		res       = httptest.NewRecorder()
		server    = NewServer()
	)
	if err != nil {
		t.Errorf("Error making a new request for getting creators history: %v", err)
	}

	server.ServeHTTP(res, req)

	// expecteds
	var (
		expectedStatus               = http.StatusOK
		history                      = []models.Opportunity{}
		expectedLen                  = 1
		expectedFirstOpportunityName = "Darvoc"
	)

	if res.Code != expectedStatus {
		t.Errorf("Error with the response status: expected %v, received %v", expectedStatus, res.Code)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	err = json.Unmarshal(payload, &history)
	if err != nil {
		t.Errorf("Error unmarshalling payload: %v", err)
		return
	}

	if len(history) != expectedLen {
		t.Errorf("Error with the length of the history: expected %d, received %d", expectedLen, len(history))
		return
	}

	fmt.Println(history)

	if history[0].Name != expectedFirstOpportunityName {
		t.Errorf("Error with the name of the first opportunity in the creators history: expected %s, received %s", expectedFirstOpportunityName, history[0].Name)
		return
	}
}

func TestGetOpportunity(t *testing.T) {

	var (
		id = "1"
		reqURL    = "/opportunities/" + id
		req, err  = http.NewRequest(http.MethodGet, reqURL, nil)
		res       = httptest.NewRecorder()
		server    = NewServer()
	)
	if err != nil {
		t.Errorf("Error making a new request: %v", err)
	}

	server.ServeHTTP(res, req)

	// expecteds
	var (
		expectedStatus               = http.StatusOK
		opportunity                  = models.Opportunity{}
		expectedFirstOpportunityName = "Darvoc"
	)

	if res.Code != expectedStatus {
		t.Errorf("Error with the response status: expected %v, received %v", expectedStatus, res.Code)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	err = json.Unmarshal(payload, &opportunity)
	if err != nil {
		t.Errorf("Error unmarshalling payload: %v", err)
		return
	}

	if opportunity.Name != expectedFirstOpportunityName {
		t.Errorf("Error with the name of opportunity: expected %s, received %s", expectedFirstOpportunityName, opportunity.Name)
		return
	}
}

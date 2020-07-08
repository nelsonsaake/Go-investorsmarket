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

	mx.HandleFunc("/investments/investors/{id}", GetInvestmentsByInvestor(formatter)).Methods(http.MethodGet)
	mx.HandleFunc("/investments/{id}", GetInvestment(formatter)).Methods(http.MethodGet)
}

func TestGetInvestmentsByInvestor(t *testing.T) {

	return

	var (
		url      = "/investments/investors/1"
		req, err = http.NewRequest(http.MethodGet, url, nil)
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating a new request for getting investments made by investor: %v", err)
		return
	}
	server.ServeHTTP(res, req)

	var (
		expectedStatus                = http.StatusOK
		investments                   = []models.Investment{}
		expectedUserId        uint64  = 1
		expectedOpportunityId uint64  = 3
		expectedBoughtAmount  float64 = 10000
	)

	if res.Code != expectedStatus {
		t.Errorf("Error with the response code: expected %v, got %v", expectedStatus, res.Code)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	err = json.Unmarshal(payload, &investments)
	if err != nil {
		t.Errorf("Error unmarshalling payload: %v", err)
		return
	}

	if len(investments) == 0 {
		t.Errorf("Found no investments made by this investor")
		return
	}

	if investments[0].UserId != expectedUserId {
		t.Errorf("Expecting user id of first investment to be: %v, received %v %v", expectedUserId, investments[0].UserId, "\n")
		return
	}

	if investments[0].OpportunityId != expectedOpportunityId {
		t.Errorf("Expecting opportunity id of first investment to be: %v, received %v %v", expectedOpportunityId, investments[0].OpportunityId, "\n")
		return
	}

	if investments[0].AmountBought != expectedBoughtAmount {
		t.Errorf("Expecting amount bought of first investment to be: %v, received %v %v", expectedBoughtAmount, investments[0].AmountBought, "\n")
		return
	}
}

func TestCreateInvestment(t *testing.T) {

	return

	//
	var (
		client    = &http.Client{}
		formatter = render.New(render.Options{IndentJSON: true})
		reqBody   = []byte(
			`
				{
					"userId": 1,
					"opportunityId": 3,
					"amountBought": 10000
				}
			`,
		)
		server   = httptest.NewServer(http.HandlerFunc(CreateInvestment(formatter)))
		req, err = http.NewRequest(http.MethodPost, server.URL, bytes.NewBuffer(reqBody))
	)
	defer server.Close()
	if err != nil {
		t.Errorf("Error creating a new request for create investments: %v", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending request to create a new request investment: %v", err)
	}

	// expecteds
	var (
		investment                      = models.Investment{}
		expectedStatus                  = http.StatusCreated
		expectedUserId          uint64  = 1
		expectedOpportunityId   uint64  = 3
		expectedAmount          float64 = 10000
		expectedLocationContent         = "/investments/"
	)

	if res.StatusCode != expectedStatus {
		t.Errorf("Error with the status code: expected %v, received %v", expectedStatus, res.Status)
		return
	}

	// test for location header
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
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}
	defer res.Body.Close()

	err = json.Unmarshal(payload, &investment)
	if err != nil {
		t.Errorf("Error unmarshalling payload: %v", err)
		return
	}

	if expectedUserId != investment.UserId {
		t.Errorf("Error with the investment id: expected %v, received %v", expectedUserId, investment.UserId)
		return
	}

	if investment.OpportunityId != expectedOpportunityId {
		t.Errorf("Error with the opportunity id: expected %v, received %v", expectedOpportunityId, investment.OpportunityId)
		return
	}

	if investment.AmountBought != expectedAmount {
		t.Errorf("Error with the amount: expected %v, received %v", expectedAmount, investment.AmountBought)
		return
	}
}

func TestGetInvestment(t *testing.T) {
	
	return

	var (
		url      = "/investments/8"
		req, err = http.NewRequest(http.MethodGet, url, nil)
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating a new request for getting an investment: %v", err)
		return
	}
	server.ServeHTTP(res, req)

	var (
		expectedStatus                = http.StatusOK
		expectedUserId        uint64  = 1
		expectedOpportunityId uint64  = 3
		expectedAmount        float64 = 10000
		investment            models.Investment
	)

	if res.Code != expectedStatus {
		t.Errorf("Error with status code: expected %v, received %v", expectedStatus, res.Code)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}

	err = json.Unmarshal(payload, &investment)
	if err != nil {
		t.Errorf("Error unmarshalling payload: %v", err)
		return
	}

	if expectedUserId != investment.UserId {
		t.Errorf("Error with the investment id: expected %v, received %v", expectedUserId, investment.UserId)
		return
	}

	if investment.OpportunityId != expectedOpportunityId {
		t.Errorf("Error with the opportunity id: expected %v, received %v", expectedOpportunityId, investment.OpportunityId)
		return
	}

	if investment.AmountBought != expectedAmount {
		t.Errorf("Error with the amount: expected %v, received %v", expectedAmount, investment.AmountBought)
		return
	}
}

func TestGetAllInvestments(t *testing.T){
	
	return 
	
	//
	var (
		client    = &http.Client{}
		formatter = render.New(render.Options{IndentJSON: true})
		server   = httptest.NewServer(http.HandlerFunc(GetAllInvestments(formatter)))
		req, err = http.NewRequest(http.MethodPost, server.URL, nil)
	)
	defer server.Close()
	if err != nil {
		t.Errorf("Error creating a new request for getting all investments: %v", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending request to get all request investments: %v", err)
	}

	// expecteds
	var (
		investments                      = []models.Investment{}
		investment                     = models.Investment{}
		expectedStatus                  = http.StatusOK
		expectedUserId          uint64  = 1
		expectedOpportunityId   uint64  = 3
		expectedAmount          float64 = 10000
	)

	if res.StatusCode != expectedStatus {
		t.Errorf("Error with the status code: expected %v, received %v", expectedStatus, res.Status)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		return
	}
	defer res.Body.Close()

	err = json.Unmarshal(payload, &investments)
	if err != nil {
		t.Errorf("Error unmarshalling payload: %v", err)
		return
	}

	if len(investments) == 0 {
		t.Errorf("Error: no investments was found!")
		return
	}

	investment = investments[0]

	if expectedUserId != investment.UserId {
		t.Errorf("Error with the investment id: expected %v, received %v", expectedUserId, investment.UserId)
		return
	}

	if investment.OpportunityId != expectedOpportunityId {
		t.Errorf("Error with the opportunity id: expected %v, received %v", expectedOpportunityId, investment.OpportunityId)
		return
	}

	if investment.AmountBought != expectedAmount {
		t.Errorf("Error with the amount: expected %v, received %v", expectedAmount, investment.AmountBought)
		return
	}
}

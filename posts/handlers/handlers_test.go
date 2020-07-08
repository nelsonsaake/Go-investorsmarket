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
	
	mx.HandleFunc("/posts", CreatePost(formatter)).Methods(http.MethodPost)
	mx.HandleFunc("/posts/opportunity/{id}", GetOpportunityPosts(formatter)).Methods(http.MethodGet)
}

func TestCreatePost(t *testing.T) {
	
	return
	var (
		url      = "/posts"
		reqBody  = []byte(`
			{
				"picture": "/images/posts/opportunity1.2.jpg",
				"description": "Request stage!",
				"opportunityId": 1
			}
		`)
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
		expectedStatus                 = http.StatusCreated
		expectedLocationContent        = "/posts/"
		post                           = models.Post{}
		expectedOpportunityId uint64 = 1
	)
	
	if res.Code != expectedStatus {
		t.Errorf("Error with status code: expected %v, received %v", expectedStatus, res.Code)
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
		t.Errorf("Error reading response body")
		return
	}
	
	err = json.Unmarshal(payload, &post)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload : %v\n\t err : %v\n", payload, err)
		return
	}
	
	if post.OpportunityId != expectedOpportunityId {
		t.Errorf("Error with the opportunity id: expected %d, received %d\n", expectedOpportunityId, post.OpportunityId)
		return
	}
}

func TestGetOpportunityPosts(t *testing.T) {
	
	var (
		url      = "/posts/opportunity/1"
		req, err = http.NewRequest(http.MethodGet, url, nil)
		res      = httptest.NewRecorder()
		server   = NewServer()
	)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}
	server.ServeHTTP(res, req)

	var (
		expectedStatus                 = http.StatusOK
		posts                           = []models.Post{}
		expectedOpportunityId uint64 = 1
	)
	
	if res.Code != expectedStatus {
		t.Errorf("Error with status code: expected %v, received %v", expectedStatus, res.Code)
		return
	}
	
	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading response body")
		return
	}
	
	err = json.Unmarshal(payload, &posts)
	if err != nil {
		t.Errorf("Error unmarshalling payload: \n\t payload : %v\n\t err : %v\n", payload, err)
		return
	}
	
	if len(posts) == 0 {
		t.Error("Didn't get any post!\n")
		return
	}
	
	post := posts[0]
	
	if post.OpportunityId != expectedOpportunityId {
		t.Errorf("Error with the opportunity id: expected %d, received %d\n", expectedOpportunityId, post.OpportunityId)
		return
	}
	
	// fmt.Println(string(payload))
}

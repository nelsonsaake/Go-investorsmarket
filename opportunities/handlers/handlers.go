package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"projects/investorsmarket/opportunities/repo"
	"strconv"
)

type createOpportunityRequest struct {
	Name        string
	Amount      float64
	Industry    string
	Description string
	UserId      uint64
	Picture     string
	Returns     float64
	Duration    float32
	Location    string
}

func (coReq createOpportunityRequest) isValid() bool {

	// this test might be expanded on later
	// that is why the statements are broken down
	if len(coReq.Name) == 0 {
		fmt.Println("Problem with the create opportunity request, the name is empty!")
		return false
	}

	if coReq.Returns <= 0.0 {
		fmt.Println("Problem with the create opportunity request, the returns from the is not acceptable! returns=", coReq.Returns)
		return false
	}

	if coReq.Duration <= 0.0 {
		fmt.Println("Problem with the create opportunity request, the duration to yield returns is not acceptable! duration=", coReq.Duration)
		return false
	}

	return true
}

func CreateOpportunity(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading request body: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		var coReq createOpportunityRequest
		err = json.Unmarshal(payload, &coReq)
		if err != nil {
			fmt.Println("Error parsing request body: ", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		if !coReq.isValid() {
			fmt.Println("Error, create opportunity request is invalid.")
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		opportunity, err := repo.CreateOpportunity(coReq.Name, coReq.Amount, coReq.Industry, coReq.Description, coReq.UserId, coReq.Picture, coReq.Returns, coReq.Duration, coReq.Location)
		if err != nil {
			fmt.Println("Error creating a new opportunity: ", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		w.Header().Add("Location", "/opportunities/"+fmt.Sprint(opportunity.ID))

		formatter.JSON(w, http.StatusCreated, opportunity)
	}
}

func GetAllOpportunities(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		opportunities, err := repo.GetAllOpportunities()
		if err != nil {
			fmt.Println("Error getting a list of all opportunities: ", err)
			formatter.Text(w, http.StatusNotFound, "")
			return
		}
		formatter.JSON(w, http.StatusOK, opportunities)
	}
}

func GetOpportunityCreators(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		creators, err := repo.GetOpportunityCreators()
		if err != nil {
			fmt.Println("Error getting creators: ", err)
			formatter.Text(w, http.StatusNotFound, "")
			return
		}
		formatter.JSON(w, http.StatusOK, creators)
	}
}

func GetCreatorHistory(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		idstr := vars["id"]

		id, err := strconv.ParseUint(idstr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing id string: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		opportunities, err := repo.GetCreatorHistory(id)
		if err != nil {
			fmt.Printf("Error getting creators history: %v", err)
			return
		}
		formatter.JSON(w, http.StatusOK, opportunities)
	}
}

func GetOpportunity(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		idstr := vars["id"]

		id, err := strconv.ParseUint(idstr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing id string: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		opportunity, err := repo.GetOpportunity(id)
		if err != nil {
			fmt.Printf("Error getting opportunity: %v", err)
			return
		}
		formatter.JSON(w, http.StatusOK, opportunity)
	}
}
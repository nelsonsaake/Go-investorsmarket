package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"projects/investorsmarket/investments/repo"
	"strconv"
)

func GetInvestmentsByInvestor(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		idstr := vars["id"]

		id, err := strconv.ParseUint(idstr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing id string: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		investments, err := repo.GetInvestmentsByInvestor(id)
		if err != nil {
			fmt.Printf("Error getting investments made by an investor: %v", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		formatter.JSON(w, http.StatusOK, investments)
	}
}

type createInvestmentRequest struct {
	UserId        uint64
	OpportunityId uint64
	AmountBought  float64
}

func (ciReq createInvestmentRequest) isValid() bool {

	// can't think of any validation that won't be done better by repo
	// but as we add more stuff ...
	return true
}

func CreateInvestment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		var ciReq createInvestmentRequest
		err = json.Unmarshal(payload, &ciReq)
		if err != nil {
			fmt.Printf("Error parsing request body: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		if !ciReq.isValid() {
			fmt.Println("Error: request is in valid")
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		investment, err := repo.CreateInvestment(ciReq.UserId, ciReq.OpportunityId, ciReq.AmountBought)
		if err != nil {
			fmt.Printf("Error creating investment: %v", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		w.Header().Add("Location", "/investments/"+fmt.Sprint(investment.ID))

		formatter.JSON(w, http.StatusCreated, investment)
	}
}

func GetInvestment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		idstr := vars["id"]

		id, err := strconv.ParseUint(idstr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing id string: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		investment, err := repo.GetInvestment(id)
		if err != nil {
			fmt.Printf("Error getting investment: %v", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		formatter.JSON(w, http.StatusOK, investment)
	}
}

func GetAllInvestments(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		investments, err := repo.GetAllInvestments()
		if err != nil {
			fmt.Printf("Error getting all investments: %v", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		formatter.JSON(w, http.StatusOK, investments)
	}
}
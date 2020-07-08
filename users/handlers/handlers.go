package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"projects/investorsmarket/users/repo"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"strconv"
)

type retrieveUserRequest struct {
	Email    string
	Password string
}

func (retrieveUserRequest retrieveUserRequest) isValid() bool {

	// email and password must not be empty
	return !(len(retrieveUserRequest.Email) == 0 || len(retrieveUserRequest.Password) == 0)
}

func GetUserA(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Error reading request body: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		var userReq retrieveUserRequest
		err = json.Unmarshal(payload, &userReq)
		if err != nil {
			fmt.Printf("Error unmarshalling json: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		if userReq.isValid() == false {
			fmt.Println("Error: invalid user")
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		user, err := repo.GetUserA(userReq.Email, userReq.Password)
		if err != nil {
			fmt.Printf("Error getting user from repos: %v", err)
			formatter.Text(w, http.StatusNotFound, "")
			return
		}

		formatter.JSON(w, http.StatusOK, user)
	}
}

type createUserRequest struct {
	Email    string
	Password string
}

func (cuReq createUserRequest) isValid() bool {

	// these assertions are intensionally spaced out;
	// just incase we want to do more test on each of them
	if len(cuReq.Email) == 0 {
		return false
	}

	if len(cuReq.Password) == 0 {
		return false
	}

	return true
}

func CreateUser(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Error reading request body: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		var cuReq createUserRequest
		err = json.Unmarshal(payload, &cuReq)
		if err != nil {
			fmt.Printf("Error parsing payload: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		if !cuReq.isValid() {
			fmt.Printf("Error user information provided is in valid!")
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		user, err := repo.CreateUser(cuReq.Email, cuReq.Password)
		if err != nil {
			fmt.Printf("Error: %v", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}
		w.Header().Add("Location", "/users/"+fmt.Sprint(user.ID))
		formatter.JSON(w, http.StatusCreated, user)
	}
}

func GetAllUsers(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		users := repo.GetAllUsers()
		formatter.JSON(w, http.StatusOK, users)
	}
}

func GetUser(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		idstr := vars["id"]

		id, err := strconv.ParseUint(idstr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing id string: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		user, err := repo.GetUser(id)
		if err != nil {
			fmt.Printf("Error getting user profile: %v", err)
			formatter.Text(w, http.StatusNotFound, "")
			return
		}
		formatter.JSON(w, http.StatusOK, user)
	}
}

type updateUserRequest struct {
	Email          string 
	Picture        string
	FirstName      string
	Surname        string
	DateOfBirth    string
	Gender         string
	PhoneNumber    string
	Nationality    string
	Occupation     string
	Address        string
	Country        string
	Region         string
	City           string
	AccName        string
	AccNumber      string
	AccBankName    string
	NkSurname      string
	NkFirstName    string
	NkRelationship string
	NkEmail        string
	NkPhoneNumber  string
	NkAddress      string
}

// ? check date of birth
// ? it ends with a default value instead of null
func UpdateUser(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		idstr := vars["id"]

		id, err := strconv.ParseUint(idstr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing id string: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}


		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Error reading request body: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		var uuReq updateUserRequest
		err = json.Unmarshal(payload, &uuReq)
		if err != nil {
			fmt.Printf("Error parsing payload: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}		

		user, err := repo.UpdateUser(id, uuReq.Email, uuReq.Picture, uuReq.FirstName, uuReq.Surname, uuReq.DateOfBirth, uuReq.Gender, uuReq.PhoneNumber, uuReq.Nationality, uuReq.Occupation, uuReq.Address, uuReq.Country, uuReq.Region, uuReq.City, uuReq.AccName, uuReq.AccNumber, uuReq.AccBankName, uuReq.NkSurname, uuReq.NkFirstName, uuReq.NkRelationship, uuReq.NkEmail, uuReq.NkPhoneNumber, uuReq.NkAddress)
		if err != nil {
			fmt.Printf("Error updating user profile: %v", err)
			formatter.Text(w, http.StatusNotFound, "")
			return
		}
		formatter.JSON(w, http.StatusOK, user)
	}
}

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hako/branca"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"projects/investorsmarket/changepasswords/repo"
	"time"
)

type createChangePasswordRequest struct {
	UserId uint64
}

func getBranca() *branca.Branca {
	// ? remember to set branca key, 32 characters
	// we removed the key so that it won't appear in the github commit
	// set key as env 
	key := ""
	return branca.NewBranca(key)
}

func newCode() (code string, err error) {

	brca := getBranca()
	code = time.Now().Format(time.RFC3339)
	code, err = brca.EncodeToString(string(code[:]))
	return
}

func CreateChangePassword(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		cpReq := createChangePasswordRequest{}
		err = json.Unmarshal(payload, &cpReq)
		if err != nil {
			fmt.Printf("Error parsing payload body: payload %v,\n err %v \n ", string(payload), err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		code, err := newCode()
		if err != nil {
			fmt.Printf("Error generating code: %v \n", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		cp, err := repo.CreateChangePassword(cpReq.UserId, code)
		if err != nil {
			fmt.Printf("Error creating change password: %v\n", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		w.Header().Add("Location", "/changepasswords/"+fmt.Sprint(cp.ID))

		formatter.JSON(w, http.StatusCreated, cp)
	}
}

type getChangePasswordRequest struct {
	UserId uint64
	Code   string
}

func GetChangePassword(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		gcpReq := getChangePasswordRequest{}
		err = json.Unmarshal(payload, &gcpReq)
		if err != nil {
			fmt.Printf("Error parsing payload body: payload %v,\n err %v \n ", string(payload), err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		cp, err := repo.GetChangePassword(gcpReq.UserId, gcpReq.Code)
		if err != nil {
			fmt.Printf("Error getting change password record: %v\n", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		formatter.JSON(w, http.StatusOK, cp)
	}
}

func GetAllChangePasswords(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cps, err := repo.GetAllChangePasswords()
		if err != nil {
			fmt.Printf("Error getting all change password records: %v\n", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		formatter.JSON(w, http.StatusOK, cps)
	}
}

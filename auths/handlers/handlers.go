package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"projects/investorsmarket/auths/repo"
	"time"

	"github.com/hako/branca"
	"github.com/unrolled/render"
)

func getBranca() *branca.Branca {
	// ? remember to set branca key, 32 characters
	// we removed the key so that it won't appear in the github commit
	// set key as env 
	key := ""
	return branca.NewBranca(key)
}

func newCode(email, password string) (code string, err error) {

	brca := getBranca()
	code = time.Now().Format(time.RFC3339)
	code = email + code + password
	code, err = brca.EncodeToString(string(code[:]))
	return
}

type createAuthRequest struct {
	Email    string
	Password string
}

func (caReq createAuthRequest) isValid() bool {

	if len(caReq.Email) < 5 {
		fmt.Println("Invalid email: too short")
		return false
	}

	if len(caReq.Password) < 8 {
		fmt.Println("Invalid password: too short")
		return false
	}

	return true
}

type authResponse struct {
	UserId uint64
	Token  string
	Role   string
}

func CreateAuth(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		caReq := createAuthRequest{}
		err = json.Unmarshal(payload, &caReq)
		if err != nil {
			fmt.Printf("Error parsing payload body: payload %v,\n err %v \n ", string(payload), err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		if !caReq.isValid() {
			err = fmt.Errorf("Error request is invaid: payload %v,\n err %v \n ", string(payload), err)
			fmt.Println(err)
			formatter.Text(w, http.StatusBadRequest, err.Error())
			return
		}

		code, err := newCode(caReq.Email, caReq.Password)
		if err != nil {
			fmt.Printf("Error generating code: %v \n", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		auth, err := repo.CreateAuth(caReq.Email, caReq.Password, code, "user")
		if err != nil {
			fmt.Printf("Error creating auth: %v\n", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		authRes := authResponse{
			UserId: auth.UserId,
			Token:  auth.Token,
			Role:   auth.Role,
		}

		formatter.JSON(w, http.StatusCreated, authRes)
	}
}

type getAuthRequestGivenEP struct {
	Email    string
	Password string
}

func (gaReq getAuthRequestGivenEP) isValid() bool {

	if len(gaReq.Email) < 5 {
		fmt.Println("Invalid email: too short")
		return false
	}

	if len(gaReq.Password) < 8 {
		fmt.Println("Invalid password: too short")
		return false
	}

	return true
}

func GetAuthGivenEP(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		gaReq := getAuthRequestGivenEP{}
		err = json.Unmarshal(payload, &gaReq)
		if err != nil {
			fmt.Printf("Error parsing payload body: payload %v,\n err %v \n ", string(payload), err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		if !gaReq.isValid() {
			err = fmt.Errorf("Error request is invaid: payload %v,\n err %v \n ", string(payload), err)
			formatter.Text(w, http.StatusBadRequest, err.Error())
			return
		}

		auth, err := repo.GetAuthGivenEP(gaReq.Email, gaReq.Password)
		if err != nil {
			fmt.Printf("Error getting auth given email: %s, and password: %s, \nerr: %v\n", gaReq.Email, gaReq.Password, err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		authRes := authResponse{
			UserId: auth.UserId,
			Token:  auth.Token,
			Role:   auth.Role,
		}

		formatter.JSON(w, http.StatusOK, authRes)
	}
}

type getAuthGivenTokenRequest struct {
	Token string
}

func GetAuthGivenToken(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		gaReq := getAuthGivenTokenRequest{}
		err = json.Unmarshal(payload, &gaReq)
		if err != nil {
			fmt.Printf("Error parsing payload body: payload %v,\n err %v \n ", string(payload), err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		auth, err := repo.GetAuthGivenToken(gaReq.Token)
		if err != nil {
			fmt.Printf("Error getting auth given token: %s \nerr: %v\n", gaReq.Token, err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		authRes := authResponse{
			UserId: auth.UserId,
			Token:  auth.Token,
			Role:   auth.Role,
		}

		formatter.JSON(w, http.StatusOK, authRes)
	}
}

type getAuthGivenUserAndTokenRequest struct {
	UserId uint64
	Token  string
}

func GetAuthGivenUserAndToken(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		gaReq := getAuthGivenUserAndTokenRequest{}
		err = json.Unmarshal(payload, &gaReq)
		if err != nil {
			fmt.Printf("Error parsing payload body: payload %v,\n err %v \n ", string(payload), err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		auth, err := repo.GetAuthGivenUserAndToken(gaReq.UserId, gaReq.Token)
		if err != nil {
			fmt.Printf("Error getting auth given token: %s \nerr: %v\n", gaReq.Token, err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		authRes := authResponse{
			UserId: auth.UserId,
			Token:  auth.Token,
			Role:   auth.Role,
		}

		formatter.JSON(w, http.StatusOK, authRes)
	}
}

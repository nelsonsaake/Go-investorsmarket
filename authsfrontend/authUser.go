package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Auth struct {
	UserId uint64
	Token  string
	Role   string
}

func extractAuthFromResponse(res *http.Response) (auth Auth, err error) {

	// interpret
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Error with the response code: expected %v, received %v ", http.StatusOK, res.StatusCode)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("Error reading the response body: %v\n", err)
		return
	}
	defer res.Body.Close()

	err = json.Unmarshal(payload, &auth)
	if err != nil {
		err = fmt.Errorf("Error unmarshalling payload: \n\t payload: %s \n\t err: %v", string(payload), err)
		return
	}

	return
}

func getAuthGivenToken(token string) (auth Auth, err error) {

	// create request
	var (
		method  = http.MethodPost
		url     = "http://localhost:3007/auths/t1"
		reqBody = []byte(`{"token": "` + token + `" }`)
		client  = http.Client{}
	)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		err = fmt.Errorf("Error creating request: %v", err)
		return
	}

	// send request
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("Error creating request: %v", err)
		return
	}

	auth, err = extractAuthFromResponse(res)
	return
}

func grant(
	formatter *render.Render,
	next func(*render.Render) http.HandlerFunc,
	errorpage func(http.ResponseWriter, *http.Request, string),
	role string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// read token in header
		var (
			tokenKey = "Token"
			token    = r.Header.Get(tokenKey)
		)

		auth, err := getAuthGivenToken(token)
		if err != nil {
			err = fmt.Errorf("failed to user with given token: %v", err)
			errorpage(w, r, err.Error())
			return
		}

		// make sure it is type of user
		if auth.Role != role {
			err = fmt.Errorf("unauthorized")
			errorpage(w, r, err.Error())
			return
		}

		// add token to response
		w.Header().Add(tokenKey, token)

		// move the the apporiate handler
		next(formatter)(w, r)
	}
}

func GrantUser(
	formatter *render.Render,
	next func(*render.Render) http.HandlerFunc,
	errorpage func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		grant(formatter, next, errorpage, "user")(w, r)
	}
}

func GrantAdmin(
	formatter *render.Render,
	next func(*render.Render) http.HandlerFunc,
	errorpage func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		grant(formatter, next, errorpage, "admin")(w, r)
	}
}

func grantAfterMatch(
	formatter *render.Render,
	next func(*render.Render) http.HandlerFunc,
	errorpage func(http.ResponseWriter, *http.Request, string),
	role string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// read token in header
		var (
			tokenKey = "Token"
			token    = r.Header.Get(tokenKey)
		)

		auth, err := getAuthGivenToken(token)
		if err != nil {
			err = fmt.Errorf("failed to user with given token: %v", err)
			errorpage(w, r, err.Error())
			return
		}

		// make sure it is type of user
		if auth.Role != role {
			err = fmt.Errorf("unauthorized")
			errorpage(w, r, err.Error())
			return
		}

		// make sure the id matched
		var (
			vars   = mux.Vars(r)
			idstr  = vars["id"]
			authId = fmt.Sprint(auth.UserId)
		)
		if authId != idstr {
			err = fmt.Errorf("unauthorized, id does not match user")
			errorpage(w, r, err.Error())
			return
		}

		// add token to response
		w.Header().Add(tokenKey, token)

		// move the the apporiate handler
		next(formatter)(w, r)
	}
}

func IsMatch(
	formatter *render.Render,
	next func(*render.Render) http.HandlerFunc,
	errorpage func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		grantAfterMatch(formatter, next, errorpage, "user")(w, r)
	}
}

func GetAuthPackage(r *http.Request) (auth Auth, err error) {

	var (
		tokenKey = "Token"
		token    = r.Header.Get(tokenKey)
	)
	auth, err = getAuthGivenToken(token)
	if err != nil {
		err = fmt.Errorf("failed to user with given token: %v", err)
		return
	}
	return
}

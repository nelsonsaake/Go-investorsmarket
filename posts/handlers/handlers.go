package handler

import (
	"encoding/json"
	"fmt"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
	"projects/investorsmarket/posts/repo"
	"github.com/gorilla/mux"
	"strconv"
)

type createPostRequest struct{
	Picture string
	Description string
	OpportunityId uint64
}

func CreatePost(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body: ", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		cpReq := createPostRequest{}
		err = json.Unmarshal(payload, &cpReq)
		if err != nil {
			fmt.Printf("Error unmarshalling payload: payload %v,\n err %v \n ", string(payload), err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		post, err := repo.CreatePost(cpReq.Picture, cpReq.Description, cpReq.OpportunityId)
		if err != nil {
			fmt.Printf("Error creating new post: %v\n", err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		w.Header().Add("Location", "/posts/"+fmt.Sprint(post.ID))

		formatter.JSON(w, http.StatusCreated, post)
	}
}

// posts are done to show the progress of a business. Such business is an opportunity as it has offering to allow investors to participate. 
// this handler gets all posts related to a particulart opportunity and returns it
func GetOpportunityPosts(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		idstr := vars["id"]

		id, err := strconv.ParseUint(idstr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing id string: %v", err)
			formatter.Text(w, http.StatusBadRequest, "")
			return
		}

		posts, err := repo.GetOpportunityPosts(id)
		if err != nil {
			fmt.Printf("Error getting all posts related to opportunity with id: %d, \n err : %v\n", id, err)
			formatter.Text(w, http.StatusExpectationFailed, "")
			return
		}

		formatter.JSON(w, http.StatusOK, posts)
	}
}
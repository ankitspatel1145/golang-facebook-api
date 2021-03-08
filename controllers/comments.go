package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/faceapi/models"
	"github.com/gorilla/mux"
)

// GetComment will return a single comment that is requested from the params
func GetComment(input http.ResponseWriter, r *http.Request) {
	input.Header().Set("Content-Type", "application/json")
	var comment models.Comment // make blank comment object
	params := mux.Vars(r)
	// fmt.Println(params)
	models.DB.First(&comment, params["id"]) //extract comment that was asked for
	if comment.ID == 0 {
		json.NewEncoder(input).Encode("no comment found")

	} else {

		json.NewEncoder(input).Encode(&comment)
	}
}

// CreateComment will let you add a comment to the DB and relate it to a post
func CreateComment(input http.ResponseWriter, r *http.Request) {
	input.Header().Set("Content-Type", "application/json")
	var newComment models.Comment
	params := mux.Vars(r)

	//this takes params and fills in our newComment object
	_ = json.NewDecoder(r.Body).Decode(&newComment)

	// this chunk of code take the id from params and turns it into the right format
	postStr := params["id"]
	post64, _ := strconv.ParseUint(postStr, 10, 32)
	postID := uint(post64)
	newComment.PostID = postID

	// we get the post from the DB
	var post models.Post
	models.DB.First(&post, postID)

	if post.ID == 0 {
		json.NewEncoder(input).Encode("no post found")
	} else {
		//if parent post exists, then we create new comment in the db and set up relationship
		models.DB.Create(&newComment)
		models.DB.Model(&post).Association("Comment").Append(newComment)
		// fmt.Println(newComment.PostID)
		json.NewEncoder(input).Encode(newComment)
	}
}


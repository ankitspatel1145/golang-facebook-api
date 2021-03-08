package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/faceapi/models"
	"github.com/gorilla/mux"
)

// TestServer - define functions here that will be called when we hit an endpoint (just the /test endpoint)
func TestServer(input http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(input).Encode("This is a test route to make sure the server is running")
}

// GetAllPosts will be used for the index page, will return all posts
func GetAllPosts(input http.ResponseWriter, r *http.Request){
	input.Header().Set("Content-Type", "application/json")
	var posts []models.Post
	models.DB.Find(&posts)
	json.NewEncoder(input).Encode(posts)
}

// GetPost will return a single post that is requested from the params
func GetPost(input http.ResponseWriter, r *http.Request){
	input.Header().Set("Content-Type", "application/json")
	var post models.Post // make blank post object
	params := mux.Vars(r)
	// fmt.Println(params)
	models.DB.First(&post, params["id"]) //extract post that was asked for
	if post.ID == 0 {
		json.NewEncoder(input).Encode("no post found")
		
	} else {
		//if the post exists, then all of the comments will also be fetched from DB
		var postComments []models.Comment
		models.DB.Model(&post).Association("Comments").Find(&postComments)
		// fmt.Println(postComments)
		// fmt.Println(post)
		post.Comments = postComments
		json.NewEncoder(input).Encode(&post)
	}
}


// CreatePost will add a new post to our DB
func CreatePost(input http.ResponseWriter, r *http.Request){
	input.Header().Set("Content-Type", "application/json")
	var newPost models.Post
	//this takes params and fills in our newPost object
	_ = json.NewDecoder(r.Body).Decode(&newPost) 

	models.DB.Create(&newPost)
	json.NewEncoder(input).Encode(newPost)
}

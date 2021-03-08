package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/faceapi/models"
	"github.com/gorilla/mux"
)



// GetComment will return a single comment that is requested from the params
func GetComment(input http.ResponseWriter, r *http.Request){
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
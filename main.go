package main

//TODO:  "add a comment to a post"
import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/faceapi/models"
	"github.com/faceapi/controllers"
	
)



func main() {
	
	//make the new router to handle requests
	router := mux.NewRouter()

	//when the serer is started, it automatically connects to the DB
	models.ConnectDB()

	models.SeedDB()
	//define our endpoints and functions to run when we hit those endpoints

	router.HandleFunc("/test", controllers.TestServer).Methods("GET")
	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/post/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")

	router.HandleFunc("/comment/{id}", controllers.GetComment).Methods("GET")
	router.HandleFunc("/post/{id}/comment", controllers.CreateComment).Methods("POST")
	fmt.Println("server is running")
	// start server
	log.Fatal(http.ListenAndServe(":3000", router))


}
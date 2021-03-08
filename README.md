# Golang REST API

This is a simple Golang backend that I made to learn Golang
## To Run/Test
from the root of the project run these commands:

 ```go build```
 
 ```.\faceapi.exe```
 
 use postman to interact with the endpoints on 
 
 http://localhost:3000/

## Endpoints 

* ```/test``` - "GET"
* ```/posts``` - "GET", "POST"
* ```/posts/{id}``` - "GET"
* ```/comment/{id}``` - "GET"
* ```/post/{id}/comment``` - "POST"

### Making a POST request
When making a POST request to the server:
1. Use Postman
2. In the headers tab, set ```Content-Type``` to ```application/json```
3. Pass in JSON for the payload under Body > raw


## Technical Specs
It uses Gorilla mux for its routing

https://github.com/gorilla/mux

It also uses GORM as an Object relational layer on the SQLite data base

https://gorm.io/index.html

and SQLite3
## Models

There are 2 classes/models:

```
Posts
    Id
    Body

Comments
    Id
    PostId
    body
```
## You can do the following:
- Create a post
- View a specific post (along with all comments of that post)
- Create a comment (must be on an exisiting parent post)
- View a specific post

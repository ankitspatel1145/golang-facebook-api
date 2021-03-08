This is a simple Golang backend that I made to learn Golang

It uses Gorilla mux for its routing
https://github.com/gorilla/mux

It also uses GORM as an Object relational layer on the SQLite data base
https://gorm.io/index.html

There are 2 classes/models:
Posts
    Id
    Body

Comments
    Id
    PostId
    body


You have can do the following:
- Create a post
- View a specific post (along with all comments of that post)
- Create a comment (must be on an exisiting parent post)
- View a specific post

package models

import (
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

// Post model  has id, body, and comments relationship  these will be the posts that show up 
type Post struct {
	gorm.Model
	// id field is not needed if you dont put index on it, else it throws warning
	ID    uint   ` gorm:"index"`
	Body  string  `json:"body"`
	Comments []Comment
}

// Comment model hash id, body, and a post that it belongs to
type Comment struct {
	gorm.Model
	ID      uint  `gorm:"index"`
	PostID  uint `json:"postid"`
	Body    string `json:"body"`
	// Post	  Post
}

// DB is going to be our data base
var DB *gorm.DB

// ConnectDB will open a connection the a sqlite3 db that we define
func ConnectDB() {
  database, err := gorm.Open(sqlite.Open("facebookapi.db"), &gorm.Config{})

  if err != nil {
    panic("Failed to connect to database!")
  }

  // this line auto migrates the tables, makes whatever tables it needs to
  database.AutoMigrate(&Post{}, &Comment{}) 
  DB = database
}

// SeedDB will seed the DB if it is empty
func SeedDB() {
	var posts []Post
	DB.Find(&posts)
	if len(posts) == 0 {
		DB.Create(&Post{Body: "does anyone know how to set up a golang backend?"})
		DB.Create(&Post{Body: "This is totally the new FB!"})
    	DB.Create(&Comment{Body: "FIRST!", PostID: 1})
    	DB.Create(&Comment{Body: "no clue man", PostID: 1})
	}
	
}
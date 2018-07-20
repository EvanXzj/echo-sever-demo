package models

import (
	"fmt"

	"github.com/evanxzj/echo-server/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // for gorm
)

// User : model for app user
type User struct {
	gorm.Model
	Username string    `gorm:"text;not null;unique" json:"username"`
	First    string    `gorm:"text" json:"first"`
	Last     string    `gorm:"text" json:"last"`
	Password string    `gorm:"text" json:"-"`
	Sessions []Session `gorm:"ForeignKey:UserID" json:"sessions"`
}

// UsersAPI : used for interacting with user models
type UsersAPI struct {
	db *gorm.DB
}

var crypt = utils.NewCrypt()

// Create : create a user in db
func (u UsersAPI) Create(username, first, last, pwd string) {
	pw, _ := crypt.Hash(pwd)

	u.db.Create(&User{
		Username: username,
		First:    first,
		Last:     last,
		Password: pw,
	})

}

// GetUserByID : get a user by its id
func (u UsersAPI) GetUserByID(id int) User {
	var user User

	u.db.Find(&user, id)

	return user
}

// GetAll : get all users in db
func (u UsersAPI) GetAll() []User {
	var users []User

	u.db.Find(&users)

	return users
}

// Login : attempt to login a user or error if login failed
func (u UsersAPI) Login(username, pw string) (User, errror) {
	var user User

	// look up user
	u.db.Where("username = ?", username).Find(&user)

	if user.Username != username {
		return user, fmt.Errorf("No user found by that username")
	}

	if crypt.Check(user.Password, pw) {
		return user, nil
	}

	return user, fmt.Errorf("Invalid login")
}

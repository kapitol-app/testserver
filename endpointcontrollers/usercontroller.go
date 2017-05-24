package endpointcontrollers

import (
	"github.com/greenac/testserver/endpoints"
	"net/http"
	"github.com/greenac/testserver/models"
	"encoding/json"
	"fmt"
)

type User struct {
	UserId int `json:"userId"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastname"`
	Dob string `json:"dob"`
	Age int `json:"age"`
	IsActive bool `json:"isActive"`
}

type users struct {
	allUsers map[int]*models.User
}

func (users *users)initialize() {
	if users.allUsers == nil {
		users.allUsers = make(map[int]*models.User)
	}
}

func (users *users)addUser(u *models.User) {
	fmt.Println("saving user:", *u)
	if u.UserId == 0 {
		fmt.Println("all users:", users.allUsers)
		u.UserId = len(users.allUsers) + 1
		users.allUsers[u.UserId] = u
	} else {
		_, hasUser := users.allUsers[u.UserId]
		if !hasUser {
			users.allUsers[u.UserId] = u
		}
	}
}

type UserController struct {
	userEndPoints *endpoints.UserEndpoints
	users *users
}

func (uc *UserController)Initialize() {
	if uc.userEndPoints == nil {
		ep := endpoints.UserEndpoints{}
		uc.userEndPoints = &ep
	}

	if uc.users == nil {
		urs := users{}
		urs.initialize()
		uc.users = &urs
	}
}

func (uc *UserController)SaveUser(w http.ResponseWriter, req *http.Request) {
	var u models.User
	err := json.NewDecoder(req.Body).Decode(&u)
	if err == nil {
		fmt.Println("Saved new user:", &u)
		uc.users.addUser(&u)
		endpoints.SaveUser(w, &u)
	} else {
		fmt.Println("Error saving new user:", err)
		endpoints.SaveUserFailed(w, &u)
	}
}

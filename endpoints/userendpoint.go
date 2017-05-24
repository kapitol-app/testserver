package endpoints

import (
	"net/http"
	"fmt"
	"github.com/greenac/testserver/models"
	"github.com/greenac/testserver/restresponse"
)

type UserEndpoints struct{}

func (ue *UserEndpoints)SaveUser() string {
	return "/save-user"
}

func SaveUserFailed(w http.ResponseWriter, u *models.User) {
	fmt.Println("Sending save user failure to client.")
	res := restresponse.Response{
		Code: restresponse.InternalServerError,
		Payload: map[string]interface{}{"user": u},
	}
	res.Respond(w)
}

func SaveUser(w http.ResponseWriter, u *models.User) {
	fmt.Println("Saving new user:", u)
	res := restresponse.Response{
		Code: restresponse.OK,
		Payload: map[string]interface{}{"data": []float64{45, 34, 98, 90.1}},
	}
	res.Respond(w)
}

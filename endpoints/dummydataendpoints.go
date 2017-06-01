package endpoints

import (
	"net/http"
	"github.com/kapitol-app/testserver/restresponse"
	"github.com/kapitol-app/testserver/models"
	"fmt"
)

type DummyDataEndpoints struct {}

func (ddep *DummyDataEndpoints)GetSenatorUri() string {
	return "/dummy-data/get-senators"
}

func (ddep *DummyDataEndpoints)GetSenators(w http.ResponseWriter, senators *[]models.Senator) {
	sens := make([]map[string]interface{}, 0, len(*senators))
	fmt.Println("senators:", sens)
	for _, senator := range *senators {
		sens = append(sens, senator.ToMap())
	}

	payload := map[string]interface{} {"senators": sens}
	res := restresponse.Response{Code: restresponse.OK, Payload: payload}
	res.Respond(w)
}

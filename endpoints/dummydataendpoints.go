package endpoints

import (
	"net/http"
	"github.com/greenac/testserver/restresponse"
	"github.com/greenac/testserver/models"
	"fmt"
)

type DummyDataEndpoints struct {}

func (ddep *DummyDataEndpoints)GetSenatorUri() string {
	return "/dummy-data/get-senators"
}

func (ddep *DummyDataEndpoints)GetSenators(w http.ResponseWriter, senators *[]models.Senator) {
	ms := make([]map[string]interface{}, 0, len(*senators))
	fmt.Println("senators:", ms)
	for i, senator := range *senators {
		ms = append(ms, senator.ToMap())
		fmt.Println("i:", i, "ms length:", len(ms))
	}

	payload := map[string]interface{} {"senators": ms}
	res := restresponse.Response{Code: restresponse.OK, Payload: payload}
	res.Respond(w)
}

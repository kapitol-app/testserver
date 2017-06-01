package endpointcontrollers

import (
	"github.com/kapitol-app/testserver/endpoints"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/kapitol-app/testserver/models"
	"io/ioutil"
	"path/filepath"
)

type DummyDataController struct {
	EndPoints *endpoints.DummyDataEndpoints
}

func (ddc *DummyDataController)Initialize() {
	if ddc.EndPoints == nil {
		ddep := endpoints.DummyDataEndpoints{}
		ddc.EndPoints = &ddep
	}
}

func (ddc *DummyDataController)GetSenators(w http.ResponseWriter, req *http.Request) {
	fp, err := filepath.Abs("files/senator.json")
	if err != nil {
		fmt.Println("Error: could not abs file path file at error:", err)
		// handle error here
		return
	}

	senJsonData, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println("Error: could not read senator json file at:", fp, "error:", err)
		// handle error here
		return
	}

	var senators []models.Senator
	err = json.Unmarshal(senJsonData, &senators)
	if err != nil {
		fmt.Println("Error: could not marshal senator json. Error:", err)
		// handle error here
		return
	}

	ddc.EndPoints.GetSenators(w, &senators)
}

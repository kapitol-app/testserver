package endpointcontrollers

import (
	"github.com/kapitol-app/testserver/endpoints"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/kapitol-app/testserver/models"
	"io/ioutil"
	"os"
	"log"
	"path"
)

const senatorJsonFilePath = "files/senator.json"

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
	ex, err := os.Executable()
	if err != nil {
		log.Fatal("Error: could not get current senator executalbe file path. Error:", err)
		panic(err)
	}

	senJsonData, err := ioutil.ReadFile(path.Join(path.Dir(ex), senatorJsonFilePath))
	if err != nil {
		fmt.Println("Error: could not read senator json file at:", senatorJsonFilePath, "error:", err)
		//handle error here
		return
	}

	var senators []models.Senator
	err = json.Unmarshal(senJsonData, &senators)
	if err != nil {
		fmt.Println("Error: could not marshal senator json. Error:", err)
		return
	}

	ddc.EndPoints.GetSenators(w, &senators)
}

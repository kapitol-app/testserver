package restresponse

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Payload map[string]interface{} `json:"payload"`
}

func (r *Response)Respond(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("headers:", w.Header())
	fmt.Println("response:", &r)
	json.NewEncoder(w).Encode(r)
}

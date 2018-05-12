package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

//JSONResp ...
type JSONResp struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"payload,omitempty"`
}

//NewJSONResp ...
func NewJSONResp(status string, code int, data interface{}) *JSONResp {
	return &JSONResp{Status: status, Code: code, Data: data}
}

//respondWithJSON responds with json encoded data
func respondWithJSON(w http.ResponseWriter, code int, stat string, payload interface{}) {
	resp := NewJSONResp(stat, code, payload)
	jsResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(jsResp)
}

//respondWithError respond with an error message
func respondWithError(w http.ResponseWriter, code int, stat, msg string) {
	respondWithJSON(w, code, stat, map[string]string{"error": msg})
}

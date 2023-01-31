package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type ResponseStruct struct {
	Authorization bool
}

func WriteAsJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	WriteAsJson(w, struct {
		Error string `json:"error"`
	}{Error: err.Error()})
}

func TransactionValidate() bool {

	URL := os.Getenv("MOCK_URL")
	response, _ := http.Get(URL)
	responseData, _ := ioutil.ReadAll(response.Body)
	simpleObject := ResponseStruct{}
	err := json.Unmarshal(responseData, &simpleObject)

	if err != nil {
		return false
	}

	return simpleObject.Authorization
}

package helper

import (
	"encoding/json"
	"golang-restful-api/model/web"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	errDecode := decoder.Decode(result)
	PanicIfError(errDecode)
}

func WriteToResponseBody(writer http.ResponseWriter, webResponse web.WebResponse) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(webResponse)
	PanicIfError(errEncode)
}
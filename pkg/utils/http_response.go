package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
)

//InternalServerError constant
const (
	InternalServerError = "Internal Server Error"
)

var responses = map[int]models.APIResponse{
	http.StatusOK:                  {StatusCode: http.StatusOK, Description: "OK"},
	http.StatusCreated:             {StatusCode: http.StatusCreated, Description: "Created"},
	http.StatusConflict:            {StatusCode: http.StatusConflict, Description: "Conflict"},
	http.StatusNotFound:            {StatusCode: http.StatusNotFound, Description: "Not Found"},
	http.StatusUnauthorized:        {StatusCode: http.StatusUnauthorized, Description: "Unauthorized Request"},
	http.StatusBadRequest:          {StatusCode: http.StatusBadRequest, Description: "Request has some field errors"},
	http.StatusInternalServerError: {StatusCode: http.StatusInternalServerError, Description: "Internal Server Error"},
	http.StatusUnprocessableEntity: {StatusCode: http.StatusUnprocessableEntity, Description: "Unprocessable Request"},
}

//HTTPResponse type
type HTTPResponse struct {
	http.ResponseWriter
	ctx context.Context
}

//WriteResponse outputs response
func (res *HTTPResponse) WriteResponse(status int, data interface{}) (int, error) {
	output, err := json.Marshal(data)
	if err != nil {
		status = http.StatusInternalServerError
		output = []byte(InternalServerError)
	}
	return writeResponse(res, "application/json", status, output)
}

//UnauthorizedResponse returns response with http status 401
func (res *HTTPResponse) UnauthorizedResponse() {
	_, _ = res.WriteResponse(http.StatusUnauthorized, responses[http.StatusUnauthorized])
}

//OKResponse returns response with http status 200
func (res *HTTPResponse) OKResponse(data interface{}) {
	apiRes := responses[http.StatusOK]
	apiRes.Data = data
	_, _ = res.WriteResponse(http.StatusOK, apiRes)
}

//CreatedResponse returns response with http status 201
func (res *HTTPResponse) CreatedResponse(data interface{}) {
	apiRes := responses[http.StatusCreated]
	apiRes.Data = data
	_, _ = res.WriteResponse(http.StatusCreated, apiRes)
}

//NotFoundResponse returns response with http status 404
func (res *HTTPResponse) NotFoundResponse(data interface{}) {
	apiRes := responses[http.StatusNotFound]
	apiRes.Data = data
	_, _ = res.WriteResponse(http.StatusNotFound, apiRes)
}

//UnprocessableResponse returns response with http status 422
func (res *HTTPResponse) UnprocessableResponse() {
	apiRes := responses[http.StatusUnprocessableEntity]
	apiRes.Errors = []models.Error{{
		Code:  http.StatusUnprocessableEntity,
		Error: "Json parse error",
	}}
	_, _ = res.WriteResponse(http.StatusUnprocessableEntity, apiRes)
}

//ServerErrorResponse returns response with http status 500
func (res *HTTPResponse) ServerErrorResponse(errors []models.Error) {
	apiRes := responses[http.StatusInternalServerError]
	apiRes.Errors = errors
	_, _ = res.WriteResponse(http.StatusInternalServerError, apiRes)
}

//BadRequestResponse returns response with http status 400
func (res *HTTPResponse) BadRequestResponse(errors []models.Error) {
	apiRes := responses[http.StatusBadRequest]
	apiRes.Errors = errors
	_, _ = res.WriteResponse(http.StatusBadRequest, apiRes)
}

//ConflictResponse returns response with http status 409
func (res *HTTPResponse) ConflictResponse(errors []models.Error) {
	apiRes := responses[http.StatusConflict]
	apiRes.Errors = errors
	_, _ = res.WriteResponse(http.StatusConflict, apiRes)
}

func writeResponse(res *HTTPResponse, contentType string, status int, data []byte) (int, error) {
	res.Header().Set("Content-Type", contentType)
	res.Header().Set("Content-Length", strconv.Itoa(len(data)))
	res.WriteHeader(status)
	return res.Write(data)
}

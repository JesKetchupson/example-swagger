// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "gen/models"
)

// GetTaskSyncOKCode is the HTTP code returned for type GetTaskSyncOK
const GetTaskSyncOKCode int = 200

/*GetTaskSyncOK Состояние задачи

swagger:response getTaskSyncOK
*/
type GetTaskSyncOK struct {

	/*
	  In: Body
	*/
	Payload *models.TaskStatus `json:"body,omitempty"`
}

// NewGetTaskSyncOK creates GetTaskSyncOK with default headers values
func NewGetTaskSyncOK() *GetTaskSyncOK {

	return &GetTaskSyncOK{}
}

// WithPayload adds the payload to the get task sync o k response
func (o *GetTaskSyncOK) WithPayload(payload *models.TaskStatus) *GetTaskSyncOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task sync o k response
func (o *GetTaskSyncOK) SetPayload(payload *models.TaskStatus) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskSyncOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTaskSyncBadRequestCode is the HTTP code returned for type GetTaskSyncBadRequest
const GetTaskSyncBadRequestCode int = 400

/*GetTaskSyncBadRequest Передан не UUID

swagger:response getTaskSyncBadRequest
*/
type GetTaskSyncBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetTaskSyncBadRequest creates GetTaskSyncBadRequest with default headers values
func NewGetTaskSyncBadRequest() *GetTaskSyncBadRequest {

	return &GetTaskSyncBadRequest{}
}

// WithPayload adds the payload to the get task sync bad request response
func (o *GetTaskSyncBadRequest) WithPayload(payload interface{}) *GetTaskSyncBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task sync bad request response
func (o *GetTaskSyncBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskSyncBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTaskSyncNotFoundCode is the HTTP code returned for type GetTaskSyncNotFound
const GetTaskSyncNotFoundCode int = 404

/*GetTaskSyncNotFound Задача с {taskId} не найдена

swagger:response getTaskSyncNotFound
*/
type GetTaskSyncNotFound struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetTaskSyncNotFound creates GetTaskSyncNotFound with default headers values
func NewGetTaskSyncNotFound() *GetTaskSyncNotFound {

	return &GetTaskSyncNotFound{}
}

// WithPayload adds the payload to the get task sync not found response
func (o *GetTaskSyncNotFound) WithPayload(payload interface{}) *GetTaskSyncNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task sync not found response
func (o *GetTaskSyncNotFound) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskSyncNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetTaskSyncInternalServerErrorCode is the HTTP code returned for type GetTaskSyncInternalServerError
const GetTaskSyncInternalServerErrorCode int = 500

/*GetTaskSyncInternalServerError Что-то пошло не так

swagger:response getTaskSyncInternalServerError
*/
type GetTaskSyncInternalServerError struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetTaskSyncInternalServerError creates GetTaskSyncInternalServerError with default headers values
func NewGetTaskSyncInternalServerError() *GetTaskSyncInternalServerError {

	return &GetTaskSyncInternalServerError{}
}

// WithPayload adds the payload to the get task sync internal server error response
func (o *GetTaskSyncInternalServerError) WithPayload(payload interface{}) *GetTaskSyncInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task sync internal server error response
func (o *GetTaskSyncInternalServerError) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskSyncInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

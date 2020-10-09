// Code generated by go-swagger; DO NOT EDIT.

package service_default_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/keptn/keptn/configuration-service/models"
)

// DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContentCode is the HTTP code returned for type DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent
const DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContentCode int = 204

/*DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent Success. Service default resource has been deleted.

swagger:response deleteProjectProjectNameServiceServiceNameResourceResourceUriNoContent
*/
type DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent struct {

	/*
	  In: Body
	*/
	Payload *models.Version `json:"body,omitempty"`
}

// NewDeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent creates DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent with default headers values
func NewDeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent() *DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent {

	return &DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent{}
}

// WithPayload adds the payload to the delete project project name service service name resource resource Uri no content response
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent) WithPayload(payload *models.Version) *DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project project name service service name resource resource Uri no content response
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent) SetPayload(payload *models.Version) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURINoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequestCode is the HTTP code returned for type DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest
const DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequestCode int = 400

/*DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest Failed. Service default resource could not be deleted.

swagger:response deleteProjectProjectNameServiceServiceNameResourceResourceUriBadRequest
*/
type DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest creates DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest with default headers values
func NewDeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest() *DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest {

	return &DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest{}
}

// WithPayload adds the payload to the delete project project name service service name resource resource Uri bad request response
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest) WithPayload(payload *models.Error) *DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project project name service service name resource resource Uri bad request response
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURIBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault Error

swagger:response deleteProjectProjectNameServiceServiceNameResourceResourceUriDefault
*/
type DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault creates DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault with default headers values
func NewDeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault(code int) *DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete project project name service service name resource resource URI default response
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault) WithStatusCode(code int) *DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete project project name service service name resource resource URI default response
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete project project name service service name resource resource URI default response
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault) WithPayload(payload *models.Error) *DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project project name service service name resource resource URI default response
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectProjectNameServiceServiceNameResourceResourceURIDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

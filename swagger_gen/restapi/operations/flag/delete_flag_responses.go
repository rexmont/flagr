// Code generated by go-swagger; DO NOT EDIT.

package flag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/rexmont/flagr/swagger_gen/models"
)

// DeleteFlagOKCode is the HTTP code returned for type DeleteFlagOK
const DeleteFlagOKCode int = 200

/*DeleteFlagOK OK deleted

swagger:response deleteFlagOK
*/
type DeleteFlagOK struct {
}

// NewDeleteFlagOK creates DeleteFlagOK with default headers values
func NewDeleteFlagOK() *DeleteFlagOK {

	return &DeleteFlagOK{}
}

// WriteResponse to the client
func (o *DeleteFlagOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*DeleteFlagDefault generic error response

swagger:response deleteFlagDefault
*/
type DeleteFlagDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteFlagDefault creates DeleteFlagDefault with default headers values
func NewDeleteFlagDefault(code int) *DeleteFlagDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteFlagDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete flag default response
func (o *DeleteFlagDefault) WithStatusCode(code int) *DeleteFlagDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete flag default response
func (o *DeleteFlagDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete flag default response
func (o *DeleteFlagDefault) WithPayload(payload *models.Error) *DeleteFlagDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete flag default response
func (o *DeleteFlagDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteFlagDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

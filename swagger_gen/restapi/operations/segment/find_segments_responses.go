// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/rexmont/flagr/swagger_gen/models"
)

// FindSegmentsOKCode is the HTTP code returned for type FindSegmentsOK
const FindSegmentsOKCode int = 200

/*FindSegmentsOK segments ordered by rank of the flag

swagger:response findSegmentsOK
*/
type FindSegmentsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Segment `json:"body,omitempty"`
}

// NewFindSegmentsOK creates FindSegmentsOK with default headers values
func NewFindSegmentsOK() *FindSegmentsOK {

	return &FindSegmentsOK{}
}

// WithPayload adds the payload to the find segments o k response
func (o *FindSegmentsOK) WithPayload(payload []*models.Segment) *FindSegmentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find segments o k response
func (o *FindSegmentsOK) SetPayload(payload []*models.Segment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindSegmentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Segment, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FindSegmentsDefault generic error response

swagger:response findSegmentsDefault
*/
type FindSegmentsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindSegmentsDefault creates FindSegmentsDefault with default headers values
func NewFindSegmentsDefault(code int) *FindSegmentsDefault {
	if code <= 0 {
		code = 500
	}

	return &FindSegmentsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find segments default response
func (o *FindSegmentsDefault) WithStatusCode(code int) *FindSegmentsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find segments default response
func (o *FindSegmentsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find segments default response
func (o *FindSegmentsDefault) WithPayload(payload *models.Error) *FindSegmentsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find segments default response
func (o *FindSegmentsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindSegmentsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

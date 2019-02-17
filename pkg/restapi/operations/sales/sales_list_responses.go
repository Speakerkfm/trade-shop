// Code generated by go-swagger; DO NOT EDIT.

package sales

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "trade-shop/pkg/models"
)

// SalesListOKCode is the HTTP code returned for type SalesListOK
const SalesListOKCode int = 200

/*SalesListOK Список лотов

swagger:response salesListOK
*/
type SalesListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Sale `json:"body,omitempty"`
}

// NewSalesListOK creates SalesListOK with default headers values
func NewSalesListOK() *SalesListOK {

	return &SalesListOK{}
}

// WithPayload adds the payload to the sales list o k response
func (o *SalesListOK) WithPayload(payload []*models.Sale) *SalesListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sales list o k response
func (o *SalesListOK) SetPayload(payload []*models.Sale) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SalesListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Sale, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// SalesListUnauthorizedCode is the HTTP code returned for type SalesListUnauthorized
const SalesListUnauthorizedCode int = 401

/*SalesListUnauthorized Пользователь не аутентифицирован в системе

swagger:response salesListUnauthorized
*/
type SalesListUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewSalesListUnauthorized creates SalesListUnauthorized with default headers values
func NewSalesListUnauthorized() *SalesListUnauthorized {

	return &SalesListUnauthorized{}
}

// WithPayload adds the payload to the sales list unauthorized response
func (o *SalesListUnauthorized) WithPayload(payload *models.ErrorResult) *SalesListUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sales list unauthorized response
func (o *SalesListUnauthorized) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SalesListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

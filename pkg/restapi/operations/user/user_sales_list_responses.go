// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "trade-shop/pkg/models"
)

// UserSalesListOKCode is the HTTP code returned for type UserSalesListOK
const UserSalesListOKCode int = 200

/*UserSalesListOK Список лотов пользователя

swagger:response userSalesListOK
*/
type UserSalesListOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Sale `json:"body,omitempty"`
}

// NewUserSalesListOK creates UserSalesListOK with default headers values
func NewUserSalesListOK() *UserSalesListOK {

	return &UserSalesListOK{}
}

// WithPayload adds the payload to the user sales list o k response
func (o *UserSalesListOK) WithPayload(payload []*models.Sale) *UserSalesListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user sales list o k response
func (o *UserSalesListOK) SetPayload(payload []*models.Sale) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserSalesListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Sale, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// UserSalesListUnauthorizedCode is the HTTP code returned for type UserSalesListUnauthorized
const UserSalesListUnauthorizedCode int = 401

/*UserSalesListUnauthorized Пользователь не аутентифицирован в системе

swagger:response userSalesListUnauthorized
*/
type UserSalesListUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewUserSalesListUnauthorized creates UserSalesListUnauthorized with default headers values
func NewUserSalesListUnauthorized() *UserSalesListUnauthorized {

	return &UserSalesListUnauthorized{}
}

// WithPayload adds the payload to the user sales list unauthorized response
func (o *UserSalesListUnauthorized) WithPayload(payload *models.ErrorResult) *UserSalesListUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user sales list unauthorized response
func (o *UserSalesListUnauthorized) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserSalesListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
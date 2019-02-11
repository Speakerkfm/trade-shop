// Code generated by go-swagger; DO NOT EDIT.

package sale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// SaleOKCode is the HTTP code returned for type SaleOK
const SaleOKCode int = 200

/*SaleOK OK

swagger:response saleOK
*/
type SaleOK struct {
}

// NewSaleOK creates SaleOK with default headers values
func NewSaleOK() *SaleOK {

	return &SaleOK{}
}

// WriteResponse to the client
func (o *SaleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SaleUnauthorizedCode is the HTTP code returned for type SaleUnauthorized
const SaleUnauthorizedCode int = 401

/*SaleUnauthorized Пользователь не авторизован

swagger:response saleUnauthorized
*/
type SaleUnauthorized struct {
}

// NewSaleUnauthorized creates SaleUnauthorized with default headers values
func NewSaleUnauthorized() *SaleUnauthorized {

	return &SaleUnauthorized{}
}

// WriteResponse to the client
func (o *SaleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

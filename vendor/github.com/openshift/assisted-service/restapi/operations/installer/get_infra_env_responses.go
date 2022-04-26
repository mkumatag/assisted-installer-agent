// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// GetInfraEnvOKCode is the HTTP code returned for type GetInfraEnvOK
const GetInfraEnvOKCode int = 200

/*GetInfraEnvOK Success.

swagger:response getInfraEnvOK
*/
type GetInfraEnvOK struct {

	/*
	  In: Body
	*/
	Payload *models.InfraEnv `json:"body,omitempty"`
}

// NewGetInfraEnvOK creates GetInfraEnvOK with default headers values
func NewGetInfraEnvOK() *GetInfraEnvOK {

	return &GetInfraEnvOK{}
}

// WithPayload adds the payload to the get infra env o k response
func (o *GetInfraEnvOK) WithPayload(payload *models.InfraEnv) *GetInfraEnvOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env o k response
func (o *GetInfraEnvOK) SetPayload(payload *models.InfraEnv) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvUnauthorizedCode is the HTTP code returned for type GetInfraEnvUnauthorized
const GetInfraEnvUnauthorizedCode int = 401

/*GetInfraEnvUnauthorized Unauthorized.

swagger:response getInfraEnvUnauthorized
*/
type GetInfraEnvUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetInfraEnvUnauthorized creates GetInfraEnvUnauthorized with default headers values
func NewGetInfraEnvUnauthorized() *GetInfraEnvUnauthorized {

	return &GetInfraEnvUnauthorized{}
}

// WithPayload adds the payload to the get infra env unauthorized response
func (o *GetInfraEnvUnauthorized) WithPayload(payload *models.InfraError) *GetInfraEnvUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env unauthorized response
func (o *GetInfraEnvUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvForbiddenCode is the HTTP code returned for type GetInfraEnvForbidden
const GetInfraEnvForbiddenCode int = 403

/*GetInfraEnvForbidden Forbidden.

swagger:response getInfraEnvForbidden
*/
type GetInfraEnvForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewGetInfraEnvForbidden creates GetInfraEnvForbidden with default headers values
func NewGetInfraEnvForbidden() *GetInfraEnvForbidden {

	return &GetInfraEnvForbidden{}
}

// WithPayload adds the payload to the get infra env forbidden response
func (o *GetInfraEnvForbidden) WithPayload(payload *models.InfraError) *GetInfraEnvForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env forbidden response
func (o *GetInfraEnvForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvNotFoundCode is the HTTP code returned for type GetInfraEnvNotFound
const GetInfraEnvNotFoundCode int = 404

/*GetInfraEnvNotFound Error.

swagger:response getInfraEnvNotFound
*/
type GetInfraEnvNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfraEnvNotFound creates GetInfraEnvNotFound with default headers values
func NewGetInfraEnvNotFound() *GetInfraEnvNotFound {

	return &GetInfraEnvNotFound{}
}

// WithPayload adds the payload to the get infra env not found response
func (o *GetInfraEnvNotFound) WithPayload(payload *models.Error) *GetInfraEnvNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env not found response
func (o *GetInfraEnvNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvMethodNotAllowedCode is the HTTP code returned for type GetInfraEnvMethodNotAllowed
const GetInfraEnvMethodNotAllowedCode int = 405

/*GetInfraEnvMethodNotAllowed Method Not Allowed.

swagger:response getInfraEnvMethodNotAllowed
*/
type GetInfraEnvMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfraEnvMethodNotAllowed creates GetInfraEnvMethodNotAllowed with default headers values
func NewGetInfraEnvMethodNotAllowed() *GetInfraEnvMethodNotAllowed {

	return &GetInfraEnvMethodNotAllowed{}
}

// WithPayload adds the payload to the get infra env method not allowed response
func (o *GetInfraEnvMethodNotAllowed) WithPayload(payload *models.Error) *GetInfraEnvMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env method not allowed response
func (o *GetInfraEnvMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvInternalServerErrorCode is the HTTP code returned for type GetInfraEnvInternalServerError
const GetInfraEnvInternalServerErrorCode int = 500

/*GetInfraEnvInternalServerError Error.

swagger:response getInfraEnvInternalServerError
*/
type GetInfraEnvInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfraEnvInternalServerError creates GetInfraEnvInternalServerError with default headers values
func NewGetInfraEnvInternalServerError() *GetInfraEnvInternalServerError {

	return &GetInfraEnvInternalServerError{}
}

// WithPayload adds the payload to the get infra env internal server error response
func (o *GetInfraEnvInternalServerError) WithPayload(payload *models.Error) *GetInfraEnvInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env internal server error response
func (o *GetInfraEnvInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvNotImplementedCode is the HTTP code returned for type GetInfraEnvNotImplemented
const GetInfraEnvNotImplementedCode int = 501

/*GetInfraEnvNotImplemented Not implemented.

swagger:response getInfraEnvNotImplemented
*/
type GetInfraEnvNotImplemented struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfraEnvNotImplemented creates GetInfraEnvNotImplemented with default headers values
func NewGetInfraEnvNotImplemented() *GetInfraEnvNotImplemented {

	return &GetInfraEnvNotImplemented{}
}

// WithPayload adds the payload to the get infra env not implemented response
func (o *GetInfraEnvNotImplemented) WithPayload(payload *models.Error) *GetInfraEnvNotImplemented {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env not implemented response
func (o *GetInfraEnvNotImplemented) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInfraEnvServiceUnavailableCode is the HTTP code returned for type GetInfraEnvServiceUnavailable
const GetInfraEnvServiceUnavailableCode int = 503

/*GetInfraEnvServiceUnavailable Unavailable.

swagger:response getInfraEnvServiceUnavailable
*/
type GetInfraEnvServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInfraEnvServiceUnavailable creates GetInfraEnvServiceUnavailable with default headers values
func NewGetInfraEnvServiceUnavailable() *GetInfraEnvServiceUnavailable {

	return &GetInfraEnvServiceUnavailable{}
}

// WithPayload adds the payload to the get infra env service unavailable response
func (o *GetInfraEnvServiceUnavailable) WithPayload(payload *models.Error) *GetInfraEnvServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get infra env service unavailable response
func (o *GetInfraEnvServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInfraEnvServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
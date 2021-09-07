// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fugue/regula/pkg/swagger/models"
)

// UpdateNotificationReader is a Reader for the UpdateNotification structure.
type UpdateNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNotificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateNotificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateNotificationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateNotificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateNotificationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNotificationOK creates a UpdateNotificationOK with default headers values
func NewUpdateNotificationOK() *UpdateNotificationOK {
	return &UpdateNotificationOK{}
}

/*UpdateNotificationOK handles this case with default header values.

New notification details.
*/
type UpdateNotificationOK struct {
	Payload *models.Notification
}

func (o *UpdateNotificationOK) Error() string {
	return fmt.Sprintf("[PUT /notifications/{notification_id}][%d] updateNotificationOK  %+v", 200, o.Payload)
}

func (o *UpdateNotificationOK) GetPayload() *models.Notification {
	return o.Payload
}

func (o *UpdateNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Notification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationBadRequest creates a UpdateNotificationBadRequest with default headers values
func NewUpdateNotificationBadRequest() *UpdateNotificationBadRequest {
	return &UpdateNotificationBadRequest{}
}

/*UpdateNotificationBadRequest handles this case with default header values.

BadRequestError
*/
type UpdateNotificationBadRequest struct {
	Payload *models.BadRequestError
}

func (o *UpdateNotificationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /notifications/{notification_id}][%d] updateNotificationBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateNotificationBadRequest) GetPayload() *models.BadRequestError {
	return o.Payload
}

func (o *UpdateNotificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationUnauthorized creates a UpdateNotificationUnauthorized with default headers values
func NewUpdateNotificationUnauthorized() *UpdateNotificationUnauthorized {
	return &UpdateNotificationUnauthorized{}
}

/*UpdateNotificationUnauthorized handles this case with default header values.

AuthenticationError
*/
type UpdateNotificationUnauthorized struct {
	Payload *models.AuthenticationError
}

func (o *UpdateNotificationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /notifications/{notification_id}][%d] updateNotificationUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateNotificationUnauthorized) GetPayload() *models.AuthenticationError {
	return o.Payload
}

func (o *UpdateNotificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationForbidden creates a UpdateNotificationForbidden with default headers values
func NewUpdateNotificationForbidden() *UpdateNotificationForbidden {
	return &UpdateNotificationForbidden{}
}

/*UpdateNotificationForbidden handles this case with default header values.

AuthorizationError
*/
type UpdateNotificationForbidden struct {
	Payload *models.AuthorizationError
}

func (o *UpdateNotificationForbidden) Error() string {
	return fmt.Sprintf("[PUT /notifications/{notification_id}][%d] updateNotificationForbidden  %+v", 403, o.Payload)
}

func (o *UpdateNotificationForbidden) GetPayload() *models.AuthorizationError {
	return o.Payload
}

func (o *UpdateNotificationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationNotFound creates a UpdateNotificationNotFound with default headers values
func NewUpdateNotificationNotFound() *UpdateNotificationNotFound {
	return &UpdateNotificationNotFound{}
}

/*UpdateNotificationNotFound handles this case with default header values.

NotFoundError
*/
type UpdateNotificationNotFound struct {
	Payload *models.NotFoundError
}

func (o *UpdateNotificationNotFound) Error() string {
	return fmt.Sprintf("[PUT /notifications/{notification_id}][%d] updateNotificationNotFound  %+v", 404, o.Payload)
}

func (o *UpdateNotificationNotFound) GetPayload() *models.NotFoundError {
	return o.Payload
}

func (o *UpdateNotificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFoundError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNotificationInternalServerError creates a UpdateNotificationInternalServerError with default headers values
func NewUpdateNotificationInternalServerError() *UpdateNotificationInternalServerError {
	return &UpdateNotificationInternalServerError{}
}

/*UpdateNotificationInternalServerError handles this case with default header values.

InternalServerError
*/
type UpdateNotificationInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *UpdateNotificationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /notifications/{notification_id}][%d] updateNotificationInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateNotificationInternalServerError) GetPayload() *models.InternalServerError {
	return o.Payload
}

func (o *UpdateNotificationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
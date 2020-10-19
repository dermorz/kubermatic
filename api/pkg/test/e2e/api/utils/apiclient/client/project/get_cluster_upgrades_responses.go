// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// GetClusterUpgradesReader is a Reader for the GetClusterUpgrades structure.
type GetClusterUpgradesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterUpgradesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterUpgradesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterUpgradesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterUpgradesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterUpgradesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterUpgradesOK creates a GetClusterUpgradesOK with default headers values
func NewGetClusterUpgradesOK() *GetClusterUpgradesOK {
	return &GetClusterUpgradesOK{}
}

/*GetClusterUpgradesOK handles this case with default header values.

MasterVersion
*/
type GetClusterUpgradesOK struct {
	Payload []*models.MasterVersion
}

func (o *GetClusterUpgradesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesOK  %+v", 200, o.Payload)
}

func (o *GetClusterUpgradesOK) GetPayload() []*models.MasterVersion {
	return o.Payload
}

func (o *GetClusterUpgradesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterUpgradesUnauthorized creates a GetClusterUpgradesUnauthorized with default headers values
func NewGetClusterUpgradesUnauthorized() *GetClusterUpgradesUnauthorized {
	return &GetClusterUpgradesUnauthorized{}
}

/*GetClusterUpgradesUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetClusterUpgradesUnauthorized struct {
}

func (o *GetClusterUpgradesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesUnauthorized ", 401)
}

func (o *GetClusterUpgradesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterUpgradesForbidden creates a GetClusterUpgradesForbidden with default headers values
func NewGetClusterUpgradesForbidden() *GetClusterUpgradesForbidden {
	return &GetClusterUpgradesForbidden{}
}

/*GetClusterUpgradesForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetClusterUpgradesForbidden struct {
}

func (o *GetClusterUpgradesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/upgrades][%d] getClusterUpgradesForbidden ", 403)
}

func (o *GetClusterUpgradesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterUpgradesDefault creates a GetClusterUpgradesDefault with default headers values
func NewGetClusterUpgradesDefault(code int) *GetClusterUpgradesDefault {
	return &GetClusterUpgradesDefault{
		_statusCode: code,
	}
}

/*GetClusterUpgradesDefault handles this case with default header values.

errorResponse
*/
type GetClusterUpgradesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster upgrades default response
func (o *GetClusterUpgradesDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterUpgradesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/upgrades][%d] getClusterUpgrades default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterUpgradesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterUpgradesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetOidcClusterKubeconfigV2Reader is a Reader for the GetOidcClusterKubeconfigV2 structure.
type GetOidcClusterKubeconfigV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOidcClusterKubeconfigV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOidcClusterKubeconfigV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOidcClusterKubeconfigV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOidcClusterKubeconfigV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOidcClusterKubeconfigV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOidcClusterKubeconfigV2OK creates a GetOidcClusterKubeconfigV2OK with default headers values
func NewGetOidcClusterKubeconfigV2OK() *GetOidcClusterKubeconfigV2OK {
	return &GetOidcClusterKubeconfigV2OK{}
}

/*GetOidcClusterKubeconfigV2OK handles this case with default header values.

Kubeconfig is a clusters kubeconfig
*/
type GetOidcClusterKubeconfigV2OK struct {
	Payload []uint8
}

func (o *GetOidcClusterKubeconfigV2OK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidckubeconfig][%d] getOidcClusterKubeconfigV2OK  %+v", 200, o.Payload)
}

func (o *GetOidcClusterKubeconfigV2OK) GetPayload() []uint8 {
	return o.Payload
}

func (o *GetOidcClusterKubeconfigV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOidcClusterKubeconfigV2Unauthorized creates a GetOidcClusterKubeconfigV2Unauthorized with default headers values
func NewGetOidcClusterKubeconfigV2Unauthorized() *GetOidcClusterKubeconfigV2Unauthorized {
	return &GetOidcClusterKubeconfigV2Unauthorized{}
}

/*GetOidcClusterKubeconfigV2Unauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetOidcClusterKubeconfigV2Unauthorized struct {
}

func (o *GetOidcClusterKubeconfigV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidckubeconfig][%d] getOidcClusterKubeconfigV2Unauthorized ", 401)
}

func (o *GetOidcClusterKubeconfigV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOidcClusterKubeconfigV2Forbidden creates a GetOidcClusterKubeconfigV2Forbidden with default headers values
func NewGetOidcClusterKubeconfigV2Forbidden() *GetOidcClusterKubeconfigV2Forbidden {
	return &GetOidcClusterKubeconfigV2Forbidden{}
}

/*GetOidcClusterKubeconfigV2Forbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetOidcClusterKubeconfigV2Forbidden struct {
}

func (o *GetOidcClusterKubeconfigV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidckubeconfig][%d] getOidcClusterKubeconfigV2Forbidden ", 403)
}

func (o *GetOidcClusterKubeconfigV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOidcClusterKubeconfigV2Default creates a GetOidcClusterKubeconfigV2Default with default headers values
func NewGetOidcClusterKubeconfigV2Default(code int) *GetOidcClusterKubeconfigV2Default {
	return &GetOidcClusterKubeconfigV2Default{
		_statusCode: code,
	}
}

/*GetOidcClusterKubeconfigV2Default handles this case with default header values.

errorResponse
*/
type GetOidcClusterKubeconfigV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get oidc cluster kubeconfig v2 default response
func (o *GetOidcClusterKubeconfigV2Default) Code() int {
	return o._statusCode
}

func (o *GetOidcClusterKubeconfigV2Default) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/oidckubeconfig][%d] getOidcClusterKubeconfigV2 default  %+v", o._statusCode, o.Payload)
}

func (o *GetOidcClusterKubeconfigV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetOidcClusterKubeconfigV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}